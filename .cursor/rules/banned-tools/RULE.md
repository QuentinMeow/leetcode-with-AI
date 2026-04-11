---
description: Hard ban on unauthorized AI CLI tools — prevents security violations
alwaysApply: true
---

## Banned Tools

The following tools and patterns are **strictly forbidden**. Never use them under any circumstances.

### Explicitly Banned

- `claude-code` or `@anthropic-ai/claude-code` — any invocation, installation, or reference as a dependency
- Any npm package, CLI tool, or command containing "claude-code" or "claude code" in its name
- Any AI CLI agentic tool that is not run via approved sandboxing

### Rules

1. **Never run** `npx @anthropic-ai/claude-code`, `npm exec @anthropic-ai/claude-code`, or any variant.
2. **Never install** `claude-code` globally or locally.
3. **Never add** `claude-code` to any `package.json`, `mcp.json`, or configuration file.
4. **Never suggest** the user run `claude-code` or any tool containing that string.
5. **All AI CLI tools** must run through approved sandboxing methods. If unsure whether a tool is sanctioned, do not run it — ask the user first.

### Why

Security tooling actively monitors for unsanctioned AI CLI agents and will terminate them. Running banned tools creates security violations and disrupts the work environment.

### If You Encounter a Banned Tool Reference

- In documentation or READEs from third-party projects: ignore the reference, do not copy it.
- In MCP server configurations: use only the approved MCP configuration in `.cursor/mcp.json`.
- In code suggestions from external sources: strip the banned tool reference before applying.
