# Commercial Product & License Foundation

Phase 0 adds a commercial product boundary because the source code is intended to be sold as a software product.

The goal is architecture, domain vocabulary, persistence boundaries, security boundaries, and release/version rules. Payment, billing, marketplace, licensing portal, and aggressive DRM remain future work.

## Domain

- Product
- ProductVersion
- Release
- Edition
- Customer
- License
- LicenseActivation
- Installation
- Entitlement
- Upgrade

## Product hierarchy

    Product
       |
       +-- ProductVersion
               |
               +-- Release

Semantic versioning is the baseline:

    MAJOR.MINOR.PATCH

Version is domain data, not an arbitrary string duplicated across source files.

## Ownership

A license grants an entitlement to a purchased version/release according to its policy.

Purchasing v0.0.1 does not automatically imply entitlement to v0.0.2. Upgrade eligibility is policy/entitlement driven.

Commercial pricing is intentionally not defined in Phase 0.

## Editions

The architecture supports editions such as Standard, Professional, and Enterprise. Editions resolve capabilities through entitlements rather than separate source-code forks.

Feature-gating implementation is future work.

## License

A license is a domain object, not a magic key string. It may include:

- customer
- product
- purchased version
- edition
- status
- activation limit
- issued timestamp
- expiration timestamp

## Installation and activation

Installation identifies a customer deployment. Activation associates an installation with a license according to activation policy.

Online activation is supported by the architecture. Offline activation is a future capability using a signed activation artifact.

## Signed license

The intended trust flow is:

    License payload
        |
        v
    Vendor signing service
        |
        v
    Digital signature
        |
        v
    License token/file
        |
        v
    Customer installation
        |
        v
    Public-key verification
        |
        v
    Entitlement resolution

The vendor private signing key must never be shipped in customer source code.

## Vendor boundary

Customer application:

- API
- Worker
- Admin
- Frontend
- Database

Vendor infrastructure:

- License Server
- Customer Portal
- Release Management
- Download Management
- Upgrade Management

These are separate bounded systems.

## Update vs upgrade

Technical releases follow semantic versioning. Commercial upgrade eligibility is a separate entitlement decision.

Example:

    v0.0.1 -> v0.0.2  patch release
    v0.0.x -> v0.1.0  minor release
    v0.x.x -> v1.0.0  major release

This does not define pricing policy.
