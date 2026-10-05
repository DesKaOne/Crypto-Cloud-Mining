# Reward Model

Phase 0 defines the conceptual calculation contract only.

## Inputs

A reward calculation may depend on:

- asset
- mining plan/version
- mining session
- effective hashrate
- active interval
- reward policy/version
- network or pool inputs when applicable
- deterministic adjustment factors

## Output

The calculator produces a deterministic reward result containing:

- user/session reference
- asset reference
- calculation period/reference
- quantity
- calculation policy/version
- idempotency key

## Constraints

- Never trust client-provided elapsed time or reward amount.
- Never calculate monetary value as the source of truth for crypto quantity.
- Do not silently overwrite a previously posted reward.
- Separate crypto quantity calculation from fiat display/conversion.
- Asset/network-specific calculations must be isolated behind adapters or policies.

No production formula is selected in Phase 0.
