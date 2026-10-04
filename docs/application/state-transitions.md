# Application State Transitions

Use cases are responsible for validating transitions before persistence.

## Mining session

Allowed:

- pending -> active
- pending -> cancelled
- active -> paused
- paused -> active
- active -> completed
- active -> cancelled

Terminal:
- completed
- cancelled

## Transaction

Allowed:

- created -> pending
- pending -> processing
- pending -> rejected
- processing -> completed
- processing -> failed
- processing -> rejected

Terminal:
- completed
- failed
- rejected

## Concurrency

Starting, pausing, resuming, or completing a session must protect against concurrent commands.

The persistence implementation will later use transactions plus row/version locking as appropriate. Redis locks are not the financial source of truth.
