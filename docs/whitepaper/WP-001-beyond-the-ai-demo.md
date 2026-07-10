# Whitepaper 001 - Beyond the AI Demo

**A Platform Engineering Guide to Building Secure, Scalable, and Observable Enterprise AI Systems**

Series: *Engineering Enterprise AI Infrastructure*
Version: 0.1 (draft) · Status: **living document** · Licence: Apache-2.0

---

## Status and honesty statement

Read this first, because it governs everything below.

This whitepaper is the **single source of truth for EAPRA's architecture**. Where
the repository and this paper disagree, this paper is correct and the repository
is a bug.

It is also a *draft*, and it is honest about what it is not:

- **No first-party measurements.** Every number required to make a claim about
  latency, cost, or throughput is absent, because none have been produced yet.
  Producing them is the purpose of the benchmark harness on the roadmap. Nothing
  in this paper reports a measurement that was not measured.
- **No production war stories.** The author has not operated the system described
  here at enterprise scale. Scenarios are labelled **Illustrative** and are
  constructed to show how the architecture behaves, not to report what happened.
- **External claims are cited.** Where this paper asserts something about the
  wider ecosystem, it cites an official source (CNCF, Kubernetes SIG-Network,
  OpenTelemetry, OWASP). Where it introduces an original framework, it says so.

A reference that reports fake precision is worse than one that admits a gap.

---

## Abstract

Most organisations can build an AI demo in an afternoon. Few can operate one in
production. The distance between those two states - measured in authentication,
cost control, routing, retrieval quality, observability, security, and change
management - is not a machine-learning problem. It is a **platform-engineering**
problem.

This paper reframes enterprise AI infrastructure around a single economic idea:
**every AI request continuously spends five currencies - latency, cost, trust,
freshness, and agency - and every architectural decision trades one for another.**
From that lens it derives a reference architecture (the *Three Planes*, wrapped
in *Two Envelopes*), a per-request accounting artifact (the *Request Receipt*), a
catalogue of anti-patterns, and a maturity model.

Its companion repository, **EAPRA**, implements that architecture as runnable
code, so the paper can be executed rather than only read.

**Thesis:** *The demo proves the idea is possible. The infrastructure proves it
is a product.*

---

## Scope and audience

**In scope:** the infrastructure *around* inference - gateways, routing,
retrieval, serving, observability, security, cost, and lifecycle.

**Out of scope:** model training and fine-tuning; model selection; prompt craft.

**Audience:** platform, infrastructure, and DevOps engineers; SREs; staff and
principal architects. Working familiarity with Kubernetes, CI/CD, and cloud
infrastructure is assumed. No ML research background is required.

---

## 1. Introduction: the afternoon and the nine months

A demo has one user, one happy path, and a forgiving audience. Production has a
*distribution*: thousands of concurrent requests, an adversary in the input
stream, a compliance reviewer, a finance owner, and a pager.

The model rarely fails. The system around the model fails - and it fails in the
places a demo never exercises.

> **The Distribution Principle.** The demo has one happy path; production has a
> distribution. Everything hard about enterprise AI lives in the tails.

| Dimension | AI demo | Enterprise AI production |
| --- | --- | --- |
| Users | One, cooperative | Many, concurrent, adversarial |
| Identity | None | AuthN/AuthZ *before* any token is spent |
| Failure | Retry by hand | Backoff, circuit breakers, fallback models |
| Cost | Ignored | An SLO, attributed per tenant |
| Data | Pasted documents | Governed pipeline: ingest → embed → index → refresh |
| Observability | `print()` | Trace, quality, and cost per request |
| Change | Edit and rerun | Shipped and rolled back like code |
| Security surface | The prompt box | The full OWASP LLM Top 10 |

---

## 2. The landscape: why cloud-native must adapt

The CNCF's *Cloud Native Artificial Intelligence* whitepaper states the mismatch
precisely: cloud-native was designed for **stateless microservices that scale
horizontally and fail fast**, whereas AI workloads are **stateful and
GPU-hungry** - weights that must persist, indexes that must stay warm,
accelerators that dwarf ordinary compute. [CNCF]

That is the field's central tension in one sentence:

> We run stateful, expensive, non-deterministic workloads on a substrate built
> for stateless, cheap, deterministic ones.

