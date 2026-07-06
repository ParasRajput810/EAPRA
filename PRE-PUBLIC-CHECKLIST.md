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
- [ ] **Reformat the remaining kickoff ADRs into MADR.** ADR-0001-0004 (language
      posture, build-vs-integrate, first adapter) are now authored directly in
      MADR. Reformat the remaining kickoff decisions (0005, 0006) into the
      canonical template so `docs/adr/` stays single-format.

## Deferred by dependency

- [ ] **Write the full `CONTRIBUTING.md`.** Three forks now confirmed (language
      posture, build-vs-integrate, first adapter), so this is unblocked pending
      the walking skeleton for accurate dev-setup and test instructions.