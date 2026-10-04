# API Error Codes

Initial stable error-code vocabulary:

| Code | Meaning |
|---|---|
| UNAUTHENTICATED | No valid authentication context |
| FORBIDDEN | Authenticated but not authorized |
| VALIDATION_ERROR | Request shape or field validation failed |
| NOT_FOUND | Requested resource does not exist |
| CONFLICT | State or concurrency conflict |
| IDEMPOTENCY_CONFLICT | Same idempotency key used with different operation/input |
| RATE_LIMITED | Request rate exceeded |
| INTERNAL_ERROR | Unexpected server failure |

The HTTP status and machine-readable code are both part of the contract.
