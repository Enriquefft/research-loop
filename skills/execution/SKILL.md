# SKILL: Execution

> How to run a research investigation with discipline: de-risking fast, killing dead projects ruthlessly, and putting in unreasonable effort on the ones that matter.
> Based on Carlini (2026) — "How to win a best paper award."

---

## When to use this skill

Use this skill **after** `idea-selection` has cleared an investigation. Apply it continuously throughout the experiment loop — at every decision point about what to run next, what to cut, and how much effort to apply.

Invoke with: `/execution` or `@empirical: am I running this right?`

---

## Principle: de-risk first

The most dangerous thing you can do is spend weeks on the part of the project you already know how to solve.

**Rule: always start with the sub-problem most likely to fail.**

Before writing any polished code, answer: "What is the one thing that, if it doesn't work, kills this entire investigation?" Run that experiment first. A small, dirty prototype that answers the core question is worth more than a clean implementation of something peripheral.

Apply to the experiment loop:
1. Identify the assumption the entire hypothesis rests on
2. Design the minimal experiment that tests that assumption
3. Run it
4. If it fails: log to `knowledge_graph.md`, do not polish, move on
5. If it succeeds: clean it up, run the full benchmark, log the result

---

## Step 1: Kill papers that are not working

This is the hardest step and the most important.

At each checkpoint, ask:

> "If I spent the next two weeks finishing this investigation, would it produce something compelling?"

If the honest answer is "probably not," stop now. Do not finish it "just to have something to show." The time cost of finishing a bad paper is the same as the opportunity cost of not starting a good one.

Log the investigation to `knowledge_graph.md` as `status: killed` with a causal annotation explaining what was tried and why it was abandoned. This is not failure — it is the knowledge graph doing its job.

**Acceptable reasons to kill a paper:**
- The core assumption turned out to be false
- The results are technically correct but the contribution is not compelling
- A stronger paper on the same topic appeared while you were working on it
- The expected conclusion (from the `idea-selection` skill) is no longer achievable

**Unacceptable reasons to continue a bad paper:**
- "I've already spent three weeks on it" — sunk cost fallacy
- "I can probably still get it accepted somewhere" — this is not the goal
- "I don't have a better idea right now" — start brainstorming; do not finish bad work

---

## Step 2: Reprioritize ruthlessly when something better arrives

When a new idea arrives that is clearly more important than the current investigation:

1. Write the best-case conclusion for the new idea (from `idea-selection`)
2. Write the best-case conclusion for the current investigation
3. Compare honestly
4. If the new idea is substantially more important: stop the current investigation, log its state to `knowledge_graph.md`, and pivot

The impact distribution of research papers is exponential. A paper that is 100x more important is worth pivoting to even if you are 90% done with the current one. The marginal return on finishing a mediocre paper is almost always lower than starting on something that matters.

Document the pivot decision in `lab_notebook.md` with the date and reasoning, so future sessions can understand why the investigation was paused.

---

## Step 3: Apply unreasonable effort to investigations that matter

Once an investigation has cleared the kill/pivot decision and is worth pursuing, apply effort that would seem disproportionate to an outside observer.

Concretely:
- Run every experiment multiple times with different seeds
- Control for every confounder a skeptical reviewer would raise
- If there is a question that "could" be answered with more data or a longer run, answer it before someone asks
- Spend hours strengthening "sometimes" to "usually" when it matters for the claim
- Do not stop when the experiment is "good enough" — stop when it is as clean as it can be

This is the wildlife photographer principle: you need to be skilled and in the right place, but the great picture comes from waiting an unreasonable amount of time for exactly the right circumstances.

---

## Step 4: Maintain singular focus

At every point in the experiment loop, the investigation should be advancing exactly one idea.

Before running any experiment, answer:
- "Does this experiment directly support the core claim of this investigation?"
- "Could I cut this experiment without weakening the paper?"

If the answer to the second question is "yes," cut it.

Write the core claim at the top of `hypothesis.md` in one sentence. Every proposed mutation, every benchmark, every new analysis must trace back to that sentence. If it does not, either the experiment is off-topic or the hypothesis needs to be revised.

---

## Step 5: The paper must be at a local optimum

Before declaring an investigation complete, perform a gap audit:

Ask: "Is there an obvious experiment I did not run that a reader will immediately notice is missing?"

If yes, run it. A paper should leave the reader satisfied — not feeling that something essential was missing. The test: would a skeptical reader, finishing your paper, say "I wish they had done X"? If you can anticipate X, do X.

The distinction between "local optimum" and "global optimum": you do not need to run every possible experiment. You need to run every experiment that is obviously implied by your claim. Leaving room for follow-up work is fine and good — it invites others to engage with your ideas. Leaving an obvious gap is not fine.

---

## Experiment loop integration

In the Research Loop experiment cycle, this skill maps to:

| Loop state | Execution skill action |
|------------|----------------------|
| `HYPOTHESIZE` | Apply Step 4 (focus check) — does this mutation advance the core claim? |
| `MUTATE → BENCHMARK` | Apply Step 3 (unreasonable effort) — run with multiple seeds, control confounders |
| `ANNOTATE` | Apply Step 1 kill check — is this investigation still worth continuing? |
| `PROPOSE` | Apply Step 2 pivot check — has something more important arrived? |

Log all kill and pivot decisions to `autoresearch.jsonl` with `status: killed` or `status: pivoted` so the loop can resume from a clean state.

---

## Checklist at each loop checkpoint

- [ ] Core claim is still written in one sentence at the top of `hypothesis.md`
- [ ] Every experiment in this loop directly supports that claim
- [ ] The sub-problem most likely to fail was tested first
- [ ] No sunk cost is being honored — if the investigation is not working, it is logged and killed
- [ ] If a more important idea arrived, a pivot decision was documented in `lab_notebook.md`
- [ ] Every obvious gap in the experimental record is either filled or explicitly deferred with a reason

---

## Anti-patterns to avoid

- **Polishing before de-risking.** Writing clean code before you know the core idea works.
- **Sunk cost continuation.** Finishing a paper because you started it, not because it matters.
- **Experiment sprawl.** Running every possible experiment because it "might" be interesting. Focus.
- **Perfectionism on the wrong thing.** Spending days on a secondary result while the primary claim is still shaky.
- **Ignoring the obvious gap.** Finishing a paper while knowing a critical experiment is missing and hoping no reviewer notices.
