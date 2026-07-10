# Architecture Decision Records (ADRs)

Each file records one significant architectural decision using the
[MADR](https://adr.github.io/madr/) template. See
[`0000-template.md`](./0000-template.md) to start a new ADR.

**Naming:** `NNNN-short-title.md`, zero-padded and incrementing.

**Status lifecycle:** `proposed` → `accepted` → later `deprecated` or
`superseded`. **ADRs are immutable once accepted.** Supersede rather than edit, so
the decision history stays honest.

**Every ADR records the cost of the chosen option, not only its benefit.** An ADR
with no "Bad, because…" line is not finished.

## Index

| ADR | Title | Status |
| --- | --- | --- |
| [0001](./0001-language-posture.md) | Language posture: Go core, polyglot plugins over a network SPI | accepted |
| [0002](./0002-build-vs-integrate.md) | Build vs. integrate: build the control layer, integrate every substrate | accepted |
| [0003](./0003-first-adapter.md) | First adapter: OpenAI-compatible chat, traced with OTel GenAI | accepted |
| [0005](./0005-plane-aligned-repository-structure.md) | Plane-aligned repository structure over strict DDD | accepted |
| [0006](./0006-walking-skeleton-first.md) | Walking skeleton before breadth | accepted |
| [0007](./0007-zero-dependency-core-and-receipt-port.md) | Zero-dependency core; OTel exporter behind the Receipt port | accepted |
| [0008](./0008-adopt-opentelemetry-genai-conventions.md) | Adopt the OpenTelemetry GenAI semantic conventions | accepted |
| [0009](./0009-reserve-then-commit-token-budgeting.md) | Reserve-then-commit token budgeting | accepted |
| [0010](./0010-provider-spi-must-report-token-usage.md) | Minimal Provider SPI; usage reporting is mandatory | accepted |
| [0011](./0011-honest-instrumentation.md) | Honest instrumentation: unmetered currencies declare themselves | accepted |
| [0012](./0012-enforce-architecture-decisions-in-ci.md) | Enforce architecture decisions in CI | accepted |

*(0004 is intentionally unused. It was reserved during planning for a decision that
was folded into ADR-0003. The number is retired rather than recycled, so that
references made while it existed do not silently point at something else.)*

## Disagreeing with a decision

You are encouraged to. Open an issue using the **"Challenge an architecture
decision"** template. If the challenge succeeds, the outcome is a **new ADR that
supersedes the old one** - never an edit to the original.

Some decisions are enforced by CI (see [ADR-0012](./0012-enforce-architecture-decisions-in-ci.md)).
A failing guard is not a wall; it is a pointer to the decision record and to the
process for changing it.