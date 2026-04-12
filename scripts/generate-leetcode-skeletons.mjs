#!/usr/bin/env node
/**
 * Creates solution skeleton files from LeetCode official starter code (GraphQL).
 * Reads slugs from data/leetcode/python/{first-100,first-300,last-100}.json (same catalog as golang).
 *
 * Usage:
 *   node scripts/generate-leetcode-skeletons.mjs              # skip existing files
 *   node scripts/generate-leetcode-skeletons.mjs --force      # overwrite all
 *   node scripts/generate-leetcode-skeletons.mjs --golang-only # only Go (still fetches each problem once)
 */
import { readFile, writeFile, mkdir, access } from "node:fs/promises";
import { constants as fsConstants } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { LeetCode } from "leetcode-query";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = join(__dirname, "..");
const DATA_DIR = join(ROOT, "data", "leetcode", "python");
const PY_OUT = join(ROOT, "solutions", "python");
const GO_OUT = join(ROOT, "solutions", "golang");

const force = process.argv.includes("--force");
const golangOnly = process.argv.includes("--golang-only");
const pythonOnly = process.argv.includes("--python-only");
if (golangOnly && pythonOnly) {
  console.error("Use at most one of --golang-only / --python-only");
  process.exit(1);
}

function padId(frontendId) {
  const n = String(frontendId).trim();
  return n.length >= 4 ? n : n.padStart(4, "0");
}

function filenameBase(frontendId, titleSlug) {
  return `${padId(frontendId)}-${titleSlug}`;
}

async function loadProblemSlugs() {
  const seen = new Map();
  for (const name of ["first-100.json", "first-300.json", "last-100.json"]) {
    const raw = await readFile(join(DATA_DIR, name), "utf8");
    const data = JSON.parse(raw);
    for (const p of data.problems ?? []) {
      if (!seen.has(p.titleSlug)) {
        seen.set(p.titleSlug, {
          frontendId: p.frontendId,
          title: p.title,
          titleSlug: p.titleSlug,
        });
      }
    }
  }
  return [...seen.values()].sort(
    (a, b) => Number(a.frontendId) - Number(b.frontendId)
  );
}

// LeetCode Go stubs wrap helper types in slash-star doc blocks; unwrap to valid Go.
// LeetCode Go doc blocks mix prose and declarations; comment only depth-0 non-decl lines.
function goCommentNonDeclarations(inner) {
  const lines = inner.split("\n");
  const out = [];
  let depth = 0;
  for (const line of lines) {
    const trimmed = line.trim();
    if (!trimmed) {
      out.push("");
      continue;
    }
    const opens = (trimmed.match(/\{/g) || []).length;
    const closes = (trimmed.match(/\}/g) || []).length;
    if (depth === 0) {
      const isTopDecl = /^(type|func|import|var|const)\s/.test(trimmed);
      if (isTopDecl) {
        out.push(line.trimEnd());
        depth += opens - closes;
      } else {
        out.push(`// ${trimmed}`);
        depth += opens - closes;
      }
    } else {
      out.push(line.trimEnd());
      depth += opens - closes;
    }
  }
  return out.join("\n");
}

function unwrapGoCommentBlocks(code) {
  return code.replace(/\/\*\*\s*\n([\s\S]*?)\*\//g, (_, block) => {
    const stripped = block
      .split("\n")
      .map((line) => line.replace(/^\s*\*\s?/, "").replace(/\s+$/, ""))
      .join("\n")
      .trim();
    if (!stripped) return "";
    const inner = goCommentNonDeclarations(stripped);
    return `\n${inner}\n`;
  });
}

function getIndent(line) {
  const m = line.match(/^(\s*)/);
  return m ? m[1].length : 0;
}

/** Insert `pass` into empty method bodies (Python requires a statement after `:`). */
function fillEmptyPythonDefs(code) {
  const lines = code.split("\n");
  const out = [];
  for (let i = 0; i < lines.length; i++) {
    out.push(lines[i]);
    const line = lines[i];
    const trimmed = line.trim();
    if (!trimmed.startsWith("def ") || !trimmed.endsWith(":")) continue;
    const ind = getIndent(line);
    let j = i + 1;
    while (j < lines.length && lines[j].trim() === "") j++;
    if (j >= lines.length) {
      out.push(`${" ".repeat(ind + 4)}pass`);
      continue;
    }
    const nextLine = lines[j];
    const ni = getIndent(nextLine);
    const nt = nextLine.trim();
    if (
      ni <= ind &&
      (nt.startsWith("def ") ||
        nt.startsWith("class ") ||
        nt.startsWith("#"))
    ) {
      out.push(`${" ".repeat(ind + 4)}pass`);
    }
  }
  return out.join("\n");
}

