# AI Gateway (walking skeleton)

The decide-and-meter control layer: EAPRA's original code (ADR-0003).
Zero third-party dependencies (ADR-0007).

## Run it in 30 seconds - no API key required

```bash
make run          # starts on :8080 with the stub provider
```

In another terminal:

```bash
curl -s -X POST localhost:8080/v1/chat/completions \
  -H 'Content-Type: application/json' \
  -H 'X-EAPRA-Caller: you' \
  -d '{"model":"demo","messages":[{"role":"user","content":"What is EAPRA?"}]}'
```

You get a completion **and**, on the gateway's stdout, a **Request Receipt**:

```json
{
  "trace_id": "37029f90c1a97a0ae78786303ae8bfd8",
  "caller": "you",
  "gen_ai.provider.name": "stub",
  "gen_ai.request.model": "demo",
  "gen_ai.usage.input_tokens": 4,
  "gen_ai.usage.output_tokens": 17,
  "currencies": {
    "latency_ms": 12,
    "cost_tokens": 21,
    "trust": "not_metered",
    "freshness": "not_metered",
    "agency": "not_metered"
  }
}
```

Three currencies say `not_metered`. That is deliberate. See
[the architecture](../../docs/architecture/README.md#3-the-five-currencies-of-production-ai).

## Watch the token budget stop a Denial-of-Wallet

Budgets are counted in **tokens, not requests** - because one prompt can cost
orders of magnitude more than another (OWASP LLM10).

```bash
EAPRA_TOKEN_BUDGET=60 make run
# repeat the curl above a few times -> HTTP 429 {"error":{"type":"token_budget_exceeded"}}
```

The budget is **reserved before the provider is called**, so an over-budget
caller never spends a single token. Budgets are per-caller: a different
`X-EAPRA-Caller` is unaffected.

## Point it at a real model

The same adapter speaks to OpenAI, vLLM, Ollama, and anything else exposing an
OpenAI-compatible `/chat/completions` (ADR-0004).

```bash
export EAPRA_PROVIDER=openai
export EAPRA_OPENAI_BASE_URL=http://localhost:11434/v1   # e.g. Ollama
export EAPRA_OPENAI_API_KEY=...                          # omit for local servers
make run
```

## Endpoints

| Method | Path | Purpose |
| --- | --- | --- |
| `POST` | `/v1/chat/completions` | Completion through the metered request path |
| `GET` | `/v1/receipts` | The last 50 receipts, so you can see them without a tracing backend |
| `GET` | `/healthz` | Liveness |

## Configuration

| Variable | Default | Meaning |
| --- | --- | --- |
| `EAPRA_ADDR` | `:8080` | Listen address |
| `EAPRA_PROVIDER` | `stub` | `stub` or `openai` |
| `EAPRA_TOKEN_BUDGET` | `1000` | Per-caller token budget; `0` = unlimited |
| `EAPRA_OPENAI_BASE_URL` | `https://api.openai.com/v1` | Any OpenAI-compatible base URL |
| `EAPRA_OPENAI_API_KEY` | *(empty)* | Omit for local servers |

## Known gaps (stated, not hidden)

- Receipts go to stdout and an in-memory recorder. **The OpenTelemetry exporter
  is not wired yet** (ADR-0007) - the conventions are, the exporter is issue #1.
- Token counting for the stub uses a crude ~4-chars/token heuristic. A real
  tokenizer is issue #2.
- The caller identity is a header, not a validated token. OIDC/JWT is issue #3.
- Budgets are process-local and reset on restart. A Redis-backed store is issue #4.

`make test` covers the request path, the budget breach, per-caller isolation,
and the receipt's contents.