# EAPRA Documentation

Welcome to the documentation for the **Enterprise AI Platform Reference
Architecture (EAPRA)**. This is the home for everything beyond the project
overview in the [root README](../README.md).

> **Status:** Early foundation phase. The structure below is in place; many
> documents are still being written. Concepts marked "forthcoming" are being
> finalized in **Whitepaper 001**, the single source of truth for the architecture.

## Where to start

- **New here?** Read the [root README](../README.md) for what EAPRA is and why.
- **Want the architecture?** See [`architecture/`](./architecture) (in progress).
- **Making a decision or reviewing one?** Browse the [ADRs](./adr).
- **Operating a component?** See the [runbooks](./runbooks).
- **Teaching or learning hands-on?** See the [labs](./labs).

## Documentation map

| Folder | Purpose |
| --- | --- |
| `adr/` | Architecture Decision Records - one file per significant decision, in MADR format. |
| `architecture/` | The system's architecture: the model, diagrams, and how the pieces fit. |
| `runbooks/` | Operational guides - how to run, operate, and recover parts of the system. |
| `guides/` | How-to guides and tutorials for using and extending EAPRA. |
| `labs/` | Hands-on teaching material: workshops, labs, and demo scenarios. |
| `assets/` | Images, diagrams, and other static files referenced by the docs. |

## Conventions

- **Format:** All documentation is written in Markdown.
- **ADRs:** Use the MADR template in [`adr/`](./adr). One decision per file; ADRs are immutable once accepted (supersede, don't edit).
- **Single source of truth:** Architecture concepts are defined once, in Whitepaper 001, and referenced from here - never redefined in two places.
- **Contributing to docs:** See the project's `CONTRIBUTING.md` (forthcoming). Documentation changes follow the same review process as code.