function inferTypingImports(code) {
  const names = [];
  const want = (re, name) => {
    if (re.test(code) && !names.includes(name)) names.push(name);
  };
  want(/\bList\b(?!\w)/, "List");
  want(/\bDict\b(?!\w)/, "Dict");
  want(/\bSet\b(?!\w)/, "Set");
  want(/\bTuple\b(?!\w)/, "Tuple");
  want(/\bOptional\b(?!\w)/, "Optional");
  want(/\bUnion\b(?!\w)/, "Union");
  want(/\bAny\b(?!\w)/, "Any");
  want(/\bDeque\b(?!\w)/, "Deque");
  want(/\bDefaultDict\b(?!\w)/, "DefaultDict");
  want(/\bCounter\b(?!\w)/, "Counter");
  want(/\bIterable\b(?!\w)/, "Iterable");
  want(/\bIterator\b(?!\w)/, "Iterator");
  want(/\bCallable\b(?!\w)/, "Callable");
  want(/\bMapping\b(?!\w)/, "Mapping");
  if (names.length === 0) return "";
  return `from typing import ${names.sort().join(", ")}\n\n`;
}

function buildPythonFile({ frontendId, title, titleSlug }, snippet) {
  const header = `# LeetCode ${frontendId} — ${title}\n# https://leetcode.com/problems/${titleSlug}/\n`;
  let body = (snippet ?? "").trimStart();
  if (!body) {
    body = `class Solution:\n    pass\n`;
    return `${header}\n${body}\n`;
  }
  const typingBlock = inferTypingImports(body);
  body = fillEmptyPythonDefs(body);
  return `${header}\n${typingBlock}${body}\n`;
}

/** Empty `func` bodies are invalid Go; LeetCode accepts them in-editor. Use panic so `go build` works locally. */
function stubEmptyGoFunctions(code) {
  let pos = 0;
  let out = "";
  while (pos < code.length) {
    const idx = code.indexOf("func ", pos);
    if (idx === -1) {
      out += code.slice(pos);
      break;
    }
    out += code.slice(pos, idx);
    const open = code.indexOf("{", idx);
    if (open === -1) {
      out += code.slice(idx);
      break;
    }
    let depth = 0;
    let j = open;
    for (; j < code.length; j++) {
      const c = code[j];
      if (c === "{") depth++;
      else if (c === "}") {
        depth--;
        if (depth === 0) {
          j++;
          break;
        }
      }
    }
    const funcBlock = code.slice(idx, j);
    const lb = funcBlock.indexOf("{");
    const rb = funcBlock.lastIndexOf("}");
    const inner = funcBlock.slice(lb + 1, rb);
    if (inner.trim() === "") {
      out +=
        funcBlock.slice(0, lb + 1) + "\n\tpanic(\"TODO\")\n" + funcBlock.slice(rb);
    } else {
      out += funcBlock;
    }
    pos = j;
  }
  return out;
}

function buildGoFile({ frontendId, title, titleSlug }, snippet) {
  const header = `package main\n\n// LeetCode ${frontendId} — ${title}\n// https://leetcode.com/problems/${titleSlug}/\n`;
  let body = (snippet ?? "").trim();
  if (!body) {
    return `${header}\n// No official Go starter returned; open the problem on LeetCode and paste the signature.\n`;
  }
  body = unwrapGoCommentBlocks(body);
  body = stubEmptyGoFunctions(body);
  return `${header}\n${body}\n\n// Local compile hook (LeetCode runs your func without this).\nfunc main() {}\n`;
}

async function main() {
  const problems = await loadProblemSlugs();
  const lc = new LeetCode();
  await lc.initialized;

  await mkdir(PY_OUT, { recursive: true });
  await mkdir(GO_OUT, { recursive: true });

  let pyW = 0,
    pySkip = 0,
    goW = 0,
    goSkip = 0,
    err = 0;

  for (const meta of problems) {
    const base = filenameBase(meta.frontendId, meta.titleSlug);
    const pyPath = join(PY_OUT, `${base}.py`);
    const goPath = join(GO_OUT, `${base}.go`);

    let question;
    try {
      question = await lc.problem(meta.titleSlug);
    } catch (e) {
      console.error(`fetch failed ${meta.titleSlug}:`, e.message);
      err++;
      continue;
    }

    const pySnip = question?.codeSnippets?.find((s) => s.langSlug === "python3");
    const goSnip = question?.codeSnippets?.find((s) => s.langSlug === "golang");

    if (!golangOnly) {
      const pyContent = buildPythonFile(meta, pySnip?.code);
      if (!force && (await fileExists(pyPath))) {
        pySkip++;
      } else {
        await writeFile(pyPath, pyContent, "utf8");
        pyW++;
      }
    }

    if (!pythonOnly) {
      const goContent = buildGoFile(meta, goSnip?.code);
      if (!force && (await fileExists(goPath))) {
        goSkip++;
      } else {
        await writeFile(goPath, goContent, "utf8");
        goW++;
      }
    }
  }

  console.log(
    JSON.stringify(
      {
        problems: problems.length,
        python: { wrote: pyW, skippedExisting: pySkip },
        golang: { wrote: goW, skippedExisting: goSkip },
        fetchErrors: err,
        force,
        golangOnly,
        pythonOnly,
      },
      null,
      2
    )
  );
}

async function fileExists(p) {
  try {
    await access(p, fsConstants.F_OK);
    return true;
  } catch {
    return false;
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
