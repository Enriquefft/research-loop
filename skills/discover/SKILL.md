---
name: research-loop:discover
description: Use when the user wants to run parallel research lanes on a topic that has already passed the Carlini scoring gate. Spawns multiple hypothesis angles and kills poor ones early.
---

# Discover Skill

You are running a parallel research discovery. Multiple independent lanes pursue different angles of the same topic simultaneously. Carlini gates kill weak lanes before they waste compute.

## Precondition

Before running discovery, confirm:
- [ ] Exploration is complete (mental models and debates are documented)
- [ ] Carlini score ≥ 0.5
- [ ] You have identified at least one uniqueness axis

If any are unchecked → go back to `research-loop:explore`.

## Run discovery

```bash
research-loop discover "<topic>" --max-lanes 4
```

Each lane runs this pipeline:
```
LITERATURE → GAP_ANALYSIS → HYPOTHESIS → EXPERIMENT → BENCHMARK → REVIEW
```

At each transition, a Carlini gate fires. If a lane scores below the threshold, it is **killed immediately** — no wasted compute.

## What to watch for

While lanes run, monitor `.research-loop/discoveries/` for the output JSON. Look for:

1. **Which lanes survived all gates?** → These are your best hypotheses.
2. **Which lanes were killed and why?** → Read the kill reason. It tells you what NOT to pursue.
3. **Best metric across surviving lanes** → This is your new baseline.

## After discovery completes

Read the lane results:
```bash
cat .research-loop/discoveries/<topic>_<date>.json
```

For each surviving lane (verdict = "promising"):
1. Read the `claim` field — this is your hypothesis
2. Read the `experiment` field — this is what to test
3. Read the `review` field — senior researcher's assessment

Pick the lane with:
- Highest overall score
- Most distinct claim (not just "more of the same")
- Feasible experiment (modifies specific constants in train.py)

## Proceed to loop

Once you have a promising lane:
```bash
research-loop loop start --repo ./autoresearch
```

Or invoke `research-loop:loop` skill for the full loop protocol.
