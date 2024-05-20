# ADR-0002: Use Decimal for Financial Types

Date: 2024-05-21

## Status

Accepted

## Context

Handling financial calculations requires precision and accuracy to avoid errors in monetary transactions. 
Using floating-point numbers can lead to precision issues due to their binary representation. 
Therefore, a decision has been made to use a library specifically designed for decimal arithmetic 
to ensure correctness in financial computations.

## Decision

We will use the [Shopspring Decimal](https://github.com/shopspring/decimal) library 
for handling financial types in our application. This library provides:

- **Arbitrary-precision**: Ensures that calculations are accurate and free from floating-point rounding errors.
- **Ease of use**: The API is simple and intuitive, making it easy to integrate into our existing codebase.
- **Performance**: Although slightly slower than native floating-point arithmetic, the trade-off is acceptable 
given the importance of precision in financial calculations.
- **Reliability**: It is a well-maintained library with a strong community and extensive documentation.

## Consequences

With the adoption of the Shopspring Decimal library, we ensure that our financial computations are precise and reliable. 

This decision will:

- **Improve accuracy**: Eliminate errors caused by floating-point arithmetic, 
ensuring that all financial calculations are accurate to the last cent.
- **Enhance maintainability**: Provide a consistent approach to handling financial types across the codebase, 
making it easier to maintain and understand.
- **Increase confidence**: Give stakeholders confidence that the financial aspects of 
the system are robust and trustworthy.
