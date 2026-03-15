---
name: research-loop:loop
description: Use when the user has a hypothesis and wants to run experiments. Drives the PROPOSE → MUTATE → BENCHMARK → ANNOTATE cycle against a baseline repo.
---

# Loop Skill

You are running the experiment loop. This is the Empirical Agent doing science.

## Setup checklist

Before starting the loop:
- [ ] `hypothesis.md` exists and is approved
- [ ] Baseline repo is cloned (karpathy/autoresearch or equivalent)
- [ ] `benchmark_command` is set in `.research-loop/config.toml`
- [ ] You have verified the baseline runs: `uv run train.py > run.log 2>&1`

If any are unchecked → fix them before proceeding.

## Start the loop

```bash
research-loop loop start --repo ./autoresearch --max-runs 20
```

For karpathy/autoresearch specifically:
```bash
# One-time setup
cd autoresearch && uv run prepare.py

# Set in .research-loop/config.toml:
# benchmark_command = "uv run train.py > run.log 2>&1"
# metric.name = "val_bpb"
# metric.direction = "lower"
# metric.timeout = 420

research-loop loop start --repo ./autoresearch
```

## What the loop does each iteration

```
PROPOSE   → LLM proposes a specific mutation to the repo (e.g. change DEPTH from 4 to 6)
MUTATE    → Applies the mutation as a git diff
BENCHMARK → Runs benchmark_command, reads val_bpb from output
ANNOTATE  → LLM writes causal annotation: why did this work/fail?
```

Each iteration appends to:
- `autoresearch.jsonl` — machine-readable record
- `knowledge_graph.md` — living DAG of explored paths
- `lab_notebook.md` — human-readable log

## Reading the metric output

For karpathy/autoresearch, the metric appears as:
```
---
val_bpb:  0.997900
```

Lower = better. The loop tracks `bestVal` and reports Δ from baseline.

## When to stop the loop

Stop when:
1. `bestVal` has not improved in 5+ consecutive runs → diminishing returns
2. The knowledge graph shows a clear pattern (what works, what doesn't)
3. You have enough data to write a clear claim + evidence

**Never run more than 50 iterations without reviewing the knowledge graph.**

## After the loop

Run `research-loop:execution` skill to review results and decide: continue, pivot, or kill.

Write the conclusion paragraph **before** looking at whether you got the result you wanted.
Then check if it matches. If it doesn't, investigate why — that gap is often the real finding.
