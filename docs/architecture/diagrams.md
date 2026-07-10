# Diagrams

Diagrams render on GitHub. Source lives beside them so they can be edited rather
than replaced.

---

## 1. The architecture: three planes, two envelopes, five currencies

![EAPRA architecture](../assets/architecture.svg)

The two envelopes - platform foundation, and governance and telemetry - wrap
**every** stage of the request path. They are not stages in it. Drawing them as
pipeline steps is how they end up built as afterthoughts.

`control-plane/` and `data-plane/` are drawn dashed because they are named and
empty. They exist in the tree so the model is visible; see
[ADR-0005](../adr/0005-plane-aligned-repository-structure.md).

---

## 2. The request lifecycle

What actually happens on `POST /v1/chat/completions`. Note the ordering: the
budget is **reserved before the provider is called**, so a caller who is out of
budget never spends a token ([ADR-0009](../adr/0009-reserve-then-commit-token-budgeting.md)).

```mermaid
sequenceDiagram
    autonumber
    participant C as Client
    participant G as AI Gateway
    participant B as Meter / Budget
    participant P as Provider (SPI)
    participant R as Receipt Emitter

    C->>G: POST /v1/chat/completions
    G->>G: Assign trace_id
    G->>G: Identify caller

    Note over G,B: Gate on the estimate...
    G->>B: Reserve(caller, estimated_tokens)

    alt Budget exhausted
        B-->>G: ErrBudgetExceeded
        G->>R: Emit receipt (cost_tokens = 0)
        G-->>C: 429 token_budget_exceeded
        Note right of C: Zero tokens spent.<br/>Provider never called.
    else Budget available
        B-->>G: OK
        G->>P: Complete(request)
        P-->>G: content + usage

        Note over G,B: ...bill the truth.
        G->>B: Commit(caller, actual_tokens)
        G->>R: Emit Request Receipt
        G-->>C: 200 + content + trace_id
    end
```

---

## 3. The Request Receipt

Every request returns one. Attribute names follow the OpenTelemetry GenAI
semantic conventions ([ADR-0008](../adr/0008-adopt-opentelemetry-genai-conventions.md))
so the telemetry is portable rather than a private dialect.

```mermaid
flowchart LR
    subgraph Receipt["Request Receipt (one per request)"]
        direction TB
        ID["trace_id · caller · timestamp"]
        GA["gen_ai.provider.name<br/>gen_ai.request.model<br/>gen_ai.usage.input_tokens<br/>gen_ai.usage.output_tokens"]
        CUR["currencies"]
    end

    CUR --> L["latency_ms<br/><b>metered</b>"]
    CUR --> CO["cost_tokens<br/><b>metered</b>"]
    CUR --> T["trust<br/>not_metered"]
    CUR --> F["freshness<br/>not_metered"]
    CUR --> A["agency<br/>not_metered"]
```

Three currencies say `not_metered` rather than reporting a comfortable zero. A
dashboard averaging `trust: 0` would report perfect safety; a dashboard
encountering `not_metered` cannot
([ADR-0011](../adr/0011-honest-instrumentation.md)).

Closing each one is an open issue. That is the roadmap.

---

## Editing these

- The SVG in `docs/assets/architecture.svg` is hand-written and diffable. Edit the
  markup; do not replace it with an exported bitmap.
- Mermaid blocks render natively on GitHub. Keep them in this file rather than
  exporting images, so a reviewer can diff a diagram change.