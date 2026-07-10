---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0012. Enforce architecture decisions in CI

## Context and Problem Statement

ADR-0007 commits the gateway core to having no third-party dependencies.
ADR-0003 draws a line between what EAPRA builds and what it integrates. Both are
prose. Prose does not stop a pull request. Decisions recorded only in documents
decay quietly: the first violating change is small, plausible, and merged, and
the ADR becomes a description of a past the code has left.

How do we keep an architecture decision true?

## Decision Drivers

- A decision nobody can violate accidentally is a decision that survives.
- The project teaches decision discipline; a repository that documents rules it
  does not enforce teaches the opposite.
- Enforcement must explain itself, or it reads as bureaucracy to a newcomer.
- Rules must remain changeable - enforcement must not become ossification.

## Considered Options

- Rely on code review to catch violations.
- Rely on documentation and contributor good faith.
- Encode enforceable decisions as CI checks that fail the build and cite the ADR.

## Decision Outcome

Chosen option: **encode enforceable decisions as CI checks.**

Two rules are enforced today:

1. **The dependency guard.** If a `require` directive appears in the gateway's
   `go.mod`, CI fails with a message stating that ADR-0007 requires a
   dependency-free core, that adding a dependency is *allowed*, and that doing so
   requires its own ADR.
2. **DCO sign-off.** Every commit in a pull request must carry `Signed-off-by`.

The failure message is part of the design. The check does not say "denied"; it
says *"this is permitted, and here is the process."* Enforcement points at the
decision record rather than replacing it.

### Consequences

- Good, because ADR-0007 cannot decay by accident. Violating it is now a conscious
  act with a paper trail.
- Good, because a contributor discovers the architecture through the tooling, at
  the moment it is relevant, rather than by reading every ADR beforehand.
- Good, because "the philosophy is a job that goes red" is a demonstrable claim.
- Bad, because CI now encodes opinions, and a wrong opinion is now expensive to
  hold. Mitigated by the failure message, which routes disagreement into an ADR
  rather than into a fight with the build.
- Bad, because not every decision is mechanically checkable. Most are not. This
  creates a temptation to over-value the decisions that happen to be enforceable.
  Enforceability is not importance.

## Pros and Cons of the Options

### Code review only

- Good, because flexible and requires no tooling.
- Bad, because it depends on a reviewer remembering every ADR, and it does not
  scale past one attentive maintainer.

### Documentation and good faith

- Good, because zero friction.
- Bad, because it is how architecture decisions actually die.

### CI enforcement (chosen)

- Good, because the rule holds whether or not anyone is watching.
- Bad, because it privileges checkable rules and makes wrong rules costly.

## Revisit When

A guard blocks work that is genuinely correct - in which case supersede the
underlying ADR rather than weakening the check quietly. Any new guard should be
added only alongside the ADR it enforces.