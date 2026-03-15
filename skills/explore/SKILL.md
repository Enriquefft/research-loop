---
name: research-loop:explore
description: Use when the user wants to explore a research topic from scratch — before any papers have been read, before any hypothesis exists. Implements the MIT grad student methodology.
---

# Explore Skill

You are about to map the intellectual landscape of a field in 5 phases.
This is the MIT grad student methodology: compress weeks of reading into one session.

## Phase 1 — Gather literature

Run:
```bash
research-loop explore "<topic>" --max-papers 15
```

Or if the binary is not available, use the LLM to gather papers directly:

Ask Claude: "Find the 15 most important recent papers on <topic>. For each: title, year, 1-sentence contribution."

Save the list. You will use it in every subsequent phase.

## Phase 2 — Extract mental models

Ask exactly this question (Anthropic MIT grad student prompt):

> "What are the 5 core mental models that every expert in <topic> shares?
> These are the intuitions and frameworks that take years to develop — not facts, but ways of thinking."

Write the models down. These become the foundation of your hypothesis.md.

## Phase 3 — Find the field's real debates

Ask exactly this question:

> "Show me the 3 places where experts in <topic> fundamentally disagree.
> For each: what is each side's position, and what is each side's strongest argument?"

Write the debates down. Gaps between debate positions are where new research lives.

## Phase 4 — Generate diagnostic questions

Ask exactly this question:

> "Generate 5 questions that would expose whether someone deeply understands <topic>
> versus someone who just memorized facts. For each: what does a good answer look like?"

Answer these questions yourself using the gathered papers. Every wrong answer triggers:
> "Explain why this is wrong and what I'm missing."

## Phase 5 — Carlini scoring

Score the topic on 4 axes (0.0 to 1.0):

| Axis | Question |
|------|----------|
| **Taste** | Would the field be meaningfully different if someone solved this? |
| **Uniqueness** | What can only YOU bring — skills, timing, framing, cross-field transfer? |
| **Impact** | Write the best-case conclusion NOW. Is it more than "X% improvement"? |
| **Feasibility** | Can you test this with a single-GPU experiment in < 1 week? |

Overall = Taste×0.30 + Uniqueness×0.25 + Impact×0.30 + Feasibility×0.15

**If Overall < 0.5: stop. This topic is not worth pursuing. Find a better one.**
**If Overall ≥ 0.5: proceed to `/discover <topic>`.**

## Output

Write a file `.research-loop/explorations/<topic>_<date>.md` with:
- The 15 papers (title, year, 1-sentence contribution)
- The 5 mental models
- The 3 debates
- The 5 diagnostic questions + your answers
- The Carlini score and verdict
