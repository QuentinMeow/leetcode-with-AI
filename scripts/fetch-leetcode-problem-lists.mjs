#!/usr/bin/env node
/**
 * Fetches LeetCode catalog slices (global, algorithms essentials): first 100, first 300,
 * and last 100. Catalog is shared across languages; output is duplicated under python/
 * and golang/ for separate tracks. Requires network access to leetcode.com.
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
/** Problems per `lc.problems` page (API caps around 100). */
const PAGE_SIZE = 100;
const LIMIT_FIRST_EXTENDED = 300;
const LIMIT_LAST = 100;

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

/** First N problems in category order (paginates; API returns at most ~100 per call). */
async function fetchFirstNProblems(lc, n) {
  let totalInCategory = 0;
  const out = [];
  for (let offset = 0; offset < n; offset += PAGE_SIZE) {
    const limit = Math.min(PAGE_SIZE, n - offset);
    const page = await lc.problems({
      category: CATEGORY,
      offset,
      limit,
      filters: {},
    });
    totalInCategory = page?.total ?? totalInCategory;
    const batch = normalizeQuestions(page?.questions);
    if (!batch.length) break;
    out.push(...batch);
    if (out.length >= n) break;
  }
  return {
    totalInCategory,
    questions: out.slice(0, n),
  };
}

async function main() {
  const lc = new LeetCode();
  await lc.initialized;

  const { totalInCategory: total, questions: first300Questions } =
    await fetchFirstNProblems(lc, LIMIT_FIRST_EXTENDED);
  const first100Questions = first300Questions.slice(0, 100);

  const lastOffset = Math.max(0, total - LIMIT_LAST);
  const lastPage = await lc.problems({
    category: CATEGORY,
    offset: lastOffset,
    limit: LIMIT_LAST,
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
        limit: 100,
        totalInCategory: total,
        fetchedAt,
        submissionLanguages: ["python3", "golang"],
        note: commonNote,
      },
      problems: first100Questions,
    });

    await writeTrack(track, "first-300", {
      meta: {
        track,
        slice: "first-300",
        categorySlug: CATEGORY,
        skip: 0,
        limit: LIMIT_FIRST_EXTENDED,
        totalInCategory: total,
        fetchedAt,
        submissionLanguages: ["python3", "golang"],
        note: commonNote,
      },
      problems: first300Questions,
    });

    await writeTrack(track, "last-100", {
      meta: {
        track,
        slice: "last-100",
        categorySlug: CATEGORY,
        skip: lastOffset,
        limit: LIMIT_LAST,
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
