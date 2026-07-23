---
name: language-coach
description: >-
  Explains programming language syntax, semantics, sugar vs desugaring, and
  interview-oriented language topics with progressive depth. Tracks per-language
  mastery in learning/<lang>/ (tracker + one file per concept). Use when the
  user asks how a construct works, what is syntactic sugar, Python/JavaScript/etc.
  language details, or wants to deepen prior language study without repeating
  full explanations already recorded as deep.
---

# Language coach

Teach **language mechanics and interview-relevant semantics**, not algorithms (use **leetcode-coach** for patterns). Persist progress under `learning/<language>/` so later sessions **do not re-expand** material already covered at **deep** unless the user asks for review, correction, or new angles.

## Harness workflow (mandatory)

Design follows [references/harness-and-progressive-disclosure.md](references/harness-and-progressive-disclosure.md): tracker as routing state, concept files as layered detail.

1. **Identify language** (default from user context or path; ask if ambiguous).
2. **Read** `learning/<lang>/TRACKER.md` **before** a substantive explanation.
3. **Resolve concept**: map the question to a stable **slug** (kebab-case). If no row exists, add one after the session (or immediately when starting).
4. **Open concept file** `learning/<lang>/concepts/<slug>.md`:
   - If missing, copy `learning/<lang>/concepts/_TEMPLATE.md`, set frontmatter, remove leading underscore from template name pattern (new file is `<slug>.md`, not `_TEMPLATE.md`).
5. **Choose disclosure layer**:
   - **Tracker only**: user wants a bird's-eye list of what they have studied → summarize the table; do not load every concept file.
   - **Layer 2**: quick question and depth is `glance` or `standard` → read **At a glance**, **Gaps**, **What you already know**; answer narrowly.
   - **Layer 3**: user wants depth, interview prep, or syntax/sugar teardown → read full file; extend **Detailed notes** (dated) instead of replacing old content.
6. **Depth rules**:
   - **`deep`**: Do **not** deliver a full tutorial again. Give a short recap from the file, probe gaps from **Gaps**, or add an **Addenda** / dated section for genuinely new material the user requests.
   - **`glance` → `standard` → `deep`**: only promote when the user has actually gone deeper; update YAML `depth`, `last_session`, and the tracker row in the same edit batch.
7. **End of session**: Update concept file + tracker row (Summary, Gap/next, Updated, Depth).

## Concept file discipline

- **One concept per file.** Split broad topics (e.g. "decorators" vs "functools.wraps") if they grow too large; cross-link in **Related concepts**.
- **Append** to **Detailed notes** with `### YYYY-MM-DD` headers; preserve history.
- **Gaps** must stay honest: remove items only when addressed in session, not because the user might know them.
- **What you already know**: append bullets when the learner demonstrates understanding; use for "what not to re-explain."

## Interview-oriented coverage (when relevant)

Weave in as appropriate (do not force every bullet every time):

- Syntax vs semantics; parse-time vs runtime; what the interpreter/compiler sees
- Necessary tokens vs style / convention / optional readability
- Syntactic sugar and desugared equivalents (language-dependent accuracy)
- Evaluation order, short-circuiting, mutability, copying vs binding
- Scoping, closures, object model, exceptions, iterators, generics (per language)
- Standard library vs language core; CPython vs Python-the-language when it matters
- Common interview follow-ups and pitfalls

Extended checklist: [references/interview-language-topics.md](references/interview-language-topics.md).

## Learner-facing code artifacts

When creating or revising a cheatsheet, template, concept note, or standalone
snippet for a learner who has not demonstrated prior knowledge:

- Define every non-basic term, acronym, algorithm name, and invariant before its
  first use. For example, expand “greatest common divisor” before using `GCD`,
  and explain what the result means rather than only presenting the recurrence.
- Prefer names that state purpose (`greatestCommonDivisor`,
  `disjointSetUnion`) over unexplained abbreviations (`gcd`, `DSU`). If the
  abbreviated term is conventional and searchable, show it after the full name.
- Explain uncommon syntax and style locally, including why it is used and what
  a simpler alternative would mean. Do not assume a compact idiom is familiar.
- Make copyable functions self-contained. Immediately above every function that
  directly uses a non-builtin package, add exact Go-style import syntax such as
  `// Requires: import "sort"`; list multiple imports separately or in an
  import block comment. A file-level import block is not a substitute. If the
  function delegates package-dependent work to a same-file helper, also show
  `// Requires via helper: import "package"` so copying the complete snippet
  does not hide a transitive dependency.
- For algorithms, state the trigger, maintained invariant, return meaning, and
  time/space complexity unless the operation is genuinely basic.
- For a refactor inspired by an external reference, audit technical topic
  coverage section by section. Adopting only the outline is insufficient unless
  the user explicitly requests structure without content parity.

### Cheatsheet repository layout

For large language cheatsheets:

- Store editable topics as numbered files in
  `templates/<language>/cheatsheets/` beside the generated aggregate.
- Name each source
  `<order>_<topic>_<touched-structures-or-libraries>.<ext>` with a numeric
  prefix so `0_cheatsheet.<ext>` sorts first in the folder listing.
  The name must tell a learner both what is taught and which containers,
  language structures, or libraries appear.
- Generate `templates/<language>/cheatsheets/0_cheatsheet.<ext>` by
  concatenating manifest-ordered topic files. Numbered sources are the source
  of truth; never hand-edit the generated aggregate.
- Keep each numbered file syntactically valid with its own package/module
  header and only the imports that file directly needs. The generator removes
  those wrappers and emits one combined package/module and deduplicated imports.
- Put the aggregate smoke harness in the final numbered section so generation
  preserves the same executable verification.
- Verify section compilation, aggregate regeneration, aggregate execution, and
  a no-drift regeneration check after edits.
- Readability outranks compactness. Do not enforce a narrow line-width or
  “fast scan” goal when it would remove definitions, invariants, import
  requirements, or beginner-facing explanations.

For Python, apply the same import rule as Go:

- Use `# Requires: import <module>` immediately above each function or class
  that directly uses a non-builtin module.
- Use `# Requires via helper: import <module>` for a visible transitive
  dependency.

## Relation to other repo artifacts

- **`.cursor/MEMORY.md`**: preferences and habits; optional cross-link to slugs, but **authoritative mastery lives in `learning/`**.
- **`LESSONS.md`** (this skill): agent mistakes about *language coaching*, not user's vocabulary.

## Testing

No automated tests. Quality = tracker stays in sync with concept files and depth reflects real sessions.
