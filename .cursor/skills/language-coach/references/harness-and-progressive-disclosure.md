# Harness thinking and progressive disclosure (for this skill)

## Agent harness (engineering sense)

In agent systems, a **harness** is everything wrapped around the model to make behavior reliable: fixed **read order**, explicit **state** (what is already true in files), **guardrails** (when not to load full context), and **update hooks** (when to persist new facts). This skill treats `learning/<language>/TRACKER.md` plus per-concept files as **durable harness state**: the agent must consult them before teaching so it does not duplicate deep work or contradict recorded mastery.

**Implications for the agent**

- Treat the tracker as the **routing table**: shallow pass first, full concept file only when the question needs it or depth is being upgraded.
- Treat each concept file as **append-friendly truth**: add detail in dated sections rather than rewriting history, so "what was covered when" stays auditable.
- **Do not** re-expand topics already at `deep` unless the user asks for review, correction, or explicitly new material; then extend the file instead of replacing it.

## Progressive disclosure (learning / UX sense)

**Progressive disclosure** means presenting information in layers: overview → relevant chunk → full detail on demand. Here:

| Layer | Artifact | When the agent uses it |
|-------|----------|-------------------------|
| 0 | Skill `SKILL.md` | Workflow only; no personal progress |
| 1 | `learning/<lang>/TRACKER.md` | Every language-coach turn for that language |
| 2 | Concept file **At a glance** + **Gaps** | Quick questions, "what do I know?" |
| 3 | Rest of concept file | Deeper questions, interview prep, syntax sugar breakdowns |

**Anti-pattern**: Dumping a full tutorial when the tracker shows `deep` and the user asked a narrow follow-up.

## Depth levels (contract)

| Depth | Meaning | Agent behavior |
|-------|---------|----------------|
| `none` | Listed in tracker only, no concept file yet | Create concept file from template on first real touch |
| `glance` | Skim / orientation only | Short answers; offer to deepen; record bullets under **At a glance** |
| `standard` | Workable understanding, some gaps | Answer with references to file sections; grow **Detailed notes** |
| `deep` | Interview-ready for this concept | No full re-teach by default; recap, stress-test, or add **Addenda** |

Upgrades must update YAML `depth` and `last_session` and a line in the tracker table.
