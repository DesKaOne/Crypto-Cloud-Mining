# Repository Structure

```text
.
├── apps/
│   ├── flutter_client/        # Mobile + Web Flutter application
│   └── telegram_miniapp/      # Telegram-specific Mini App surface
├── backend/
│   ├── api/                   # HTTP/WebSocket entrypoint
│   ├── worker/                # Background processing entrypoint
│   └── internal/
│       ├── domain/             # Core business domains
│       ├── application/        # Use cases / orchestration
│       ├── infrastructure/     # DB, cache, external adapters
│       └── transport/          # HTTP, WebSocket, auth transport
├── packages/
│   ├── api_contracts/          # OpenAPI/schema contracts
│   └── shared_config/          # Non-secret shared configuration
├── infrastructure/
│   ├── docker/                 # Local container definitions
│   └── environments/           # Environment-specific deployment config
├── docs/
│   └── architecture/
├── scripts/
└── .github/
    └── workflows/
```

Directory names are intentionally technology-aware only at application boundaries. Domain code should not be coupled to Flutter widgets, HTTP handlers, or database models.
