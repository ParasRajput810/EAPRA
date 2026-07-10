---
status: "accepted"
date: 2026-07-10
deciders: [The EAPRA Authors]
---

# 0009. Reserve-then-commit token budgeting

## Context and Problem Statement

The gateway enforces a per-caller token budget. Tokens are only known exactly
*after* the provider responds - but by then the money is spent. If we check the
budget only on the way out, an over-budget caller has already cost us a full
inference. When and how should the budget be enforced?

## Decision Drivers

- An over-budget caller must not be able to spend anything (OWASP LLM10,
  unbounded consumption / Denial of Wallet).
- Budget accounting must reflect reality, not an estimate.
- The check sits on the hot path, so it must be cheap.
- The estimate and the truth will always differ; the design must accept that.

## Considered Options

- Check the budget after the response, using actual usage.
- Check only an estimate, and bill the estimate.
- **Reserve** against an estimate before the call; **commit** the actual usage after.

## Decision Outcome

Chosen option: **reserve-then-commit.**

1. Estimate the request's token cost from the prompt.
2. `Reserve(caller, estimated)` - if it does not fit the remaining budget, reject
   with a typed `429 token_budget_exceeded` **before the provider is called.**
3. Call the provider.
4. `Commit(caller, actual)` - bill the real usage the provider reported.

We *gate* on the estimate and *bill* the truth.

### Consequences

- Good, because a rejected request spends exactly zero tokens. This is verified by
  a test that asserts `cost_tokens == 0` on a budget breach.
- Good, because accounting never drifts from what the provider actually charged.
- Good, because the typed error gives callers something to program against.
- Bad, because a bad estimate can reject a request that would have fitted
  (false positive) or admit one that overshoots (small overshoot past the limit).
  A crude estimator makes both worse - which is exactly why a real tokenizer is a
  tracked issue and why the current heuristic is documented as crude rather than
  presented as exact.
- Bad, because reserve and commit are two operations, so a crash between them
  under-bills. Accepted at this scale; a distributed store must make it atomic.

## Pros and Cons of the Options

### Check after the response

- Good, because it bills perfectly accurately.
- Bad, because the abuse has already happened. It detects Denial of Wallet
  instead of preventing it.

### Estimate only, bill the estimate

- Good, because one operation, simple.
- Bad, because the ledger diverges from the invoice - the one thing a cost SLO
  cannot tolerate.

### Reserve-then-commit (chosen)

- Good, because it prevents spend and bills truthfully.
- Bad, because the estimate's quality bounds its precision, and it is two steps.

## Revisit When

The budget moves to a distributed store, at which point reserve and commit should
become an atomic operation (see the Redis-backed budget issue), or when a real
tokenizer replaces the heuristic and the false-positive rate can be measured.