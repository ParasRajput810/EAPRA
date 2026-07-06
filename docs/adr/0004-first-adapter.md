---
status: "accepted"
date: 2026-07-06
deciders: [The EAPRA Authors]
---

# 0004. First adapter: an OpenAI-compatible chat adapter, traced with OpenTelemetry GenAI

## Context and Problem Statement

The walking skeleton needs exactly one end-to-end slice to prove the
architecture: a request entering the gateway, reaching a model through a plugin,
and returning with full telemetry. Which first integration proves the most
architecture for the least effort?

## Decision Drivers

- Maximum architecture proven per unit of effort (the thinnest real slice).
- Proves the provider SPI (ADR-0004) as a real, working contract.
- Emits a genuine Request Receipt from the very first request (what makes this a
  *reference*, not a demo).
- Reaches multiple backends without multiple adapters.

## Considered Options

- A single proprietary provider SDK first
- A self-hosted server (e.g. vLLM) first
- An OpenAI-compatible chat adapter + OpenTelemetry GenAI tracing, together
- A mock/echo provider first

## Decision Outcome

Chosen option: "An OpenAI-compatible chat adapter, traced with OpenTelemetry
GenAI from request one." The OpenAI-compatible chat API is a de facto standard,
so the *same* adapter also talks to vLLM, Ollama, and other compatible
endpoints - one piece of code unlocks several backends. Pairing it with OTel
GenAI tracing means the skeleton emits a real receipt (model, token usage,
latency) on day one, proving the provider SPI and the observability envelope in
a single slice.

### Consequences

- Good, because one adapter reaches many backends (OpenAI, vLLM, Ollama, …).
- Good, because the SPI and honest telemetry are proven together, immediately.
- Bad, because "OpenAI-compatible" is not perfectly uniform across providers;
  small surface differences will need handling later. (Accepted.)

## Pros and Cons of the Options

### Proprietary SDK first

- Good, because quickest single-provider path.
- Bad, because it proves one backend and risks coupling the core to an SDK.

### Self-hosted (vLLM) first

- Good, because it exercises the serving path early.
- Bad, because GPU/serving ops burden slows the *first* slice unnecessarily.

### OpenAI-compatible adapter + OTel

- Good, because widest reach and honest telemetry from the thinnest slice.
- Bad, because compatibility edges differ across providers.

### Mock/echo provider

- Good, because zero external dependency.
- Bad, because it proves plumbing, not a real provider integration or receipt.

## Revisit When

Revisit if a target backend diverges from the OpenAI-compatible contract enough
to warrant a dedicated adapter, at which point add it as a second plugin behind
the same SPI.