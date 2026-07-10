---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0006. Walking skeleton before breadth

## Context and Problem Statement

EAPRA's component list -- gateway, auth, routing, registries, evaluation,
caching, retrieval, agents, observability -- is a multi-year roadmap. Building
breadth-first produces a sprawling repository of stubs that runs nothing and
teaches nobody. What is built first?

## Decision Drivers

- Finishability. Ten components that run beat forty that do not.
- Contributors engage with things that **run**, not with foundations.
- The architecture must be proven end to end before it is widened.
- A reference must emit honest telemetry from its first request.

## Considered Options

- Breadth-first: scaffold every component as a stub, fill them in later.
- Depth-first on one component (e.g. a complete AI gateway) before any other.
- A walking skeleton: one thin end-to-end slice through the whole request path.

## Decision Outcome

Chosen option: **a walking skeleton**. Milestone one is a single request
traversing the full path -- caller → token budget → provider → Request Receipt --
running under one command, covered by tests, with no API key required.

Every later feature slots into a working spine rather than into a diagram.

### Consequences

- Good, because the architecture and the delivery loop are proven on day one.
- Good, because a first-time visitor can run the project in under a minute,
  which is the single strongest predictor of whether they contribute.
- Good, because each subsequent feature is a small, reviewable addition.
- Bad, because the skeleton is unimpressive in feature count and may read as
  trivial to someone counting components rather than reading the receipt.

## Pros and Cons of the Options

### Breadth-first stubs

- Good, because the repo *looks* complete quickly.
- Bad, because nothing runs; stubs rot; contributors have nothing to pull on.

### Depth-first on one component

- Good, because one component reaches production quality early.
- Bad, because the end-to-end path -- where the interesting failures live --
  stays unproven.

### Walking skeleton (chosen)

- Good, because it proves the whole path with the least code.
- Bad, because breadth arrives slowly and must be defended against impatience.

## Revisit When

The skeleton is complete and stable enough that parallel workstreams on separate
planes no longer risk invalidating the spine.