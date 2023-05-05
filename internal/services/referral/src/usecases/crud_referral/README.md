## Use Case: UC-1 CRUD referral

### Description

CRUD operations for referral:

  * Create referral 
  * Read referral
  * Update referral
  * Delete referral
  * List referrals

### Actors

  * Manager
  * Referral

### Flow of events

```mermaid
graph LR
    style Manager fill:#7dcfe3
    style Referral fill:#9dcc7a

    Manager -- create --> Referral
    Manager -- read --> Referral
    Manager -- update --> Referral
    Manager -- delete --> Referral
    Manager -- list --> Referral
```
