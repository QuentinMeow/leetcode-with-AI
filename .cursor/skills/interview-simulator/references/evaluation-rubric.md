# Interview Evaluation Rubric

Structured criteria for evaluating interview simulation performance.

## Dimensions

### 1. Problem Understanding

| Rating | Criteria |
|--------|----------|
| **Strong** | Clarified constraints, identified edge cases, confirmed understanding with examples before coding |
| **Acceptable** | Understood the problem correctly, asked some clarifying questions |
| **Needs Work** | Started coding immediately, missed constraints, or misunderstood the problem |

### 2. Approach Selection

| Rating | Criteria |
|--------|----------|
| **Strong** | Identified the optimal pattern, considered alternatives, explained trade-offs |
| **Acceptable** | Found a working approach, but didn't explore alternatives |
| **Needs Work** | Chose a suboptimal approach without recognizing it, or couldn't find an approach |

### 3. Communication

| Rating | Criteria |
|--------|----------|
| **Strong** | Explained thinking clearly throughout, narrated decisions, caught own mistakes verbally |
| **Acceptable** | Explained approach at the start, but went silent while coding |
| **Needs Work** | Coded silently, couldn't explain decisions when asked |

### 4. Correctness

| Rating | Criteria |
|--------|----------|
| **Strong** | Code works for all cases including edge cases, no bugs |
| **Acceptable** | Code works for main cases, minor edge case issues caught during review |
| **Needs Work** | Code has logical errors, doesn't compile, or fails on basic cases |

### 5. Complexity Analysis

| Rating | Criteria |
|--------|----------|
| **Strong** | Correctly analyzed both time and space, identified the bottleneck, knows if it's optimal |
| **Acceptable** | Got the right complexity but couldn't fully justify it |
| **Needs Work** | Incorrect analysis, confused about complexity, or couldn't analyze at all |

### 6. Optimization Response

| Rating | Criteria |
|--------|----------|
| **Strong** | Successfully optimized when challenged, or correctly argued current solution is optimal |
| **Acceptable** | Had ideas for optimization but couldn't fully implement |
| **Needs Work** | Couldn't think of any improvements when prompted |

## Overall Assessment Template

```
## Interview Assessment: [Problem Name]

**Date**: [date]
**Difficulty**: [Easy/Medium/Hard]
**Pattern**: [algorithm pattern]

| Dimension | Rating |
|-----------|--------|
| Problem Understanding | [Strong/Acceptable/Needs Work] |
| Approach | [Strong/Acceptable/Needs Work] |
| Communication | [Strong/Acceptable/Needs Work] |
| Correctness | [Strong/Acceptable/Needs Work] |
| Complexity Analysis | [Strong/Acceptable/Needs Work] |
| Optimization | [Strong/Acceptable/Needs Work] |

**Key Strengths**: [what went well]
**Areas to Improve**: [specific, actionable improvements]
**Recommended Practice**: [specific problem types or skills to work on]
```

## Recording in MEMORY.md

After assessment, update `.cursor/MEMORY.md` with:

- Move persistent patterns (e.g., "consistently struggles with complexity analysis") to **Weak Areas**
- Move demonstrated improvements to **Strengths**
- Keep individual problem scores out — MEMORY tracks patterns, not events
