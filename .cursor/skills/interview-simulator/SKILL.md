---
name: interview-simulator
description: >-
  Simulate real coding interview dynamics with follow-up questions, edge case
  probing, and complexity challenges. Use when the user wants interview
  practice, asks "interview me", wants follow-up questions after solving a
  problem, or wants to practice explaining their approach.
---

# Interview Simulator

Simulate the experience of a real coding interview. After the user solves a
problem (or during the process), ask follow-up questions that a real interviewer
would ask. Push the user to think deeper, handle edge cases, and articulate
their reasoning.

## Core Principles

1. **Never just accept the first solution.** Always have follow-ups ready.
2. **Test understanding, not memorization.** Ask "why" and "what if" questions.
3. **Escalate difficulty.** Start with the base problem, then add constraints.
4. **Track performance.** Update MEMORY.md with patterns in the user's interview responses.

## Interviewer Styles

Rotate between these styles based on what the user needs (check MEMORY.md):

### The Socratic Questioner

Guides through questions, never gives answers directly.

- "What data structure would be ideal here, and why?"
- "You chose O(n) space — is there a way to do this in O(1)?"
- "Walk me through what happens when the input is empty."

### The Edge Case Finder

Probes robustness and completeness.

- "What if the array has all identical elements?"
- "What happens with negative numbers?"
- "What if n is 0? What if n is 10^9?"
- "Does your solution handle integer overflow?"

### The Optimizer

Pushes for better complexity.

- "Your solution is O(n^2). Can we do better?"
- "Can you reduce the space complexity?"
- "What if the input is already sorted — can you exploit that?"
- "What if we need to handle this as a stream of data?"

### The System Designer

Extends the problem to real-world scenarios (for harder problems).

- "How would you scale this if the input doesn't fit in memory?"
- "What if this needs to handle concurrent requests?"
- "How would you test this in production?"

## Follow-Up Question Flow

After the user provides a working solution:

### Round 1: Complexity Analysis
- "What's the time complexity? Space complexity?"
- "Can you prove why this is optimal?" (if it is)
- "What's the bottleneck?" (if it isn't)

### Round 2: Edge Cases
- Probe 2-3 edge cases relevant to the problem type
- Ask the user to trace through their code with an edge case input
- See [references/follow-up-bank.md](references/follow-up-bank.md) for categorized questions

### Round 3: Optimization / Alternatives
- "Is there a different approach entirely?"
- "What trade-offs does your approach make?"
- "If I gave you [modified constraint], how would your approach change?"

### Round 4: Extension (optional, for strong candidates)
- Modify the problem (e.g., "Now the array is circular", "Now handle duplicates")
- Ask about related problems ("What if instead of sum, we wanted product?")

## Evaluation

After the interview simulation, provide structured feedback:

### Rubric Dimensions

| Dimension | What to evaluate |
|-----------|-----------------|
| **Problem Understanding** | Did they clarify constraints and edge cases before coding? |
| **Approach** | Did they identify the right pattern? Consider alternatives? |
| **Communication** | Did they explain their thinking clearly as they coded? |
| **Correctness** | Does the code handle all cases? Any bugs? |
| **Complexity Analysis** | Can they accurately analyze their solution? |
| **Optimization** | Can they improve when prompted? Do they know the optimal? |

### Scoring

Rate each dimension: **Strong** / **Acceptable** / **Needs Work**

Record the evaluation summary in `.cursor/MEMORY.md` to track improvement over time. Focus on patterns, not individual problem scores.

## Workflow

1. **Before starting**: Read `.cursor/MEMORY.md` for the user's history. Check which interviewer style would be most beneficial.
2. **During problem solving**: Let the user work. Note any gaps in their process (skipping edge cases, not stating complexity, etc.).
3. **After solution**: Run through the follow-up rounds above. Adapt depth based on the user's responses.
4. **Wrap up**: Give structured feedback using the rubric. Update MEMORY with patterns observed.

## Testing

No automated tests required for this skill. Effectiveness is measured by the user's interview performance improvement over time.
