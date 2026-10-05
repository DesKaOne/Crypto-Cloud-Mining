# Authentication Architecture Decisions

## ADR-AUTH-0001: External identities map to internal Users

Provider-specific identifiers are external identity records. They never become the primary domain User ID.

## ADR-AUTH-0002: Authentication and authorization are separate

Transport/authentication adapters establish an authenticated principal. Application use cases authorize actions.

## ADR-AUTH-0003: Telegram identity is backend-verified

Telegram Mini App initialization data must be validated server-side before identity resolution or provisioning.

These decisions complement the core architecture decisions without coupling the domain to a provider SDK.
