---
name: research-loop:idea-selection
description: Use before committing to any research direction. Applies Carlini's methodology to evaluate whether an idea is worth pursuing. Gate: if it doesn't pass, don't run experiments.
---

# Idea Selection Skill

> "The single most important skill is good taste in what problems are worth solving."
> — Nicholas Carlini, "How to win a best paper award"

## The core question — answer this first

> **"What is the most important open problem in my field right now — and is this it?"**

Write your answer in one sentence. If you cannot, stop and read more papers.

## Gate 1 — Taste (weight 0.30)

Answer each in one sentence:
1. Would the field be meaningfully different if this were solved?
2. Is this on the critical path, or a footnote to something that matters?
3. Do you read existing work and want to scream at the wrong approach?

Score: 0.0 (no to all) → 1.0 (yes to all)

## Gate 2 — Uniqueness (weight 0.25)

What can ONLY YOU bring to this?

| Axis | Question |
|------|----------|
| Skills | Do you have a combination others lack? |
| Timing | Are you 3–12 months ahead? |
| Framing | Do you see the problem structured differently? |
| Cross-field | Does a tool from another field apply here? |

Score: 0.0 (no advantage) → 1.0 (strong advantage on ≥2 axes)

## Gate 3 — Impact (weight 0.30)

Write the best-case conclusion **now**, before running any experiment:

> "If this paper succeeds wildly — all experiments turn out exactly as I want — what would the conclusion say?"

A high-impact conclusion sounds like:
- "This changes how the field should frame [problem]."
- "This shows that [widely held assumption] is wrong."
- "This connects [field A] and [field B] in a way that opens a new direction."
- "This makes [capability] accessible on modest hardware for the first time."

**"Our method achieves X% improvement" is NOT high impact.** Score: 0.0

Score: 0.0 (benchmark improvement only) → 1.0 (changes how the field thinks)

## Gate 4 — Feasibility (weight 0.15)

Can you test the core claim with:
- Single GPU (or equivalent)
- ≤ 1 week of experiments
- A clear metric (val_bpb, accuracy, latency, etc.)

Score: 0.0 (requires months) → 1.0 (testable in days)

## Final score

```
Overall = Taste×0.30 + Uniqueness×0.25 + Impact×0.30 + Feasibility×0.15
```

| Score | Verdict |
|-------|---------|
| ≥ 0.70 | **Pursue** — strong idea, start experiments |
| 0.50–0.69 | **Conditional** — pursue if you can strengthen the weak axis |
| < 0.50 | **Skip** — find a better problem |

## Checklist before proceeding to experiments

- [ ] Best-case conclusion written and is compelling beyond "number goes up"
- [ ] At least one uniqueness axis identified
- [ ] Field critique logged (what is the field doing wrong?)
- [ ] First-principles analysis done independent of the literature
- [ ] Required collaborators identified

**All boxes must be checked before running a single experiment.**
