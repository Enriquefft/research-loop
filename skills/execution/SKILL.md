---
name: research-loop:execution
description: Use when experiments are running or just completed. Annotate results, update knowledge graph, decide whether to continue/pivot/kill.
---

# Execution Skill

You are the Empirical Agent reviewing experimental results and deciding next steps.

## After each experiment run

Read the latest entry in `autoresearch.jsonl`:
```bash
tail -1 .research-loop/sessions/<id>/autoresearch.jsonl | python3 -m json.tool
```

Answer these three questions:
1. **What happened?** (metric value, direction, magnitude of change)
2. **Why did it happen?** (causal, not correlational — what mechanism explains it?)
3. **What should happen next?** (continue same direction / pivot / kill)

Write the answer to (2) as a causal annotation in `knowledge_graph.md`.

## The kill/pivot/continue decision

At every checkpoint, apply this decision tree:

```
Has the metric improved in the last 5 runs?
├── YES → continue along this direction
└── NO
    ├── Are we > 10 runs in with no improvement?
    │   └── YES → KILL this lane. Move to next hypothesis.
    └── NO → PIVOT: try a different mutation direction
```

**Do not run more than 10 iterations without improvement. Sunk cost is not a reason to continue.**

## Updating the knowledge graph

Every run should add a node to `knowledge_graph.md`:

```markdown
## Run #N — <node name>

- **Mutation**: what changed
- **Result**: metric value (Δ from baseline)
- **Annotation**: WHY this worked or failed (mechanistic)
- **Next**: what this implies for the next run
```

## When to declare success

Declare success when:
- Best metric is a meaningful improvement from baseline (not noise)
- You can explain WHY the improvement happened (not just that it did)
- The effect is reproducible (run it twice, same result)
- You have enough failed runs to know what DOESN'T work

## Writing the paper from experiments

Do not start writing until:
- [ ] Knowledge graph has ≥ 5 annotated runs
- [ ] You can state the core claim in one sentence
- [ ] You have at least one clear negative result (what didn't work and why)
- [ ] The best-case conclusion from idea-selection matches what the experiments found

Then invoke `research-loop:writing-papers` (coming soon).
