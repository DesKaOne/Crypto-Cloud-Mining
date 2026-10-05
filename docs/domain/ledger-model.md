# Ledger Model

The ledger is the financial source-of-truth boundary.

## Conceptual accounts

At minimum the design must be able to represent:

- user wallet account
- platform mining/reward liability account
- platform fee/revenue account
- pending withdrawal account
- external settlement account

## Ledger entry

A ledger entry should contain:

- unique ID
- asset ID
- transaction/reference ID
- debit account
- credit account
- quantity
- created timestamp
- idempotency key
- immutable metadata/reference

## Double-entry principle

A balanced financial operation moves the same asset quantity from one account to another. No operation should create or destroy a user balance by directly editing a wallet total.

Database implementation belongs to the next persistence step.
