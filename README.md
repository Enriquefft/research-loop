

**[Quickstart](#quickstart)** · **[Docs](https://research-loop.dev/docs)** · **[GitHub](https://github.com/research-loop/research-loop)** · **[Discord](https://discord.gg/research-loop)**



  


## What is Research Loop?

# Open-source Agent OS for scientific researchers

**If Claude Code is a *coding agent*, Research Loop is the *research environment***

Research Loop is a standalone terminal application that orchestrates a team of AI agents to run the full scientific discovery cycle — from reading a paper to running reproducible experiments, building a living knowledge graph, and drafting your results.

It looks like a terminal workspace — but under the hood it has dual agents, experiment loop state machines, persistent knowledge graphs, and portable `.research` bundles.

**Manage hypotheses, not terminals.**


|        | Step              | Example                                                              |
| ------ | ----------------- | -------------------------------------------------------------------- |
| **01** | Point at a paper  | `research-loop start "https://arxiv.org/abs/2403.xxxxx"`             |
| **02** | Review hypothesis | The Epistemic agent extracts the core claim. You approve or edit it. |
| **03** | Run the loop      | The Empirical agent mutates code, benchmarks, annotates — overnight. |


  


> **COMING SOON: Hand Registry** — Download and run entire research workflows with one command. Browse pre-built Hand definitions — Ingestor, Experimenter, Replicator, Deep Learner — and import them into your workspace in seconds.

  



|                    |        |         |        |        |       |     |
| ------------------ | ------ | ------- | ------ | ------ | ----- | --- |
| **Works** **with** | Claude | GPT-4/5 | Gemini | Ollama | ArXiv | MCP |


*If it outputs a metric, it can be benchmarked.*

## Research Loop is right for you if

- ✅ You read ArXiv papers and spend **days translating them into runnable code**
- ✅ You run **20+ experiments** and lose track of what you've tried and why things failed
- ✅ You want experiments running **autonomously overnight**, but still want full scientific control
- ✅ You want a **living knowledge graph** that remembers what failed and why — so you never repeat it
- ✅ You want to **share reproducible experiments** as a portable bundle, not a pile of scripts
- ✅ You're on modest hardware — a **MacBook or single consumer GPU** — not an H100

  


## Features


|                                                                                                                                                                                           |                                                                                                                                                                                                        |                                                                                                                                                                                                                     |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 🧪 Dual-Agent Loop Epistemic agent reads theory and proposes hypotheses. Empirical agent writes code, runs benchmarks, and logs results. They run together, overnight, on your hardware. | 🧠 Living Knowledge Graph Every hypothesis tried, every result observed, every causal annotation for every failure — in a single human-readable Markdown file. Any agent or human can resume from it. | 📦 .research Bundles The portable unit of reproducible science. Everything needed to reproduce or continue an investigation, in a single ZIP-compatible archive. Like `.ipynb` — but for the full discovery cycle. |
| 🖐️ Hands Autonomous capability packages that run on schedules. Ingestor, Experimenter, Librarian, Scribe, Reviewer, Replicator, Deep Learner. Build your own with a `HAND.toml`.        | 📚 Paper Library Ingest ArXiv URLs, DOIs, or local PDFs. Full-text search across your entire library. ArXiv RSS feed monitoring. Citation graph. Cross-session paper linking.                         | ✍️ Writer Pane Draft paper sections from your experiment data. Scribe agent auto-populates Methods, Results, Related Work. Vim keybindings. LaTeX export.                                                          |
| 💰 Cost Control Real-time LLM API spend in the dashboard. Per-session budget caps. Token tracking per agent. No runaway costs.                                                           | 🔌 Model-Agnostic Claude, GPT-4/5, Gemini, Ollama, LM Studio, OpenRouter, Groq, DeepSeek — any OpenAI-compatible endpoint. Configure per-agent.                                                       | 🔬 Deep Learner Feed it a topic and your entire source corpus. It extracts expert mental models, maps field debates, generates deep-understanding questions, and tutors you through them.                          |


  


## Problems Research Loop solves


| Without Research Loop                                                                                                 | With Research Loop                                                                                             |
| --------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| ❌ You read a promising paper and spend 2–3 days translating it into runnable code before running a single experiment. | ✅ `research-loop start <arxiv-url>` produces a running baseline in under 3 minutes.                            |
| ❌ You run 40 experiments and can't remember which variants you tried or why the promising-looking one failed.         | ✅ The knowledge graph logs every hypothesis, result, and causal annotation. Nothing is forgotten.              |
| ❌ You rerun experiments you already ran 3 weeks ago because you forgot the outcome.                                   | ✅ Duplicate detection against the knowledge graph prevents redundant attempts automatically.                   |
| ❌ You share results with a collaborator by emailing a ZIP of scripts and hoping they can reproduce it.                | ✅ `research-loop export` produces a `.research` bundle that any researcher or agent can resume on any machine. |
| ❌ Runaway experiment loops cost you hours of GPU time on dead-end mutations with no way to trace back why.            | ✅ Backpressure checks auto-revert failing mutations. Cost tracking enforces session budgets.                   |
| ❌ You have to be at your computer babysitting experiments instead of thinking about science.                          | ✅ The Experimenter Hand runs loops autonomously on a schedule. You review results when you're ready.           |


  


## Why Research Loop is different


|                                        |                                                                                                                                                 |
| -------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| **Checkpoint-based recovery.**         | Every loop state transition is persisted to `autoresearch.jsonl`. Crash, reboot, context reset — the loop resumes in under 10 seconds.          |
| **Causal annotations, not just logs.** | The Epistemic agent writes *why* a hypothesis failed, not just that it did. The knowledge graph is a causal map, not a run log.                 |
| **Domain-agnostic metrics.**           | Any CLI metric works: `val_bpb`, `val_loss`, NDCG@10, Sharpe ratio, Lighthouse score. Declare direction, parse `METRIC name=value` from stdout. |
| **Plain-text persistence.**            | All state lives in Markdown and JSONL — human-readable, git-diffable, LLM-resumable. No databases. No proprietary formats.                      |
| **Hands, not scripts.**                | Each autonomous capability is a self-contained Hand with a manifest, system prompt, tools, and guardrails. Composable, pausable, shareable.     |
| **MCP bridge included.**               | Don't want to switch tools? The MCP server exposes Research Loop capabilities to Claude Code, OpenCode, Cursor, or any MCP host.                |


  


## What Research Loop is not


|                                      |                                                                                                                                                  |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Not an AGI research system.**      | The tool proposes. The researcher approves. The agent executes. Scientific judgment stays with you.                                              |
| **Not a cloud product.**             | Local-first. Runs on a MacBook. No account required. No data leaves your machine unless you push a bundle.                                       |
| **Not a coding agent.**              | Research Loop isn't a general-purpose coding assistant with research features bolted on. It's a research environment that happens to write code. |
| **Not a GUI dashboard.**             | v0.1 is strictly CLI and Markdown-first. The TUI is a terminal workspace, not a web app.                                                         |
| **Not a statistical analysis tool.** | Experiment execution, not significance testing. Bring your own stats.                                                                            |


  


## Quickstart

One binary. No database. No account.

```bash
curl -fsSL https://research-loop.dev/install | sh
research-loop init
research-loop start "https://arxiv.org/abs/2403.05821"
```

Or build from source:

```bash
git clone https://github.com/research-loop/research-loop.git
cd research-loop
go build ./cmd/research-loop
./research-loop init
```

> **Requirements:** Go 1.22+ (to build from source) · Any OpenAI-compatible LLM API key

  


## The magic moment

```
$ research-loop start "https://arxiv.org/abs/2403.05821"

Downloading paper... done (2.1s)
Extracting hypothesis... done (8.4s)

┌─────────────────────────────────────────────────────────────┐
│ Hypothesis extracted                                        │
│                                                             │
│ Paper: GQA: Training Generalized Multi-Query Transformer    │
│ Claim: Grouped-query attention reduces memory bandwidth     │
│        during decoding while matching MHA quality with      │
│        1/G of the KV cache memory.                         │
│                                                             │
│ Proposed experiment: Implement GQA on nanoGPT baseline     │
│ and benchmark val_bpb on OpenWebText.                      │
└─────────────────────────────────────────────────────────────┘

Approve this hypothesis and start the loop? [y/N] y

Cloning baseline (nanoGPT)... done
Running baseline... done  →  val_bpb: 3.42

🔬 research-loop | run 1/∞ | baseline: 3.42 bpb | current: 3.38 bpb | Δ -0.04 ✓
```

  


## Terminal UI

Four panes. One terminal. No browser required.


| Pane          | Key  | What it shows                                                           |
| ------------- | ---- | ----------------------------------------------------------------------- |
| **Chat**      | `F1` | Agent conversation. `/slash` commands. Natural language.                |
| **Dashboard** | `F2` | Live experiment metrics, cost tracker, run history, Hand status.        |
| **Library**   | `F3` | Paper browser, full-text search, reading annotations, citation graph.   |
| **Writer**    | `F4` | Markdown editor, Vim keybindings, Scribe auto-population, LaTeX export. |


### Dashboard widget

```
╔══════════════════════════════════════════════════════════════╗
║  🔬 research-loop                                    v0.2.0 ║
║  Session: attention-mechanism-2024                           ║
╠══════════════════════════════════════════════════════════════╣
║  Runs: 47/∞        Baseline: 3.42 bpb     Best: 3.21 bpb   ║
║  Current: 3.28 bpb   Δ best: -0.21 ✓      Δ last: +0.07    ║
║  Status: BENCHMARK   Hand: experimenter    Uptime: 4h 12m   ║
╠══════════════════════════════════════════════════════════════╣
║  Cost: $2.84 API  ·  0.6 GPU-hrs  ·  ~$0.06/run            ║
║  Graph: 47 nodes  ·  12 improvements  ·  3 dead ends        ║
╠══════════════════════════════════════════════════════════════╣
║  Last 5 runs:                                                ║
║  #47  3.28 bpb  +0.07  ▼  lr_cosine_warmup_v2              ║
║  #46  3.21 bpb  -0.03  ▲  dropout_schedule_linear           ║
║  #45  3.24 bpb  -0.01  ▲  head_dim_128_rope                ║
║  #44  3.25 bpb  +0.02  ▼  gqa_4_groups (checks_failed)     ║
║  #43  3.23 bpb  -0.04  ▲  swiglu_activation                ║
╚══════════════════════════════════════════════════════════════╝
```

  


## Hands

**Hands** are autonomous capability packages that run independently, on schedules, without you prompting them.


| Hand             | What it does                                                                                                                                            |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Ingestor**     | ArXiv URL → structured hypothesis in under 30 seconds.                                                                                                  |
| **Experimenter** | The core loop. Proposes mutations, runs benchmarks, annotates failures. Runs in 5-minute cycles.                                                        |
| **Librarian**    | Monitors ArXiv RSS feeds. Indexes, deduplicates, cross-references. Runs daily.                                                                          |
| **Scribe**       | Drafts paper sections from your experiment data. Methods, Results, Related Work, tables, figures.                                                       |
| **Reviewer**     | Internal scientific critic. Checks methodology, flags p-hacking, validates statistical claims.                                                          |
| **Replicator**   | Attempts to reproduce a paper's claimed results. Generates pass/fail verdict per claim.                                                                 |
| **Deep Learner** | Feeds on your entire source corpus. Extracts expert mental models, maps field debates, generates deep-understanding questions, tutors you through them. |


```bash
research-loop hand activate experimenter
research-loop hand activate librarian --schedule daily
research-loop hand status experimenter
research-loop hand list
```

Build your own: define a `HAND.toml` and publish to the registry.

  


## Collaboration

```bash
# Export your session as a portable bundle
research-loop export --session attention-2024 --output ./attention.research

# A collaborator loads it and continues from where you left off
research-loop resume ./attention.research

# They export again — the bundle contains both sessions
research-loop export --output ./attention-v2.research
```

A `.research` bundle contains: `hypothesis.md`, `knowledge_graph.md`, `lab_notebook.md`, `autoresearch.jsonl`, code diffs for every mutation, and an auto-generated README with the best result summary.

  


## MCP Bridge

Don't want to switch tools? Install the MCP server and get Research Loop capabilities inside Claude Code, OpenCode, Cursor, or any MCP-compatible host.

```bash
research-loop mcp serve
```

Exposes: `research_paper_ingest`, `research_loop_start`, `research_kg_query`, `research_library_search`, `research_export`, and more.

The MCP bridge is ~20% of the full experience. The full TUI workspace — Dashboard, Library pane, Writer, Hand lifecycle management — requires the standalone binary.

  


## Workspace layout

```
~/research/
  .research-loop/
    config.toml                      # LLM backend, metric config
    sessions/
      attention-mechanism-2024/
        hypothesis.md                # extracted claim + experiment design
        knowledge_graph.md           # living DAG of every idea tried
        lab_notebook.md              # human-readable experiment log
        autoresearch.jsonl           # machine-readable run history (checkpoint)
        checkpoints/                 # git patches per mutation
    library/
      papers/                        # ingested PDFs + extracted metadata
      index.json                     # full-text search index
      feeds.toml                     # ArXiv RSS subscriptions
    drafts/
      paper-v1.md
    bundles/
```

All state in plain Markdown and JSONL. No database. Fully git-diffable.

  


## Development

```bash
go build ./cmd/research-loop    # Build binary
go test ./...                   # Run tests
go vet ./...                    # Lint
```

See [doc/DEVELOPING.md](doc/DEVELOPING.md) for the full development guide.

  


## Roadmap

**v0.1 — Foundation**

- ⚪ Go binary + Bubble Tea TUI (Chat + Dashboard)
- ⚪ Epistemic + Empirical agents
- ⚪ Paper ingestion pipeline (ArXiv → `hypothesis.md`)
- ⚪ Experiment loop state machine + JSONL checkpointing
- ⚪ Knowledge graph (Markdown DAG)
- ⚪ `.research` bundle export + resume
- ⚪ `research-loop start <url>` magic moment

**v0.2 — Full Workspace**

- ⚪ Library pane (F3) with full-text search
- ⚪ Writer pane (F4) with Vim keybindings
- ⚪ Scribe + Reviewer agents
- ⚪ ArXiv RSS feed monitoring
- ⚪ Cost tracking dashboard
- ⚪ Hand lifecycle management

**v0.3 — Collaboration**

- ⚪ MCP bridge server
- ⚪ Bundle registry
- ⚪ Multi-paper sessions
- ⚪ Replicator Hand
- ⚪ LaTeX/PDF export

**v1.0 — Production**

- ⚪ Deep Learner Hand
- ⚪ Community Hand registry
- ⚪ Homebrew / pip distribution
- ⚪ Documentation site

  


## Contributing

We welcome contributions. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

  


## Community

- [Discord](https://discord.gg/research-loop) — Join the community
- [GitHub Issues](https://github.com/research-loop/research-loop/issues) — Bugs and feature requests
- [GitHub Discussions](https://github.com/research-loop/research-loop/discussions) — Ideas and RFCs

  


## License

MIT &copy; 2026 Research Loop

## Star History

[Star History Chart](https://www.star-history.com/?repos=research-loop%2Fresearch-loop&type=date&legend=top-left)

  


---



Open source under MIT. Built for researchers who want to run experiments, not babysit terminals.