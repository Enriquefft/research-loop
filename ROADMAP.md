# Roadmap

Research Loop is an open-source Agent OS for scientific researchers. This roadmap reflects the PRD v0.2 vision. Items marked **shipped** are in the current binary.

---

## Phase 1 — Foundation (v0.1, shipped)

The core loop: ingest a paper, propose mutations, run benchmarks, log results.

- [x] Go binary with Bubble Tea TUI (Chat pane)
- [x] Epistemic + Empirical agents with OpenAI-compatible backend
- [x] Paper ingestion pipeline (ArXiv URL → `hypothesis.md`)
- [x] Experiment loop state machine with JSONL persistence
- [x] Knowledge graph (`knowledge_graph.md` Markdown DAG)
- [x] `.research` bundle export/resume
- [x] `research-loop start <url>` magic moment
- [x] karpathy/autoresearch baseline integration
- [x] Claude Code skill system (explore, discover, loop, idea-selection, execution)
- [x] SessionStart hook (active session state injected on every open)
- [x] Conversational advisor mode (research-loop skill)
- [x] Parallel discovery orchestrator with Carlini decision gates
- [x] `learn` skill — MIT grad student methodology + Socratic reverse learning

---

## Phase 2 — Full Workspace (v0.2, in progress)

Target: Q2 2026

### TUI

- [ ] Dashboard pane (F2): live experiment metrics, color-coded results, cost tracker, run history
- [ ] Library pane (F3): paper browser with full-text search, reading annotations, citation graph
- [ ] Writer pane (F4): Markdown editor with Vim keybindings, LaTeX preview, Scribe integration

### Hands system

- [ ] Hand lifecycle management (activate / pause / resume / schedule)
- [ ] `HAND.toml` manifest format + guardrails
- [ ] Ingestor Hand — scheduled PDF ingestion from ArXiv RSS feeds
- [ ] Experimenter Hand — overnight experiment loops with approval gates
- [ ] Librarian Hand — RSS feed monitoring and auto-ingestion
- [ ] Scribe Hand — paper section drafting from experiment data
- [ ] Reviewer Hand — methodology checking, p-hacking detection
- [ ] Replicator Hand — automated reproduction of published results
- [ ] Deep Learner Hand — 5-phase corpus pipeline (see below)

### Deep Learner Hand

The flagship feature for rapid field mastery:

- [ ] Phase 1: Corpus assembly — accept textbooks, papers, transcripts, blog posts
- [ ] Phase 2: Mental model extraction → `mental_models.md`
- [ ] Phase 3: Intellectual landscape mapping (debates, consensus, open questions) → `landscape.md`
- [ ] Phase 4: Deep understanding question generation → `deep_questions.md`
- [ ] Phase 5: Socratic tutoring loop — answer evaluation against full corpus → `learning_notebook.md`
- [ ] Auto-trigger when 3+ papers on same topic accumulate in library
- [ ] Feed outputs into Epistemic Agent and knowledge graph

### Paper ingestion improvements

- [ ] PDF parser pipeline: full-text extraction with mathematical notation support
- [ ] Fallback to abstract-only with notification when full-text parsing fails
- [ ] BibTeX entry support
- [ ] DOI resolution
- [ ] Structured extraction: title, authors, abstract, core claim, math, baselines → `hypothesis.md`

### Core improvements

- [ ] Semantic similarity check to skip duplicate hypotheses
- [ ] Hardware auto-detection (CUDA / Apple MPS / CPU) with benchmark suggestions
- [ ] Cost tracking: real-time LLM API spend + GPU-hours in Dashboard
- [ ] Budget cap per session with auto-pause
- [ ] Backpressure correctness checks (`research.checks.sh`)
- [ ] Multi-paper sessions: chain papers for cross-hypothesis synthesis
- [ ] ArXiv RSS feed subscriptions

---

## Phase 3 — Collaboration & Distribution (v0.3)

Target: Q3 2026

- [ ] MCP bridge server — expose ingestion, loop control, knowledge graph, library as MCP tools/resources
- [ ] Bundle registry — `research-loop publish` / `research-loop search`
- [ ] Collaborator bundle mode (read/write/run) vs. read-only reviewer mode
- [ ] Replicator Hand — pass/fail replication reports per claim
- [ ] LaTeX/PDF export from Writer pane (via pandoc)
- [ ] Knowledge graph Mermaid export
- [ ] Cross-session insight synthesis
- [ ] Community Hand registry (REST API or GitHub-backed index)

---

## Phase 4 — Production (v1.0)

Target: Q4 2026

- [ ] Replay mode for `.research` bundles (step through experiment history)
- [ ] Homebrew formula, npm package, `curl | sh` installer
- [ ] Documentation site (research-loop.dev)
- [ ] Performance optimization: cold start < 500ms, TUI @ 60fps
- [ ] SQLite FTS5 index for libraries > 1000 papers
- [ ] Statistical significance testing integration
- [ ] Multi-session parallel experiment management
- [ ] Automated paper citing Research Loop in methodology sections

---

## Long-term / Speculative

These are directions we find interesting but have not committed to:

- Voice mode (research conversations via microphone)
- Web UI companion for mobile review of lab notebooks
- Federated bundle sharing (peer-to-peer, no central registry)
- Agent-to-agent collaboration across institutions
- Integration with Jupyter for hybrid notebook/loop workflows
- Fine-tuned domain-specific Epistemic Agents per field

---

## What We Are Not Building

- A GUI or web dashboard (terminal-first, always)
- A cloud service or SaaS product
- A fork of Claude Code or OpenCode
- Multi-node distributed training support
- Automatic institutional login for paywalled papers

---

## How to Influence the Roadmap

Open an issue with the label `roadmap` and describe:
1. The user problem you're trying to solve
2. How you currently work around it
3. What success would look like

We prioritize based on researcher impact and implementation feasibility.
