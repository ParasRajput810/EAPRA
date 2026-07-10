---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0005. Plane-aligned repository structure over strict Domain-Driven Design

## Context and Problem Statement

EAPRA is a teaching artifact as much as a running system. Its directory tree is
the first thing a visitor reads. How should the repository be organized so the
structure itself communicates the architecture?

## Decision Drivers

- The repo should teach the model, not obscure it.
- Contributors must know instantly where a new component belongs.
- Structure should mirror Whitepaper 001 so the paper and the code never drift.
- Navigability for a newcomer beats theoretical purity.

## Considered Options

- Strict Domain-Driven Design folders (bounded contexts, aggregates).
- Layer-based folders (`api/`, `services/`, `db/`).
- Plane-aligned folders mirroring the whitepaper (`request-plane/`, `control-plane/`, `data-plane/`).

## Decision Outcome

Chosen option: **plane-aligned folders**. The top level is `request-plane/`,
`control-plane/`, and `data-plane/`, matching the Three Planes of the
whitepaper. A reader who has read the paper can navigate the repo without a map,
and a reader who starts with the repo learns the model from the tree.

DDD's *tactical* patterns remain welcome **inside** a component. This decision
concerns strategic layout only.

### Consequences

- Good, because the repository teaches the architecture by its shape.
- Good, because the Blast-Radius Boundary between planes becomes a real directory boundary.
- Good, because a new component has one obvious home.
- Bad, because `control-plane/` and `data-plane/` sit empty early on, which can
  read as incompleteness. Accepted: the empty directories are honest signposts,
  and each carries a README explaining what will live there.

## Pros and Cons of the Options

### Strict DDD folders

- Good, because it is rigorous and familiar to some enterprise teams.
- Bad, because the ceremony (bounded contexts, ubiquitous-language scaffolding)
  makes a teaching repo harder to navigate, not easier.

### Layer-based folders

- Good, because conventional and immediately familiar.
- Bad, because it teaches nothing about *this* system; any CRUD app looks the same.

### Plane-aligned folders (chosen)

- Good, because structure and mental model become one artifact.
- Bad, because empty planes early on must be explained rather than hidden.

## Revisit When

A plane grows large enough to warrant its own repository, or the whitepaper's
plane model itself changes (which would supersede this ADR, not amend it).