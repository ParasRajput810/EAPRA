---
status: "accepted"
date: 2026-07-06
deciders: [The EAPRA Authors]
---

# 0003. Build vs. integrate: build the decide-and-meter control layer, integrate every substrate

## Context and Problem Statement

EAPRA's component list is a multi-year roadmap. Without a crisp boundary between
the code we author and the projects we integrate, the project sprawls, loses
credibility (by reinventing mature tools), and never finishes. Where exactly is
the line between "build" and "integrate and never rebuild"?

## Decision Drivers

- Finishability: the tighter "build" is drawn, the more achievable the project.
- Credibility: reinventing CNCF-grade tooling reads as naïve, not ambitious.
- Maintainability: a small original surface is a maintainable one.
- Teaching value: an enterprise AI platform is *assembled*, not authored from
  scratch - the assembly is the lesson.

## Considered Options

- Build most of it (own proxy, telemetry pipeline, secrets store, serving)
- Integrate everything; write almost no original code
- Build only the "decide-and-meter" control layer; integrate all substrate

## Decision Outcome

Chosen option: "Build only the decide-and-meter control layer; integrate all
substrate."

- **Build (original code):** the gateway control logic - token metering and
  budgets, routing, the plugin SPI contract, and the request-receipt wiring.
- **Integrate and never rebuild:** OpenTelemetry (telemetry), the model
  providers, Envoy / Gateway API (edge), HashiCorp Vault (secrets),
  Prometheus / Grafana / Jaeger (telemetry backend), and - later - the
  Kubernetes inference gateway.

**Tripwire:** if you find yourself writing a metrics store, a proxy, or a
secrets manager, you have crossed the line - stop and integrate instead.

### Consequences

- Good, because the maintained surface stays small, credible, and finishable.
- Good, because "assemble proven parts" is itself the reference's core lesson.
- Bad, because EAPRA depends on upstream projects and their changes, and forgoes
  "we built it all" bragging rights. (Accepted.)

## Pros and Cons of the Options

### Build most of it

- Good, because maximum control and no external dependencies.
- Bad, because unmaintainable, not credible, and effectively never ships.

### Integrate everything

- Good, because minimal maintenance.
- Bad, because there is no original contribution - nothing to study or teach.

### Build only the control layer

- Good, because it isolates the one genuinely original, high-value component.
- Good, because it maps cleanly onto the walking-skeleton milestone.
- Bad, because the boundary requires discipline to hold (hence the tripwire).

## Revisit When

Revisit a specific integration only when an upstream project cannot express
something the reference genuinely needs - and even then, contribute upstream
before forking or rebuilding.