Everything downstream - routing, scaling, caching, observability - is an
adaptation to that mismatch.

Two forces make this *now* rather than later. Inference has become a
revenue-bearing production workload, dragging it into the same reliability and
cost discipline as any other tier. And the ecosystem has begun to standardise the
adaptations: the Kubernetes inference gateway [K8s], the OpenTelemetry GenAI
semantic conventions [OTel], and the OWASP LLM risk taxonomy [OWASP] each
appeared to close a specific gap this paper describes. When standards bodies
move, the patterns are no longer speculative.

---

## 3. The Three Planes

*(Original framework introduced by this series.)*

Every production AI system resolves into three planes. Confusing them is the root
cause of most "it worked in staging" incidents.

| Plane | What it is |
| --- | --- |
| **Request path** | What one user request traverses in real time: gateway → auth → AI gateway → model → retrieval → business services → response. |
| **Control plane** | What governs the system *off* the hot path: model and prompt registries, deployment and promotion, policy, quotas, routing configuration, secrets. |
| **Data plane** | What grounds answers: knowledge base → ingestion and embedding pipeline → vector index. |

The seam between planes is a **Blast-Radius Boundary**: an isolation seam that
bounds the reach of a failure *or of a change*. You can re-index the data plane
without touching the request path. You can roll a new prompt through the control
plane without redeploying the model server.

> **Design principle.** Maximise the work you can change on one plane without
> moving the others.

---

## 4. The Two Envelopes

*(Original framework.)*

Two concerns wrap **every stage** of the request path. They are not stages in it.

- **The platform-foundation envelope** - Kubernetes, CI/CD, secrets, networking,
  scaling, caching. The substrate every stage runs on.
- **The governance-and-telemetry envelope** - security, compliance,
  observability, monitoring, logging, cost tracking. The oversight every stage is
  subject to.

> **Key insight.** Draw observability and security as pipeline stages and you will
> build them as afterthoughts. Draw them as envelopes and you will build them as
> infrastructure.

This is the most common architectural error in the domain, and it is committed at
the whiteboard, not in the code.

---

## 5. The Five Currencies of Production AI

*(Original framework. The organising idea of this paper.)*

Every AI request spends five currencies simultaneously, from a fixed budget.
Every design choice trades one for another. There is no free lunch.

| Currency | Unit | What it buys |
| --- | --- | --- |
| **Latency** | milliseconds | Responsiveness |
| **Cost** | tokens | The computation itself |
| **Trust** | safety and correctness | Confidence in the answer |
| **Freshness** | recency of retrieved knowledge | Relevance |
| **Agency** | tool access and autonomy | Capability to act |

The trades are constant and mostly invisible unless you instrument them:

- A **guardrail** buys *trust* with *latency*.
- A **semantic cache** buys *latency* and *cost*, and risks *trust* (a wrong
  similar-hit).
- A **larger retrieval set** buys *freshness* at the cost of *cost* and *latency*.
- An **autonomous tool call** spends *agency* to buy capability, and enlarges the
  attack surface.
- **Routing to a smaller model** buys *cost* and *latency*, and spends *trust*.

Two consequences follow, and they are the practical heart of this paper:

1. **The AI gateway is where the allocation is set.** It is the one place in the
   request path with the authority to decide how much of each currency a request
   may spend.
2. **Observability is the meter.** A currency you do not measure is a currency you
   spend blindly.

### 5.1 Cost is denominated in tokens, not requests

This deserves its own statement because it is the single most common production
error.

> **Request-count rate limiting is a lie your API gateway tells you about cost.**

One prompt can cost orders of magnitude more than another. A limiter that counts
calls will pass while the invoice triples. OWASP names the resulting abuse
**unbounded consumption** (LLM10) - colloquially, *Denial of Wallet*. [OWASP]

You budget in tokens, or you do not budget at all.

---

## 6. The Request Receipt

*(Original framework.)*

The operational question in production is never "is it up?" It is:

> *Which* prompt, retrieving *which* chunks, from *which* index version, through
> *which* model, at *what* cost - and was the answer any good?

The **Request Receipt** is the per-request record that answers it: one trace tying
an answer to its prompt, retrieval set, model version, token counts, latency,
cost, and quality score.

