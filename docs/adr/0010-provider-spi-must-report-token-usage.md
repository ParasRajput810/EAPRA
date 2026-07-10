---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0010. The Provider SPI is minimal, and every provider must report token usage

## Context and Problem Statement

The `Provider` interface is one of the few things EAPRA builds itself
(ADR-0003). Every method added to it is a method every future adapter must
implement, forever. SPIs are where "thin" is hardest to hold -- a risk flagged in
ADR-0003 itself. What shape should this contract take?

## Decision Drivers

- Every method on the SPI is a tax on every future adapter.
- Cost is an SLO; a component that cannot report its cost cannot be governed.
- Adapters must be swappable without the core knowing which one is loaded.
- A newcomer should be able to write an adapter by reading one file.

## Considered Options

- A rich SPI: streaming, tool calling, embeddings, caching hints, retries.
- A minimal SPI where usage reporting is optional.
- A minimal SPI where usage reporting is **mandatory**.

## Decision Outcome

Chosen option: **a minimal SPI -- `Name()` and `Complete()` -- in which the response
type requires input and output token counts.**

A provider that cannot report token usage is unusable in a system where cost is a
first-class currency, so usage is not an optional field on the response. It is
part of the contract.

Capabilities the SPI does *not* have today -- streaming, tool calling, embeddings --
are deliberately absent. Each will be added only when a concrete need exists, and
adding one warrants its own ADR.

### Consequences

- Good, because a new adapter is one small file (see `internal/provider/stub`).
- Good, because the gateway can meter, budget, and produce a Receipt without
  knowing which provider answered.
- Good, because the constraint surfaces bad providers early rather than silently.
- Bad, because a backend that genuinely cannot report usage cannot be adapted
  without estimating on its behalf -- and an estimate must then be labelled as such,
  never presented as reported truth.
- Bad, because streaming responses do not fit `Complete()` and will require the
  interface to grow. That growth must be deliberate, not incidental.

## Pros and Cons of the Options

### Rich SPI up front

- Good, because it anticipates future needs.
- Bad, because it taxes every adapter with methods most will stub out, and it
  guesses at requirements before they exist.

### Minimal SPI, optional usage

- Good, because it accepts any backend.
- Bad, because cost governance becomes best-effort, and the Five Currencies model
  collapses at its most important currency.

### Minimal SPI, mandatory usage (chosen)

- Good, because smallest contract that still supports the architecture's core claim.
- Bad, because it excludes backends that report nothing, and it must grow for streaming.

## Revisit When

Streaming, tool calling, or embeddings become a concrete requirement. Each is a
separate decision; none should be added pre-emptively.