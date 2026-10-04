# ADR-012: Customer Source-Code Model

**Status:** Accepted

The product is sold with source code. The architecture does not assume customer source can be made immutable.

Security focuses on:
- protecting official release authenticity
- protecting license authenticity
- validating entitlements
- tracing activations
- controlling upgrade eligibility

Customer modifications are technically possible and are not treated as a solvable DRM problem.
