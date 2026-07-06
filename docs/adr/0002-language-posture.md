---
status: "accepted"
date: 2026-07-06
deciders: [The EAPRA Authors]
---

# 0002. Language posture: Go core, polyglot plugins over a network SPI

## Context and Problem Statement

The EAPRA gateway core sits on the request hot path and belongs to the
cloud-native infrastructure world, where Go is the lingua franca (Kubernetes,
Envoy, Prometheus, the OpenTelemetry Collector). Its plugins provider
adapters, re-rankers, guardrails pull toward Python, where the AI ecosystem
lives. What language(s) do we use, and is the plugin boundary in-process or
across a network?

## Decision Drivers

- Hot-path performance and concurrency for the core.
- Fit with the CNCF ecosystem the reference is meant to teach.
- Access to the AI/provider ecosystem, which is predominantly Python.
- Avoiding runtime coupling: one plugin's language must not dictate the core's.
- Contributor accessibility at the edges.

## Considered Options

- All-Go (core and plugins)
- All-Python (core and plugins)
- Go core + polyglot plugins behind a **network** SPI
- Go core + **in-process** (embedded) plugins

## Decision Outcome

Chosen option: "Go core + polyglot plugins behind a network SPI", because it
puts each concern in the language that serves it best while keeping the plugin
boundary a language-agnostic network contract. A provider adapter is a small
service the gateway calls; its implementation language is then irrelevant to the
core.

### Consequences

- Good, because the hot-path core stays in Go while AI-ecosystem plugins can be
  Python (or anything) without dragging the core into another runtime.
- Good, because a network boundary is a clean, versionable contract and lets
  plugins scale and deploy independently.
- Bad, because a network hop adds latency versus in-process calls, and CI must
  cover two toolchains. (Accepted; see "Revisit When".)

## Pros and Cons of the Options

### All-Go

- Good, because a single toolchain and maximum hot-path performance.
- Bad, because it fights the Python-centric AI ecosystem at the edges.

### All-Python

- Good, because it matches the AI ecosystem and lowers contributor friction.
- Bad, because it is a weaker fit for a high-throughput gateway and idiomatic
  Kubernetes controllers.

### Go core + polyglot plugins over a network SPI

- Good, because right-tool-per-layer with no runtime coupling.
- Good, because the SPI is an explicit, testable contract.
- Bad, because of the per-call network hop and a two-toolchain CI matrix.

### Go core + in-process plugins

- Good, because it avoids the network hop.
- Bad, because a Python plugin would force the core to embed a Python runtime 
  exactly the coupling we want to avoid.

## Revisit When

Revisit if the Go surface stays small enough that collapsing to Python-first
would lower contributor friction, or if the network hop's latency proves
unacceptable for a co-located plugin in which case allow in-process **Go**
plugins as an opt-in optimization without reopening the polyglot boundary.