Attribute names follow the **OpenTelemetry GenAI semantic conventions**
(`gen_ai.*`) so telemetry is portable across vendors rather than bespoke. [OTel]

> Incidents without receipts are archaeology. Incidents with receipts are replay.

**Quality is a telemetry signal, not a batch job.** Continuous evaluation should
emit scores into the same traces, so a quality regression is an alert rather than
a customer complaint.

---

## 7. The AI gateway: five jobs, five trade-offs

The AI gateway is the component a demo lacks, and the reason a system is an *AI*
platform rather than a generic API platform. It sits after authentication and
performs five jobs - each a deliberate spend against the Five Currencies.

| Job | Buys | Trade-off |
| --- | --- | --- |
| Token-aware rate limiting | Bounds *cost*; protects against abuse | Needs per-model tokenizer awareness; small pre-flight cost |
| Model routing | Trades *cost* for *trust* per intent | A silent mis-route degrades quality invisibly; routing must itself be observable |
| Semantic + prompt caching | *Latency* and *cost* | Risks *trust*: a subtly wrong similar-hit. Thresholds must be tuned and audited |
| Guardrails | *Trust* | Spends *latency*; false positives spend user experience |
| Resilience (backoff, breakers, failover) | Reliability | Aggressive retries against a degraded provider amplify load |

### When *not* to build a heavy AI gateway

A single internal tool with a handful of trusted users and no cost sensitivity
does not need model routing or semantic caching on day one. Build authentication
and a token budget first; add the rest when traffic or spend justifies it.
**Premature gateway complexity is its own failure mode.**

---

## 8. Engineering challenges

Almost nothing here is a *new* distributed-systems problem. It is the old problems
re-weighted: latency is longer, cost is variable, inputs are adversarial by
construction, and correctness is statistical.

**Non-determinism.** The same input can produce different outputs. This breaks
exact-match testing and forces a shift from assertions to *evaluations* - scored,
statistical, continuous.

**Tail latency.** An LLM call is seconds, not milliseconds, with high variance.
p95/p99 dominate user experience, and every synchronous hop (auth, guardrail,
retrieval, model) adds to the tail. Each guardrail that buys *trust* is paid for
from the *latency* budget.

**Provider instability.** Managed endpoints return 429s and 503s under load. The
classic toolkit applies directly:

- **Retries with jittered exponential backoff**
- **Circuit breakers** to stop hammering a degraded dependency
- **Timeouts and bulkheads** to contain blast radius
- **Idempotency** so a retried request does not double-act

> **Warning.** Naive retries against a struggling provider are a self-inflicted
> outage. Backoff and circuit-breaking are not optional.

**Prompts and indexes are artifacts.** A prompt is code; a retrieval index is a
build output. Both must be versioned, reviewed, gated by evaluation, and rolled
back independently. Editing a prompt directly in production is the AI-era
equivalent of SSH-ing into prod to change a config - and it fails the same way.

---

## 9. Security

Start from the property that makes this domain different:

> **An LLM processes instructions and data in the same channel, so it cannot
> reliably tell content from command.**

That single fact generates most of the risk surface. It is why **prompt injection
is not a bug you patch in the model - it is a system property you contain in the
architecture.**

Map controls to a recognised taxonomy rather than inventing one. The **OWASP Top
10 for LLM Applications (2025)** provides the shared vocabulary. [OWASP]

| ID | Risk | Where it is contained |
| --- | --- | --- |
| LLM01 | Prompt injection (direct and indirect) | Segregate retrieved content as untrusted data; input guardrails at the gateway; least-privilege tools |
| LLM02 | Sensitive information disclosure | Output filtering; PII redaction; retention policy on logs and caches |
| LLM03 | Supply chain | Vet, sign, and pin models, adapters, and datasets |
| LLM04 | Data and model poisoning | Provenance on knowledge-base sources; index write-access control |
| LLM05 | Improper output handling | Never pass model output to a downstream system unvalidated |
| LLM06 | Excessive agency | Least-privilege tools; human-in-the-loop for irreversible actions |
| LLM07 | System prompt leakage | Never place secrets in system prompts; assume they leak |
| LLM08 | Vector and embedding weaknesses | Tenant isolation in the vector store; access control on retrieval |
| LLM09 | Misinformation | Groundedness checks |
| LLM10 | Unbounded consumption | Token-aware rate limits and quotas; timeouts; abuse monitoring |

