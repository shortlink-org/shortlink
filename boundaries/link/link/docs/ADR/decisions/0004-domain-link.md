# 4. Domain: Link

Date: 2024-03-17

## Status

Accepted

## Context

Link is a core domain of the system. It is used to store and manage links.

## Decision

### URL Length Policy

**Decision:**  
We have decided **not to impose an explicit limit on the length of the original (long) URL**. Instead, the length constraint will be governed by the capabilities and limitations of the underlying database system.

**Rationale:**

- **Database Constraints:** Modern databases can handle lengthy text fields efficiently. Relying on the database to enforce length limits ensures consistency and leverages existing infrastructure capabilities.
- **User Flexibility:** Allowing unrestricted URL lengths ensures that users can shorten any valid URL without encountering unnecessary restrictions, enhancing user satisfaction.
- **Future-Proofing:** As URL standards evolve or as users generate longer URLs, the system remains adaptable without requiring architectural changes.

### Short URL Generation Strategy

- **Optimal Hash Length:** `6` characters
- **Character Set:** `62` symbols (a-z, A-Z, 0-9)
- **Total Unique URLs:** `62^6 ≈ 56,800,235,584`

**Rationale:**

- **Collision Probability:** With approximately 56.8 billion unique combinations, the probability of collision is extremely low, especially when distributing URL creations evenly.
- **Daily URL Creation Estimate:**
  - **Total Daily Creations:** `5,000,000`
  - **Unique URLs Needed Per Day:** `5,000,000 / 7 ≈ 714,285` (assuming a 7-day retention or processing window)
  - **Adequacy:** The chosen hash length provides ample unique combinations well beyond the daily requirements, minimizing the risk of collisions.

**Notes:**

> **Daily URL Creations:**  
> `5,000,000 / 7` ≈ `714,285`

> **Optimal Hash Length:** `6`  
> **Count of Symbols:** `62` (a-z, A-Z, 0–9)  
> **Count of Unique URLs:** `62^6` ≈ `56,800,235,584`

## Consequences

This decision allows us to create a large number of unique URLs and not limit the length of the URL. 
This will help us to avoid collisions and provide a better user experience.
