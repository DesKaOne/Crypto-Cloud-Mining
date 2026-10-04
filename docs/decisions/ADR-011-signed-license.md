# ADR-011: Signed License Verification

**Status:** Accepted

Official license payloads and activation artifacts may be digitally signed by vendor infrastructure.

Customer source code may contain the verification mechanism and public key required to validate official artifacts, but the vendor private signing key must remain outside the customer repository and deployment.

The goal is authenticity, entitlement integrity, and traceability, not impossible prevention of source-code modification by a customer who owns the source.
