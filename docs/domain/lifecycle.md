# Domain Lifecycles

## MiningSession

`pending -> active -> completed`

Exceptional paths:

`pending -> cancelled`
`active -> paused`
`paused -> active`
`active -> cancelled`

Terminal states should not transition back to active.

## Transaction

`created -> pending -> processing -> completed`

Failure paths:

`pending -> rejected`
`processing -> failed`
`processing -> rejected`

A completed transaction is terminal.

## User

`active -> suspended -> active`
`active -> closed`

Closed accounts are terminal for normal product activity.

## Design note

Exact transition rules and authorization will be implemented in the application/domain layer, not in UI clients.
