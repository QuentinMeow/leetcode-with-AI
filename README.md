# LeetCode with AI

A self-evolving LeetCode practice environment in Cursor: **your** machine runs a **local** LeetCode MCP server (not a shared public service), project rules keep problem-solving disciplined, and skills provide coaching plus interview-style follow-ups. Solutions live under `solutions/<language>/`.

---

## How it fits together

| Piece | What it is | What you do |
|-------|------------|-------------|
| **LeetCode MCP** | A Node process Cursor starts from **`.cursor/mcp.json`** (you create this file locally; it is **not** committed). | Copy `.cursor/mcp-example.json` → `.cursor/mcp.json`, fill in paths, reload Cursor. |
| **Cursor rules** | Always-on guidance in `.cursor/rules/` (problem-solving steps, memory, banned tools). | Nothing special — they steer the agent automatically. |
| **Skills** | `leetcode-coach` and `interview-simulator` under `.cursor/skills/`. | Ask for hints, reviews, or “interview me” in natural language. |
| **`.cursor/MEMORY.md`** | Gitignored file the agent updates with durable preferences, weak areas, and progress patterns. | Optional: skim it; let the agent maintain it after substantive sessions. |
| **`progress.md`** | Committed tracker table for problems and languages. | Update with the agent or by hand as you solve problems. |

There is **no** third-party “public MCP host” for LeetCode in this setup. The server code is the npm package `@jinzcdev/leetcode-mcp-server`, installed in **this repo’s** `node_modules/`, executed by **your** `node`.

---

## Prerequisites

