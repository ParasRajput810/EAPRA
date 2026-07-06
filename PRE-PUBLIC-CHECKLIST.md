# Pre-Public Checklist

Honest stopgaps taken during early development that must be resolved before
or shortly after the first public announcement. This file exists so that
temporary answers do not silently become permanent.

## Must resolve before / around launch

- [ ] **Confirm & monitor the Code of Conduct contact.** `paraswork810@gmail.com`
      is committed publicly and will be scraped. Confirm it is monitored (or
      forwarded to an inbox that is read).
- [ ] **Graduate the CoC contact to a role alias.** Once EAPRA has more than one
      maintainer, replace the personal address with `conduct@<domain>` so reports
      outlive any single person.
- [ ] **Land the Architecture section.** Replace the placeholder in `README.md`
      and `docs/architecture/` once Whitepaper 001 finalizes the three planes,
      two envelopes, and Five Currencies.
- [ ] **Reformat ADR-0001…0005 into MADR.** Re-express the five kickoff ADRs in
      the canonical MADR template so `docs/adr/` has one consistent format.

## Deferred by dependency

- [ ] **Write the full `CONTRIBUTING.md`.** Held until three forks are confirmed:
      language posture, build-vs-integrate, and the first adapter.