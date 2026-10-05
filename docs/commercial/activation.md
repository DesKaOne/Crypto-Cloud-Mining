# Activation Boundary

## Online activation

    Customer Installation
           |
           v
    Vendor License Service
           |
           v
    License validation
           |
           v
    Activation decision
           |
           v
    Signed activation result
           |
           v
    Customer Installation

## Offline activation

    Customer
       |
       +-- License Key
       +-- Installation ID
              |
              v
       Vendor License Service
              |
              v
       Signed Activation File
              |
              v
       Customer Server
              |
              v
       Local Public-Key Verification

Offline activation lifetime, revocation semantics, and clock policy are future implementation decisions.

The customer installation never receives the vendor private signing key.