- **Cursor** with MCP support enabled for the workspace.
- **Node.js** (v20+ recommended; this repo is tested with a current LTS via nvm).
- A **LeetCode** account on [leetcode.com](https://leetcode.com) (for logged-in features; many read-only tools work without a session cookie).

---

## One-time setup

### 1. Install the MCP server package

From the **repository root**:

```bash
npm install
```

If your default npm registry is a **corporate mirror** that does not include `@jinzcdev/leetcode-mcp-server`, install from the public registry once:

```bash
npm install @jinzcdev/leetcode-mcp-server --registry https://registry.npmjs.org
```

That populates `node_modules/` (listed in `.gitignore` — not committed).

### 2. Create `.cursor/mcp.json` from the example (not committed)

**`.cursor/mcp.json` is gitignored** so machine-specific paths and session cookies never land in git. The repo ships **`.cursor/mcp-example.json`** as the template.

From the repository root:

```bash
cp .cursor/mcp-example.json .cursor/mcp.json
```

Edit `.cursor/mcp.json` and replace **every** placeholder:

| Placeholder | Replace with |
|-------------|----------------|
| `/ABSOLUTE/PATH/TO/node` | Output of `which node` (Cursor does not load your shell profile; nvm users **must** use an absolute path). |
| `/ABSOLUTE/PATH/TO/THIS-REPO` | Absolute path to **this** repository root (where `package.json` lives). |
| `/ABSOLUTE/PATH/TO/node-directory` | Directory containing the `node` binary (same as `dirname $(which node)`). |
| `LEETCODE_SESSION` | Leave `""` for read-only tools, or paste your `LEETCODE_SESSION` cookie value for authenticated tools (see below). **Never commit** this value. |

Cursor reads the filename **`mcp.json`** only — renaming the example file in place (copy → `mcp.json`) is the intended workflow.

**Why not `npx` in `command`?** Corporate npm registries sometimes omit this package, so `npx` can hang. Pointing `args` at `node_modules/.../build/index.js` after `npm install` is reliable.

**Agents:** If `mcp.json` is missing, tell the user to copy from `mcp-example.json` and fill placeholders; do not commit `mcp.json`.

### 3. Reload Cursor

**Command Palette** → **Developer: Reload Window** (or restart Cursor).

### 4. Confirm the MCP server is running

Open **Cursor Settings → MCP** (or your MCP status UI). You should see a server named like **`leetcode`** connected. If it errors, open the MCP log: typical causes are wrong `node` path, missing `npm install`, or blocked network to leetcode.com.

---

## LeetCode account and session cookie (optional)

**Without** a session cookie, you can still use tools such as **`get_problem`**, **`search_problems`**, **`get_daily_challenge`**, **`list_problem_solutions`**, and **`get_problem_solution`** (read-oriented).

**With** a session cookie (`LEETCODE_SESSION`), the same server can expose **authenticated** tools (exact set depends on package version — check the tool list in Cursor). Set it only in your **local** `.cursor/mcp.json` (never in `mcp-example.json`):

```json
"env": {
  "LEETCODE_SITE": "global",
  "LEETCODE_SESSION": "<paste-cookie-value-here>",
  "PATH": "..."
}
```

How to get the cookie: log in on leetcode.com in your browser → DevTools → **Application** → **Cookies** → copy the **`LEETCODE_SESSION`** value. Treat it like a password; **do not commit** it. Prefer local-only overrides or environment injection your team approves.

Reload Cursor after changing `mcp.json`.

---

## End-to-end workflow (typical session)

1. **Open this repo** as the Cursor workspace.
2. **Verify MCP** — `leetcode` server shows as connected.
3. **Pick a problem** — e.g. ask the agent: *“Fetch the problem `two-sum` via LeetCode MCP”* or *“Search easy array problems”*.
4. **Understand and plan** — project rules push a structured flow: clarify constraints, name the pattern, sketch approach, then code. Expect the agent to ask about edge cases before diving into implementation.
5. **Implement** — create or edit a file under `solutions/<language>/` using the naming rule **`NNNN-slug.ext`** (see [solutions/README.md](solutions/README.md)).
6. **Get coaching** — *“Give me a hint”*, *“Review my solution”*, *“What pattern is this?”* The **leetcode-coach** skill uses progressive hints (category → approach → detail → pseudocode) before a full solution unless you ask otherwise.
7. **Interview practice** — after you have a solution, *“Interview me on this”* or *“Ask follow-ups”* triggers **interview-simulator** (complexity, edge cases, optimizations, extensions).
8. **Submit / run on LeetCode** — if your MCP tool list includes run/submit and you configured `LEETCODE_SESSION`, ask the agent to use those tools; otherwise paste your code into the LeetCode site (still a valid end-to-end loop).
9. **Track progress** — update **`progress.md`** and let the agent refresh **`.cursor/MEMORY.md`** after substantive sessions (weak areas, recurring mistakes, preferences).

---

## Important features and expected behavior

### LeetCode MCP (local server)

- Cursor **starts and stops** the server; you do not run a separate terminal daemon for normal use.
- The server uses LeetCode’s **official-ish HTTP APIs**; behavior can change if LeetCode changes their site.
- **Tool names** appear in Cursor’s MCP panel. For `@jinzcdev/leetcode-mcp-server` v1.2.x, typical read tools include: **`get_problem`**, **`search_problems`**, **`get_daily_challenge`**, **`list_problem_solutions`**, **`get_problem_solution`**, plus several **user** tools that may require auth.

### Rules (automatic)

- **Problem-solving discipline** — structured steps from understanding through complexity analysis; hints before spoilers when you’re stuck.
- **Memory keeper** — after non-trivial work, the agent should update `.cursor/MEMORY.md` with durable, actionable notes (gitignored).
- **Skill learning** — mistakes that affect skills can be recorded under `.cursor/skills/*/issues/` and summarized in `LESSONS.md`.
- **Banned tools** — agents must **not** run `claude-code`, `@anthropic-ai/claude-code`, or other disallowed CLI patterns; follow your org’s sanctioned AI tooling policy.

### Skills (on demand)

| Skill | When to use | Expected behavior |
|-------|-------------|-------------------|
| **leetcode-coach** | Hints, reviews, patterns, complexity | Hints escalate in levels; reads `MEMORY.md` and skill `LESSONS.md` when relevant. |
| **interview-simulator** | Post-solve debrief, “real interview” feel | Follow-up questions (complexity, edge cases, variants); structured feedback; may write patterns to `MEMORY.md`. |

---

## Repository layout

```
solutions/                 # One folder per language; files: NNNN-slug.ext
  python/
  golang/
.cursor/
  mcp-example.json         # Committed template — copy to mcp.json locally
  mcp.json                 # Local only (gitignored) — real MCP config
  rules/                   # Workflow + memory + safety rules
  skills/                  # leetcode-coach, interview-simulator
  MEMORY.md                # Local only (gitignored) — learning memory
progress.md                # Committed progress table
package.json               # Pins @jinzcdev/leetcode-mcp-server (run npm install)
README.md                  # This file
AGENTS.md                  # Agent-facing contract (read order, conventions)
```

---

## Adding a language

See [solutions/README.md](solutions/README.md).

---

## Troubleshooting

| Symptom | Likely cause | What to try |
|---------|----------------|------------|
| MCP `leetcode` missing or failing | Wrong `node` path, or deps not installed | Run `npm install`; set absolute `command` to `node` in `mcp.json`; reload window. |
| `npx` hangs during MCP start | Corporate npm registry without this package | Use `node` + `build/index.js` as in this README; install with `--registry https://registry.npmjs.org`. |
| Tools work but “user” calls fail | No / expired session | Set `LEETCODE_SESSION` in `env`, reload Cursor. |
| Agent does not use MCP | Server disconnected | Fix MCP first; then explicitly ask to “use the LeetCode MCP tool `get_problem`”. |

---

## License

See [LICENSE](LICENSE).
