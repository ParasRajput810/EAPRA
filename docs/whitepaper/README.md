# Whitepapers

The *Engineering Enterprise AI Infrastructure* series. These papers are the
**single source of truth for EAPRA's architecture**. Where the code and a paper
disagree, the paper is correct and the code is a bug.

| # | Title | Status |
| --- | --- | --- |
| [001](./WP-001-beyond-the-ai-demo.md) | Beyond the AI Demo - A Platform Engineering Guide to Building Secure, Scalable, and Observable Enterprise AI Systems | v0.1 draft, living |

## How these papers work

- **They are versioned like code.** Changes arrive through pull requests and are
  reviewed like any other change.
- **They declare their own gaps.** Whitepaper 001 contains a section
  ("What EAPRA implements today") that lists, unflatteringly, every concept the
  paper describes and the repository does not yet implement. That table exists so
  the paper and the code cannot quietly drift apart.
- **They invent no measurements.** No latency figure, cost figure, or throughput
  figure appears unless it was measured and its methodology published. None have
  been yet.
- **They cite official sources.** CNCF, Kubernetes SIG-Network, OpenTelemetry, and
  OWASP - not blogs.
- **They label illustration as illustration.** Scenarios marked *Illustrative* are
  constructed to show how the architecture behaves. They are not reports of real
  deployments, and the papers do not claim production experience the author does
  not have.

## Disagreement

If a claim in a paper is wrong, unsupported, or overstated, open an issue -
especially if it is a claim the author would prefer to keep.