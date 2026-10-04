# Telegram Mini App Authentication

Telegram is an external identity provider for the platform.

## Trust boundary

Telegram Mini App
       |
       | signed init data
       v
Backend authentication adapter
       |
       | validate signature + freshness
       v
ExternalIdentity(provider=telegram)
       |
       v
User

## Rules

1. The Mini App never directly chooses the authenticated UserID.
2. Telegram initialization data is validated on the backend.
3. Freshness/expiration rules are enforced server-side.
4. The validated Telegram subject is mapped to ExternalIdentity.
5. Existing identity resolves to its User; otherwise account provisioning creates one.
6. Telegram authentication does not grant administrative permissions by itself.

Exact verification code and key management are later implementation steps.
