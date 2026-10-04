# Phase 0 Threat Model

| Threat | Boundary | Phase 0 control |
| --- | --- | --- |
| Forged API identity | API/auth | Backend authentication and principal boundary |
| Unauthorized domain action | Application | Use-case authorization |
| Reward manipulation | Domain | Server-authoritative calculation inputs |
| Ledger tampering | Persistence | Append-oriented ledger + idempotency |
| Secret leakage | Deployment | Environment/secret-manager boundary |
| License forgery | Commercial | Signed license architecture |
| License replay | Commercial | Activation/session identity and policy boundary |
| Activation abuse | Commercial | Installation + activation records |
| Private signing-key theft | Vendor | Key kept outside customer repository |
| Source modification | Customer installation | Verify official release/license artifacts |

Phase 0 does not claim to prevent a customer with source access from modifying their own installation.