Two deserve emphasis because teams underestimate them.

**Indirect prompt injection (LLM01)** is the dangerous variant. The attacker never
speaks to your system. They plant instructions in a document, web page, or ticket
that your retrieval pipeline later fetches and the model obeys. The defence is
architectural: treat *all* retrieved content as untrusted, denote it as data, and
constrain what the model may *do* with tools regardless of what it is told. This
paper names that pattern **Zero-Trust Retrieval**.

**Excessive agency (LLM06)** is the agentic-era risk. Give an agent more tools,
broader permissions, or more autonomy than its task requires and you have built
the attack surface yourself. Contain it with least-privilege tool scoping and
human approval gates on high-impact actions. This paper names that pattern the
**Safe Agent Envelope**.

**Authentication is the first cost control.** An unauthenticated inference
endpoint is not merely an access-control failure; it is a metered resource anyone
can drain. Establish identity before a single token is spent.

**When *not* to add another guardrail.** Guardrails cost latency and produce false
positives. Layer them by the *risk of the action*, not uniformly. A read-only
summariser and an agent that can issue refunds do not warrant the same gate.

---

## 10. Scaling

**Inference load is not request-count load.** Two requests to the same endpoint can
differ by orders of magnitude in tokens, latency, and GPU memory. Scale on request
count and you will over-provision on cheap traffic and fall over on expensive
traffic.

**Scale the inference tier independently from the API tier.** The API tier is
stateless and cheap - scale it on CPU and concurrency. The inference tier is
GPU-bound and expensive - scale it on the signals that predict saturation: queue
depth, in-flight concurrency, and KV-cache pressure.

**Round-robin is the wrong router for GPUs.** Traditional L7 load balancing sends a
request to a busy accelerator while an idle one waits, because it cannot see
model-server internals. The **Kubernetes Gateway API Inference Extension** exists
to close exactly this gap: it turns an ext-proc-capable gateway into an
*inference gateway* whose **Endpoint Picker** routes on live model-server metrics
- queue length, loaded adapters, KV-cache state - with request-cost-aware
scheduling, per-request criticality, and safe model rollouts via the
`InferencePool` and `InferenceObjective` CRDs. [K8s]

### Managed vs self-hosted

This is the decision with the largest cost and operational consequence. Ratings
below are qualitative and **illustrative of typical conditions**, not benchmarks.

| Dimension | Managed API | Self-hosted |
| --- | --- | --- |
| Time to first deploy | Fast | Slow (GPU capacity, serving stack) |
| Marginal cost at low volume | Moderate | High (idle GPUs waste money) |
| Marginal cost at high steady volume | High (per-token dominates) | Low (amortised GPU) |
| Tail-latency control | Limited | Tunable (batching, KV cache, routing) |
| Data control and residency | Provider-constrained | Full |
| Operational burden | Low | High |
| Vendor lock-in | High | Low |

**When *not* to self-host:** if you cannot keep GPUs utilised, self-hosting is more
expensive *and* more work. Utilisation, not principle, is the deciding variable.

### Resilience for a stateful data plane

Because the data plane is stateful, classic stateless HA assumptions break. Plan
for vector-index replication and rebuild procedures, cache warming after failover
(a cold semantic cache after a region flip is a latency cliff), and provider or
region failover at the gateway.

---

## 11. Observability

Standardise on the **OpenTelemetry GenAI semantic conventions**. [OTel] Emit the
`gen_ai.*` attributes - provider, request and response model, token usage, finish
reason, retrieval documents, evaluation scores - and, for agentic systems, the
agent and tool-call span conventions, so one `trace_id` links a decision through
each tool execution to the final response.

Content capture (the actual prompts and completions) is supported but defaults
off, because it can contain PII. Enable it deliberately, with redaction. [OTel]

| | Traditional monitoring | AI observability |
| --- | --- | --- |
| Unit | Request/response | Prompt → retrieval → model → tool chain |
| Golden signals | Latency, errors, saturation | *plus* tokens, cost, groundedness |
| "Correct" means | HTTP 200 | 200 *and* grounded *and* on-budget |
| Standard | OTel HTTP semconv | OTel **GenAI** semconv (`gen_ai.*`) |

