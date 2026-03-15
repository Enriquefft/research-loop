# SKILL: Idea Selection

> Encode the judgment calls that separate high-impact research from research that is technically correct but that no one reads.
> Based on Carlini (2026) — "How to win a best paper award."

---

## When to use this skill

Use this skill at the **start of any new investigation**, before committing compute, time, or writing effort to a hypothesis. Also use it to periodically re-evaluate ongoing work.

Invoke with: `/idea-selection` or `@epistemic: evaluate this idea`

---

## The core question

Before anything else, ask and answer in writing:

> **"What is the most important open problem in my field right now — and is this it?"**

If you cannot answer this question, stop. Read more papers. Talk to more people. Come back.

---

## Step 1: Evaluate taste

Taste is the ability to recognize, early, whether a direction is worth exploring. It is developed by practicing research and ruthlessly tracking what worked and what did not.

For the idea under evaluation, answer each question in one sentence:

1. **Is this problem worth solving?** Would the field be meaningfully different if someone solved it?
2. **Is this problem on the critical path?** Or is it a footnote to something that matters?
3. **Does this feel "right"?** Not just correct — elegant. Does the approach feel pulled toward a natural solution?
4. **Is anyone doing this badly right now?** Do you read the existing work and want to scream at the wrong approach?

If you answer "no" or "not sure" to more than one of these, the idea may not be worth pursuing. Note why and move on.

---

## Step 2: Check for uniqueness

Ask: **what is only you can bring to this problem?**

| Axis | Question |
|------|----------|
| Skills | Do you have a specific combination of skills (e.g., security + ML, math + systems) others in this area lack? |
| Timing | Are you 3–12 months ahead of where others will get to? |
| Framing | Do you see the problem structured differently from how everyone else frames it? |
| Cross-field transfer | Does a tool or concept from a distant field apply here in a way others haven't seen? |

If you cannot identify at least one axis where you have a comparative advantage, this is a problem others are equally or better positioned to solve. That is not a reason to stop — but it is a reason to reconsider urgency and effort level.

---

## Step 3: Impact filter

Write the best-case conclusion *now*, before running any experiment.

Prompt: "If this paper succeeds wildly — if all experiments turn out exactly as I want — what would I write in the conclusion?"

If the conclusion is only "our method achieves X% improvement on benchmark Y," this paper does not have high impact potential. It may still be worth doing as craft practice — but do not confuse craft practice with a high-impact investigation.

A high-impact conclusion sounds like one of:
- "This changes how the field should frame and evaluate [problem]."
- "This shows that [widely held assumption] is wrong."
- "This connects [field A] and [field B] in a way that opens a new research direction."
- "This makes [important capability] accessible on modest hardware for the first time."

---

## Step 4: Literature check

Read the relevant literature with two goals:

**Goal A — understand what is possible.** For each paper: what is the one sentence that makes this paper useful? If you cannot find it, the paper may not be useful.

**Goal B — identify what is wrong.** Look for:
- Arbitrary early decisions that the field accepted uncritically
- Evaluation metrics that do not measure what they claim to
- Approaches that work empirically but have no principled justification
- Problems where "no one has solved it" is because no one tried a fundamentally different approach

Write these observations in `knowledge_graph.md` under a `field-critique` node before proposing hypotheses.

---

## Step 5: Ignore the literature

After completing Step 4, close every paper and work from first principles.

Ask: "How would I approach this problem if I had never read any of these papers?"

This is not an invitation to reinvent the wheel. It is a guard against inheriting bad framing. If your first-principles answer matches the field's approach, proceed with confidence. If it differs, investigate why before assuming the field is right.

---

## Step 6: Collaborator audit

List the skills this investigation requires. For each required skill, answer:
- Do you have it?
- If not, who does — and can you reach them?

A cold email works when it includes: (1) what you have already done, (2) the specific problem you need help with, (3) why their work is relevant. Do not send emails that say only "I love your work, can we collaborate?"

Document required collaborators in `hypothesis.md` under `## Required Skills`.

---

## Checklist before proceeding

- [ ] The best-case conclusion is written down and is compelling beyond "number goes up"
- [ ] At least one uniqueness axis is identified
- [ ] The field critique is logged in `knowledge_graph.md`
- [ ] You have done first-principles analysis independent of the literature
- [ ] Required collaborators are identified or you have the skills yourself

If all boxes are checked: proceed to the `execution` skill.
If any box is unchecked: do not proceed until it is.

---

## Anti-patterns to avoid

- **Minimum viable contribution mindset.** "What is the least I can do to have a publishable paper?" This produces mediocre work. Reject this framing.
- **Solving problems that do not matter.** Being technically correct is necessary but not sufficient. The question must be worth answering.
- **Following the herd.** If all you're doing is what someone else would have submitted to the same conference, you have contributed nothing beyond timing.
- **Keeping ideas secret.** Ideas are cheap; execution is hard. Share your ideas. The cost of being scooped is far lower than the cost of working in isolation.
