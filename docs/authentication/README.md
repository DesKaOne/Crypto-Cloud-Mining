# Authentication & Identity

Phase 0 establishes authentication and identity boundaries without implementing a production identity provider.

## Identity model

A User is the stable internal identity.

External identities are linked records:
- Telegram
- email/password or email verification
- OAuth/OpenID Connect providers
- future providers

An external provider identifier is never the primary domain User ID.

## Authentication vs authorization

Authentication answers who the identity is. Authorization answers what the identity may do.

Authentication produces an authenticated principal. Application use cases remain responsible for authorization decisions.

## Token boundary

HTTP APIs use bearer authentication at the transport boundary.

The authentication adapter validates the token and produces a server-side principal containing:
- internal user ID
- authentication method/provider
- session/token identity
- relevant authentication timestamps

Business code does not parse JWT claims directly.

## Telegram

Telegram Mini App authentication is treated as an external identity proof.

The backend must validate Telegram signed initialization data server-side before creating or resolving a User identity.

The client must never be trusted to declare its own Telegram user ID.

## Account linking

Multiple external identities may map to one User when explicitly linked.

A provider identity must map to at most one User.

## Session security

- access credentials are short-lived where practical;
- refresh/session credentials are revocable;
- secrets are never stored in source control;
- authentication failures do not reveal whether another account exists.

Provider-specific token issuance, refresh endpoints, Telegram verification implementation, password storage, MFA, and production key management are later implementation work.
