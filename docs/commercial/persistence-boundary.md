# Commercial Persistence Boundary

Phase 0 does not add commercial tables to the core migration yet.

The architecture reserves persistence boundaries for:

    products
    product_versions
    releases
    editions
    customers
    licenses
    license_activations
    installations
    entitlements
    upgrades

Before a migration is introduced, each aggregate must have:

- ownership rules
- lifecycle states
- uniqueness constraints
- audit requirements
- idempotency requirements where mutations are retryable
- relationship to the vendor/customer trust boundary

The commercial domain must not make the mining ledger depend on license tables for financial correctness. Licensing authorizes software usage; the financial ledger remains its own source-of-truth boundary.
