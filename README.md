# Enterprise AI Platform Reference Architecture (EAPRA)

## What is EAPRA?

EAPRA shows you how to build a real, production-ready AI system and gives you
the working code to prove it.

Think of it as a **worked example**. Instead of describing how the pieces of an
enterprise AI system should fit together, EAPRA actually wires them together so
you can run it, read it, and learn from it. It combines trusted, existing tools
(the same ones large companies rely on) and adds a thin layer of its own code to
connect them into one working whole.

It is the runnable companion to **Whitepaper 001**. Reading the whitepaper and
cloning this repo should feel like the same idea, explained two ways: one in
words, one in code.

**What EAPRA is *not*:**

- **Not a framework.** It doesn't take over your application. It works underneath
  it, at the infrastructure level.
- **Not a finished product.** It's built to *teach*, not to run your business as-is.
- **Not a rebuild of existing tools.** It uses proven projects like OpenTelemetry,
  KServe, vLLM, and Vault it doesn't replace them. Its own code is the glue that
  connects them, plus the explanations of *why* they're connected that way.

## Why EAPRA exists

There's a big gap between an AI *demo* and an AI system that actually runs in
production and no clear, trustworthy guide for crossing it.

Almost anyone can build an AI demo in an afternoon. But running one for real is
much harder. You suddenly need to handle logins, cost limits, request routing,
answer quality, monitoring, security, and more. The knowledge to do this exists,
but it's scattered across company blogs, private chats, and lessons nobody writes
down.

There's no single, reliable place the way there is for Kubernetes where an
engineer can see a complete, correct AI system working in code they can actually
run.

EAPRA aims to be that place: turning hard-won, real-world experience into
something you can clone and study. It also backs its claims with **real
measurements** - actual cost and speed numbers from running the system, with the
method shown openly, not marketing.

## Who this is for

- **Platform and DevOps engineers** setting up AI at their company a ready
  starting point instead of a blank page.
- **Architects** clear records of every design decision, so you can defend your
  choices in a review.
- **Engineers growing into AI infrastructure** a step-by-step path with real,
  connected code instead of scattered tutorials.
- **Teachers, workshop leaders, and speakers** accurate, repeatable material for
  labs and demos.
- **Open-source contributors** a small, stable core that's easy and safe to
  extend (like adding a new provider) in a single, self-contained change.

**Not for:** teams wanting a ready-to-sell product, people looking for an app or
model framework, or researchers focused on training and fine-tuning models.

> **Project status:** EAPRA is in its early foundation phase. Right now this
> repository shows *where the project is headed* - its direction and design - not
> a finished, adopted system.

## Architecture overview

EAPRA's architecture is organized around a small set of core concepts - **three
planes**, **two envelopes**, and the **Five Currencies**. These are being
finalized in **Whitepaper 001**, which is the single source of truth for the
model.

To avoid two documents drifting apart, this section intentionally does **not**
define those concepts yet. The full architecture - with diagrams and the
executable mapping - will land here once the whitepaper is locked. See
[`docs/architecture/`](./docs/architecture) for progress.

## License

Licensed under the [Apache License 2.0](./LICENSE).
Copyright 2026 The EAPRA Authors.