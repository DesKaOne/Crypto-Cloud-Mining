# Commercial Domain Model

## Product

Owns the commercial identity of the software.

## ProductVersion

Represents a semantic version of a Product.

Responsibilities:
- version identity
- compatibility metadata
- lifecycle status

## Release

Represents a distributable official release for a ProductVersion.

Suggested metadata:
- version
- release type
- release date
- minimum supported version
- release status
- checksum
- artifact reference

## Edition

Defines a named commercial packaging level.

Edition does not own implementation code. It contributes entitlement policy.

## Customer

Represents the commercial owner/customer identity. It is intentionally separate from an end-user User account.

## License

Connects Customer, Product, purchased version, and Edition with lifecycle and activation policy.

## LicenseActivation

Represents an activation event/association for a License and Installation.

## Installation

Represents a customer deployment identity and environment metadata.

## Entitlement

Represents an explicit right granted by a License/Edition, such as:

- multi_asset
- telegram
- white_label
- advanced_mining
- blockchain_integration

## Upgrade

Represents a policy decision from one entitled version/release to another.

Upgrade eligibility must not be hardcoded as version comparisons inside unrelated application code.
