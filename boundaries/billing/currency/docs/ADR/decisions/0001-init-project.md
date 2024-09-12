# 1. Init project

Date: 2024-09-12

## Status

Accepted

## Context

To support the company's need for a multi-currency billing system, we are starting a new service to handle currency 
conversion. This service will provide real-time and historical exchange rate data.

## Decision

We will create a new service called `currency` to handle currency conversion.

## Consequences

- We will need to integrate with a third-party service to get exchange rate data.
- We will need to store historical exchange rate data.
- We will need to provide an API for converting between currencies.
- We will need to provide an API for getting exchange rate data.
- We will need to provide an API for getting historical exchange rate data.