> **Observability is not a dashboard. It is the meter on five currencies.**

---

## 12. Anti-patterns

Named anti-patterns transfer faster than principles.

1. **The Hardcoded Provider.** Services call one model SDK directly. *Fix:* route
   all model traffic through one gateway; the model becomes a swappable dependency.
2. **Counting Requests, Not Tokens.** Limits pass while the bill explodes. *Fix:*
   token-aware budgets (LLM10).
3. **Prompt-in-Prod.** Behaviour changes with no diff and no rollback. *Fix:*
   prompts as gated, versioned artifacts.
4. **The Untraced Request.** Every incident becomes archaeology. *Fix:* the Request
   Receipt.
5. **Trusting Your Own Data.** Retrieved text treated as safe. *Fix:* Zero-Trust
   Retrieval (LLM01).
6. **Round-Robin GPUs.** Tail-latency cliffs under load. *Fix:* an inference
   gateway that routes on KV-cache and queue state.
7. **CPU-Percentage Autoscaling for Inference.** Wrong signal entirely. *Fix:* scale
   on queue depth, concurrency, and cache pressure.
8. **The Uncapped Agent.** Broad tools, no budget, no human gate. *Fix:* the Safe
   Agent Envelope (LLM06).
9. **Semantic Cache Without a Threshold Policy.** Fast, confident, wrong answers.
   *Fix:* tuned and audited similarity thresholds; log and evaluate cache hits.
10. **Secrets in the System Prompt.** *Fix:* a vault. Assume prompts leak (LLM07).
11. **Evaluation as a One-Time Gate.** Quality drifts silently after launch. *Fix:*
    continuous evaluation as telemetry.
12. **Observability as a Dashboard, Not a Schema.** Metrics exist but do not
    correlate. *Fix:* OpenTelemetry GenAI conventions.
13. **Every Team Builds Its Own.** Six inconsistent, insecure gateways. *Fix:* a
    paved road owned by a platform team.

---

## 13. Illustrative scenarios

**These are constructed illustrations, not accounts of real deployments.** They
exist to show that the architecture does not change across industries - only the
*currency allocation* does.

**Illustrative - regulated healthcare intake assistant.** Bursty, low daily volume;
strict residency; high trust requirement. *Allocation:* spend heavily on trust
(guardrails, human-in-the-loop) and on control (a managed model inside a private
boundary); accept higher latency; keep agency near zero.

**Illustrative - high-volume SaaS support assistant.** Steady, latency-sensitive,
margin-sensitive. *Allocation:* self-host behind an inference gateway with
cache-aware routing; aggressive semantic caching with audited thresholds;
progressive model routing. Utilisation is high enough to amortise GPUs and every
millisecond and token is user-visible.

**Illustrative - financial research assistant.** Moderate volume; freshness- and
trust-critical; some agency (querying data tools). *Allocation:* spend on
freshness (frequent re-index) and trust (groundedness checks); constrain agency
with read-only tools and a Safe Agent Envelope.

> **The architecture does not change across industries. The currency allocation does.**

---

## 14. The Production AI Maturity Model

*(Original framework.)* Locate yourself honestly before choosing what to build next.

| Level | Name | Characteristics |
| --- | --- | --- |
| **L0** | Demo | One path, no auth, no cost control, no telemetry |
| **L1** | Guarded | AI gateway, authentication, token budget, basic tracing |
| **L2** | Governed | OWASP-mapped security, Zero-Trust Retrieval, prompts and indexes as gated artifacts, standards-based observability |
| **L3** | Scaled | Independent inference scaling, cache-aware routing, audited caching, cost as an SLO |
| **L4** | Self-optimising | Continuous evaluation as telemetry, adaptive routing and retrieval, safe agents, multi-region resilient data plane |

---

## 15. What EAPRA implements today

This section exists so the paper and the repository can never quietly drift apart.
It is deliberately unflattering.

