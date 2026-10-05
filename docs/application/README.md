# Application Layer

The application layer orchestrates use cases and coordinates domain rules with persistence ports.

## Rules

- Use cases are the entry points for business operations.
- Use cases depend on domain interfaces, not PostgreSQL, Redis, HTTP, Flutter, or Telegram.
- Authorization is evaluated at the application boundary.
- A use case must be safe to retry when its operation can be triggered more than once.
- Financial mutations are committed atomically with their required ledger operation.
- Client-provided timestamps, reward quantities, and balances are never authoritative.

## Phase 0 use-case map

### Mining lifecycle
- StartMiningSession
- PauseMiningSession
- ResumeMiningSession
- CompleteMiningSession
- CancelMiningSession

### Reward
- CalculateReward
- PostReward

### Wallet / finance
- GetWalletBalance
- CreateDeposit
- RequestWithdrawal

### Configuration
- CreateAsset
- UpdateAsset
- CreateMiningPlan
- PublishMiningPlanVersion

Only the contracts and orchestration boundaries are established in this step. Production calculations, blockchain settlement, and payment integrations remain later phases.
