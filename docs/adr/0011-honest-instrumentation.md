---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0011. Honest instrumentation: unmetered currencies declare themselves

## Context and Problem Statement

The whitepaper's central framework holds that every AI request spends five
currencies: latency, cost, trust, freshness, and agency. The gateway meters two of
them today. The other three are not instrumented. What should the Request Receipt
report for a currency that is not yet measured?

The obvious answers are to omit the field, or to emit a zero. Both are lies of a
kind, and this project's stated identity is that it does not tell them.

## Decision Drivers

- A reference that reports fake precision is worse than one that admits a gap.
- Downstream consumers must be able to distinguish "measured as zero" from "not measured".
- The gaps are the roadmap; hiding them hides the roadmap.
- The project's public claim is that it "tracks its own gaps out loud". Code should
  make that claim true, not merely documentation.

## Considered Options

- Omit unmetered currency fields from the Receipt entirely.
- Emit `0` for unmetered currencies.
- Emit the sentinel string `not_metered` for unmetered currencies.

## Decision Outcome

Chosen option: **emit `not_metered`.**

```json
"currencies": {
  "latency_ms": 12,
  "cost_tokens": 21,
  "trust": "not_metered",
  "freshness": "not_metered",
  "agency": "not_metered"
}
```

The Five Currencies are a type in the codebase, not a metaphor in a document.
Three of its fields say, on every single request, that this system does not yet
know what it is spending.

### Consequences

- Good, because no consumer can mistake an unmeasured currency for a cheap one.
  A dashboard averaging `trust: 0` would report perfect safety; a dashboard
  encountering `not_metered` cannot.
- Good, because the roadmap is legible from a single response body. Closing a
  currency is a well-defined contribution.
- Good, because it makes the project's honesty discipline mechanical rather than
  aspirational.
- Bad, because the currency fields are heterogeneously typed - two numbers and
  three strings - which is awkward for a strict schema and for naive consumers.
  Accepted: awkwardness that forces a reader to notice a gap is doing its job.
- Bad, because a sentinel string is a weaker contract than a typed optional. A
  future revision may model it as a nullable measurement with an explicit status.

## Pros and Cons of the Options

### Omit the fields

- Good, because clean payload, correctly typed.
- Bad, because absence is ambiguous: is it unmeasured, or was the field dropped?
  The Five Currencies stop being visible in the artifact that is supposed to
  demonstrate them.

### Emit zero

- Good, because uniformly typed and trivially chartable.
- Bad, because it is false. Zero trust cost, zero staleness, and zero agency are
  each a specific, meaningful, and wrong claim.

### Emit `not_metered` (chosen)

- Good, because unambiguous, visible, and honest.
- Bad, because mixed types; a stronger schema exists and should eventually replace it.

## Revisit When

A currency becomes metered - at which point its field carries a real measurement -
or when the Receipt gains a formal schema, in which case model each currency as a
measurement with an explicit `measured | not_measured` status rather than a
sentinel string.