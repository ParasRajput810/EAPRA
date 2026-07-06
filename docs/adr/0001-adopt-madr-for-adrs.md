---
status: "accepted"
date: 2026-07-06
deciders: [Paras Rajput]
---

# 0001. Adopt MADR for Architecture Decision Records

## Context and Problem Statement

EAPRA will accumulate many architectural decisions, and the reasoning behind
each must survive long after the moment it was made for future contributors
and for future-me. We need a consistent, lightweight, version-controlled way to
record decisions. What format should ADRs use?

## Decision Drivers

- Decisions must live alongside the code, be diffable, and be reviewable in PRs.
- The format must be light enough that a solo maintainer will actually use it.
- It should match conventions open-source contributors already recognize.

## Considered Options

- MADR (Markdown Any Decision Records)
- Michael Nygard's original ADR format
- No formal ADRs (decisions live in commit messages and issues)

## Decision Outcome

Chosen option: "MADR", because it is Markdown-native, widely used across the
CNCF ecosystem, and structured (drivers, options, consequences) without being
heavy. It gives every decision one consistent, reviewable home.

### Consequences

- Good, because every decision has a single, consistent, discoverable location.
- Good, because the format is familiar to the contributors we hope to attract.
- Bad, because it adds a small authoring step per decision an accepted cost.

## Pros and Cons of the Options

### MADR

- Good, because it is Markdown, diffable, and widely adopted.
- Good, because it structures the comparison of options and their consequences.
- Bad, because it carries slightly more ceremony than a commit message.

### Nygard ADR

- Good, because it is minimal and quick to write.
- Bad, because it offers less structure for weighing options against each other.

### No formal ADRs

- Good, because it has zero overhead.
- Bad, because the "why" is lost; future contributors cannot reconstruct intent.

## Revisit When

Revisit if the ceremony proves too heavy for a solo maintainer in practice, or
if EAPRA later adopts a documentation toolchain with a native decision log.