| Concept in this paper | Status in the repository |
| --- | --- |
| Three Planes | Structure only - `request-plane/` is real; `control-plane/` and `data-plane/` are named, empty, and documented as such |
| Two Envelopes | Partial - telemetry envelope present as the Receipt; platform envelope is CI only |
| Five Currencies | **Latency and cost metered.** Trust, freshness, and agency report `not_metered` |
| Request Receipt | Implemented; `gen_ai.*` attribute names emitted per request |
| Token-aware budgeting | Implemented; reserved before the provider is called |
| AI gateway - routing | Not implemented |
| AI gateway - caching | Not implemented |
| AI gateway - guardrails | Not implemented |
| AI gateway - resilience | Not implemented (no backoff, no circuit breaker) |
| Provider SPI | Implemented; two adapters (stub, OpenAI-compatible) |
| OpenTelemetry exporter | **Not wired.** Conventions are emitted; the SDK exporter sits behind a port |
| Authentication | **Not implemented.** Caller identity is an unvalidated header |
| Retrieval / RAG data plane | Not implemented |
| Inference gateway | Not implemented |
| Benchmarks | **None published.** This is the paper's largest gap |

Current maturity: **L0 → L1 in progress.**

Each row marked *not implemented* is an open issue, not an aspiration hidden in
prose.

---

## 16. Open problems

- **Quality has no SLO vocabulary.** There is no widely accepted way to express a
  service-level objective for *answer quality* the way there is for latency and
  errors. Whoever standardises it will shape the decade.
- **Delegated agency has no settled security model.** Agent-to-agent protocols and
  tool-integration standards are maturing faster than the authorisation models
  that should constrain them.
- **Evaluation as infrastructure.** Continuous, online evaluation is still bespoke
  per team. It should be a platform service.
- **The economics of self-optimising systems.** Routers and caches that tune their
  own currency allocation require observability we do not yet trust.

---

## References

Official sources only. Version-specific details evolve; verify against the linked
documents before relying on them.

- **[CNCF]** CNCF TAG Runtime - *Cloud Native Artificial Intelligence Whitepaper.*
  <https://www.cncf.io/reports/cloud-native-artificial-intelligence-whitepaper/>
- **[K8s]** Kubernetes SIG-Network - *Gateway API Inference Extension.*
  <https://gateway-api-inference-extension.sigs.k8s.io/> ·
  <https://kubernetes.io/blog/2025/06/05/introducing-gateway-api-inference-extension/>
- **[OTel]** OpenTelemetry - *Generative AI Semantic Conventions.*
  <https://opentelemetry.io/docs/specs/semconv/gen-ai/>
- **[OWASP]** OWASP GenAI Security Project - *Top 10 for LLM Applications (2025).*
  <https://genai.owasp.org/llm-top-10/>

Supporting sources to cite where used: Kubernetes HPA; KEDA; vLLM; KServe;
NVIDIA Triton; HashiCorp Vault; Redis; NIST AI Risk Management Framework.

---

## Glossary

Terms marked ★ are original coinings introduced by this series. They are offered
as shared language, not as established industry standards.

- **Request path / control plane / data plane** - the three planes every AI system resolves into. ★
- **Blast-Radius Boundary** ★ - an isolation seam that bounds the reach of a failure or a change.
- **Two Envelopes** ★ - the platform-foundation and governance-and-telemetry concerns that wrap every stage of the request path.
- **Five Currencies of Production AI** ★ - latency, cost, trust, freshness, agency.
- **Request Receipt** ★ - the per-request record tying an answer to prompt, model, tokens, latency, cost, and quality.
- **AI gateway** - the request-path control point for token-aware limiting, routing, caching, guardrails, and failover.
- **Inference gateway** - a Kubernetes gateway extended to route on live model-server metrics.
- **Zero-Trust Retrieval** ★ - treating all retrieved content as untrusted input.
- **Safe Agent Envelope** ★ - least-privilege tools plus a human gate on irreversible actions.
- **Denial of Wallet** - cost-exhaustion abuse of a metered endpoint (OWASP LLM10).
- **Paved road** - the platform team's supported golden path.

---

## Contributing to this paper

This document is versioned in the repository and changes through pull requests,
like code. Disagreement is the point. If you believe a claim here is wrong,
unsupported, or overstated, open an issue - especially if it is a claim the author
would prefer to keep.

*Copyright 2026 The EAPRA Authors. Licensed under the Apache License, Version 2.0.*