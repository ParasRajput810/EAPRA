---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0007. Zero-dependency core; OpenTelemetry exporter behind the Receipt port

## Context and Problem Statement

ADR-0004 chose an OpenAI-compatible adapter "traced with OpenTelemetry GenAI" as
the first slice. In implementation, a question surfaced: should the gateway core
depend on the OpenTelemetry SDK from the first commit, or should it emit the
GenAI *semantic conventions* through an interface that an OTel exporter later
implements?

## Decision Drivers

- A reference's core should be readable in one sitting and trivially auditable.
- Supply-chain surface in a teaching artifact is a cost, not a neutral.
- Telemetry must be honest and portable from request one (ADR-0004).
- ADR-0003 says: integrate substrates, do not build them. Telemetry backends are substrate.
- A contributor should be able to run the project with no network and no accounts.

## Considered Options

- Depend on the OpenTelemetry Go SDK in the core from the first commit.
- Emit ad-hoc logs now and adopt conventions later.
- Zero-dependency core emitting `gen_ai.*` conventions through an `Emitter` port,
  with the OTel SDK exporter as a plugin.

## Decision Outcome

Chosen option: **zero-dependency core with a Receipt `Emitter` port.** The
`gen_ai.*` semantic-convention attribute names and the Request Receipt shape land
on day one. The OTel SDK exporter is an implementation of `receipt.Emitter`, added
as its own reviewable change.

This **refines ADR-0004 rather than contradicting it**: the conventions -- the part
that makes telemetry portable and honest -- are present from the first request.
Only the exporter wiring is sequenced.

### Consequences

- Good, because the core has zero third-party dependencies: it compiles anywhere,
  reviews quickly, and carries no supply-chain surface.
- Good, because `go test ./...` and `make run` work offline, with no accounts.
- Good, because the Emitter port makes the exporter a substrate integration,
  exactly as ADR-0003 requires.
- Bad, because until the exporter lands, receipts reach stdout and an in-memory
  recorder rather than a tracing backend. This is a real gap, tracked as an open
  issue and stated in the README rather than hidden.

## Pros and Cons of the Options

### OTel SDK in the core from commit one

- Good, because traces reach a backend immediately.
- Bad, because it couples the core to a large dependency tree before the core's
  own shape is settled, and blocks offline/no-account contribution.

### Ad-hoc logs, conventions later

- Good, because it is the least work now.
- Bad, because retrofitting conventions is exactly the "bolt telemetry on later"
  anti-pattern the whitepaper argues against.

### Zero-dep core + Emitter port (chosen)

- Good, because honest conventions now, clean substrate integration next.
- Bad, because "OpenTelemetry" appears in the receipt's field names before it
  appears in `go.mod` -- which must be stated plainly, not implied away.

## Revisit When

The OTel exporter lands (closing the gap above), or a core requirement genuinely
cannot be met without a third-party dependency. Each new dependency in the core
warrants its own ADR.