# Architecture

> **Status: model frozen.** The concepts below are the canonical EAPRA
> architecture, frozen from **Whitepaper 001 - Beyond the AI Demo**. The
> whitepaper remains the single source of truth for the *reasoning*; this page
> is the *executable mapping* of that model onto the repository.
> Nothing here contradicts the paper.

EAPRA's architecture rests on three ideas. Together they answer: what does a
request traverse, what wraps it, and what does it spend?

![EAPRA architecture](../assets/architecture.svg)

*Full diagrams, including the request lifecycle and the Receipt, live in
[`diagrams.md`](./diagrams.md).*

## 1. The Three Planes

Every production AI system resolves into three planes. Confusing them is the
root cause of most "it worked in staging" incidents.

| Plane | What it is | Repository home |
| --- | --- | --- |
| **Request path** | What one user request traverses in real time: gateway → auth → AI gateway → model → retrieval → response. | `request-plane/` |
| **Control plane** | What governs the system *off* the hot path: prompt/model registries, deployment, policy, quotas, routing config, secrets. | `control-plane/` |
| **Data plane** | What grounds answers: knowledge base → ingestion/embedding pipeline → vector index. | `data-plane/` |

The seam between planes is a **Blast-Radius Boundary**. You can re-index the
data plane without touching the request path, and roll a prompt through the
control plane without redeploying the model server.

> **Design principle:** maximize the work you can change on one plane without
> moving the others.

## 2. The Two Envelopes

Two concerns wrap *every* stage of the request path. They are not steps in it.

- **Platform foundation** - Kubernetes, CI/CD, secrets, networking, scaling, caching.
- **Governance & telemetry** - security, compliance, observability, monitoring, logging, cost tracking.

> **Key insight.** Draw observability and security as pipeline stages and you
> will build them as afterthoughts. Draw them as envelopes and you will build
> them as infrastructure.

## 3. The Five Currencies of Production AI

Every AI request spends five currencies simultaneously, from a fixed budget.
Every architecture decision trades one for another. There is no free lunch.

| Currency | Unit | Metered today? |
| --- | --- | --- |
| **Latency** | milliseconds | ✅ yes |
| **Cost** | tokens | ✅ yes |
| **Trust** | safety & correctness | ❌ `not_metered` |
| **Freshness** | recency of retrieved knowledge | ❌ `not_metered` |
| **Agency** | tool access & autonomy | ❌ `not_metered` |

A guardrail buys *trust* with *latency*. A semantic cache buys *latency* and
*cost* but risks *trust*. A larger retrieval set buys *freshness* at the cost of
*cost* and *latency*.

**The AI gateway sets the allocation. Observability is the meter.**

### This is not a metaphor - it is a struct

The Five Currencies are literally a type in the codebase
(`internal/receipt.Currencies`). Two currencies are metered today. The other
three report `not_metered` rather than a comfortable zero, because a reference
that reports fake precision is worse than one that admits a gap.

```json
"currencies": {
  "latency_ms": 12,
  "cost_tokens": 21,
  "trust": "not_metered",
  "freshness": "not_metered",
  "agency": "not_metered"
}
```

Each unmetered currency is an open issue. Closing them is the roadmap.

## The Request Receipt

One trace ties an answer to the prompt, model, tokens, latency, and cost that
produced it. Attribute names follow the OpenTelemetry GenAI semantic
conventions (`gen_ai.*`) so the telemetry is portable rather than bespoke.

> Incidents without receipts are archaeology. Incidents with receipts are replay.

## Where the code lives

```
request-plane/ai-gateway/     the decide-and-meter control layer (original code)
  internal/provider/          the Provider SPI + adapters (stub, openai-compatible)
  internal/meter/             token metering and per-caller budgets
  internal/receipt/           the Request Receipt and its emitters
  internal/server/            the request path, wired
```

`control-plane/` and `data-plane/` are deliberately empty. They are named so the
model is visible from the directory tree, and filled as the roadmap reaches them.