# LeetCode Coach

An AI coaching skill that helps you practice LeetCode problems through progressive hints and pattern recognition, rather than giving away answers.

## What It Does

- Provides **4-level progressive hints** when you're stuck (category, approach, algorithm detail, pseudocode) before showing any solution
- Reviews your code for correctness, efficiency, and style before submission
- Analyzes submissions (both accepted and failed) with complexity discussion
- Identifies which algorithm pattern a problem belongs to
- Adapts coaching depth based on your tracked strengths and weaknesses

## How to Use

Ask naturally:

- *"Help me with two-sum"* — starts the coaching workflow
- *"I'm stuck"* — provides the next hint level
- *"Review my code"* — code review before submission
- *"What pattern is this?"* — pattern identification
- *"Why did my submission fail?"* — post-failure analysis

## How It Works

1. Reads your learning history from `.cursor/MEMORY.md` to personalize coaching
2. Fetches the problem from LeetCode via MCP
3. Walks you through a structured approach (understand, pattern, plan, implement, verify, analyze)
4. Tracks mistakes in `LESSONS.md` for future reference
