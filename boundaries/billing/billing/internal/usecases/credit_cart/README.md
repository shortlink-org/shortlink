## Use Case: Credit Card

> [!NOTE]
>
> This use case handles billing operations related to credit cards,
> including CRUD operations and validation through the Luhn algorithm.

### Use Cases

#### 1. Create Credit Card

- Add a new credit card entry with card details such as number, expiration date, etc.
- Applies Luhn validation to ensure the card number is valid.

#### 2. Read Credit Card

- Retrieves details of a specific credit card based on its unique identifier.

#### 3. Update Credit Card

- Modifies existing credit card’s information (e.g., expiration date).

#### 4. Delete Credit Card

- Removes a credit card from the system.

#### 5. Validate Credit Card

- Run a check using the Luhn algorithm to verify that a credit card number is valid before processing any transactions. [(Luhn algorithm on Wikipedia)](https://en.wikipedia.org/wiki/Luhn_algorithm)
