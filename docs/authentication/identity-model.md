# Identity Model

## Entities

User:
- stable internal ID
- lifecycle status
- account timestamps

ExternalIdentity:
- internal ID
- user ID
- provider
- provider subject
- verified timestamp
- created timestamp

AuthenticationSession:
- internal ID
- user ID
- session status
- created/expired/revoked timestamps

## Uniqueness

The persistence layer must enforce uniqueness for:
- provider + provider subject
- session identifier

A disabled/closed User cannot authenticate into normal product activity.

## Principal

After successful authentication, transport passes a principal to the application layer:

UserID + provider + session identity

Application code uses UserID for authorization and domain operations.
