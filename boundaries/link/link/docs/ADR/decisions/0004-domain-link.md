# 4. Domain: Link

Date: 2024-03-17

## Status

Accepted

## Context

Link is a core domain of the system. It is used to store and manage links.

## Decision

### What should be the length of the URL?

We've decided don't limit the length of the URL. The URL length is limited by the database.

> [!NOTE]
>
> **Daily URL creations:** `5,000,000 / 7` ≈ `714,285`

> [!TIP]
>
> **Optimal Hash length**: `6`
> **Count of symbols**: `62` (a-z, A-Z, 0–9)
> **Count of unique URLs**: `62^6` ≈ `56,800,235,584`

## Consequences

This decision allows us to create a large number of unique URLs and not limit the length of the URL. 
This will help us to avoid collisions and provide a better user experience.
