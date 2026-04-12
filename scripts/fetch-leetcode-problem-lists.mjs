#!/usr/bin/env node
/**
 * Fetches first / last 100 LeetCode catalog problems (global, algorithms essentials).
 * Catalog is shared across languages; output is duplicated under python/ and golang/
 * for separate tracks. Requires network access to leetcode.com.
 *
 * Next step: turn those JSON lists into solution stubs (official LeetCode signatures):
 *   npm run leetcode:generate-skeletons
 * See solutions/README.md.
 */
import { mkdir, writeFile } from "node:fs/promises";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import { LeetCode } from "leetcode-query";

const __dirname = dirname(fileURLToPath(import.meta.url));
const ROOT = join(__dirname, "..");
const OUT_BASE = join(ROOT, "data", "leetcode");

const CATEGORY = "all-code-essentials";
const LIMIT = 100;

function normalizeQuestions(questions) {
  return (questions ?? []).map((q) => ({
    frontendId: q.questionFrontendId,
    title: q.title,
    titleSlug: q.titleSlug,
    difficulty: q.difficulty,
    acRate: q.acRate,
    isPaidOnly: q.isPaidOnly,
    topicTags: (q.topicTags ?? []).map((t) => t.slug),
  }));
}

async function writeTrack(track, slice, payload) {
  const dir = join(OUT_BASE, track);
  await mkdir(dir, { recursive: true });
  const file = join(dir, `${slice}.json`);
  await writeFile(file, JSON.stringify(payload, null, 2) + "\n", "utf8");
  console.log("wrote", file);
}

async function main() {
  const lc = new LeetCode();
  await lc.initialized;

  const firstPage = await lc.problems({
    category: CATEGORY,
    offset: 0,
    limit: LIMIT,
    filters: {},
  });

  const total = firstPage?.total ?? 0;
  const firstQuestions = normalizeQuestions(firstPage?.questions);

  const lastOffset = Math.max(0, total - LIMIT);
  const lastPage = await lc.problems({
    category: CATEGORY,
    offset: lastOffset,
    limit: LIMIT,
    filters: {},
  });
  const lastQuestions = normalizeQuestions(lastPage?.questions);

  const fetchedAt = new Date().toISOString();
  const commonNote =
    "LeetCode problem catalog is the same for all supported languages; Python 3 and Go are standard judge languages.";

  for (const track of ["python", "golang"]) {
    await writeTrack(track, "first-100", {
      meta: {
        track,
        slice: "first-100",
        categorySlug: CATEGORY,
        skip: 0,
        limit: LIMIT,
        totalInCategory: total,
        fetchedAt,
        submissionLanguages: ["python3", "golang"],
        note: commonNote,
      },
      problems: firstQuestions,
    });

    await writeTrack(track, "last-100", {
      meta: {
        track,
        slice: "last-100",
        categorySlug: CATEGORY,
        skip: lastOffset,
        limit: LIMIT,
        totalInCategory: total,
        fetchedAt,
        submissionLanguages: ["python3", "golang"],
        note: commonNote,
      },
      problems: lastQuestions,
    });
  }

  console.log(`total problems in category "${CATEGORY}": ${total}`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
