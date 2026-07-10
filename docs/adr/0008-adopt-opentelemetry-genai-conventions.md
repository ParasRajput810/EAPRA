---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0008. Adopt the OpenTelemetry GenAI semantic conventions for the Request Receipt

## Context and Problem Statement

Every request through the gateway produces telemetry: which model, how many
tokens, how long, at what cost. That telemetry can be shaped however we like. If
we invent our own field names, every consumer of EAPRA's output must learn a
vocabulary that exists nowhere else, and the project's observability advice
becomes unfollowable outside the project. What schema should the Request Receipt use?

## Decision Drivers

- Portability: telemetry should outlive any single vendor or backend.
- Teaching value: the reference should model the practice it recommends.
- The whitepaper argues that observability is the meter on the Five Currencies;
  a bespoke meter reads in units nobody else recognises.
- Adoption cost for a contributor already familiar with OpenTelemetry.

## Considered Options

- Invent EAPRA-specific field names (`model`, `tokens_in`, `tokens_out`).
- Adopt a vendor's proprietary AI-monitoring schema.
- Adopt the OpenTelemetry GenAI semantic conventions (`gen_ai.*`).

## Decision Outcome

Chosen option: **the OpenTelemetry GenAI semantic conventions.** The Receipt emits
`gen_ai.provider.name`, `gen_ai.request.model`, `gen_ai.response.model`,
`gen_ai.usage.input_tokens`, `gen_ai.usage.output_tokens`, and
`gen_ai.response.finish_reason`.

This is a schema decision, independent of whether the OpenTelemetry SDK is a
dependency (it is not — see ADR-0007). Conventions and SDK are separable, and we
take the conventions first.

### Consequences

- Good, because a receipt is intelligible to any OTel-aware tool without translation.
- Good, because wiring the OTel exporter later becomes mechanical rather than a redesign.
- Good, because the repository practises the standard it advocates.
- Bad, because the GenAI conventions are still stabilising; attribute names may
  change and we will have to follow them. Pin to the version tested, and treat a
  convention change as a normal upstream upgrade.
- Bad, because `gen_ai.*` field names appear in a codebase whose `go.mod` contains
  no OpenTelemetry dependency, which is surprising until ADR-0007 is read.

## Pros and Cons of the Options

### EAPRA-specific field names

- Good, because total freedom and no upstream churn.
- Bad, because it makes EAPRA's telemetry a dialect. Precisely the
  "observability as a dashboard, not a schema" anti-pattern the whitepaper names.

### A vendor's proprietary schema

- Good, because immediate rich tooling.
- Bad, because it embeds lock-in into the reference architecture itself.

### OpenTelemetry GenAI conventions (chosen)

- Good, because portable, community-governed, and already the ecosystem's direction.
- Bad, because still evolving; we accept following it.

## Revisit When

The GenAI conventions reach a stable release, at which point pin the stable
version; or if a convention change is incompatible enough to warrant a
superseding ADR.