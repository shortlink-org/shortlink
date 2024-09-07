import cartFragment from '../fragments/cart';

export const getCartQuery = /* GraphQL */ `
  query Carts_getCart(
    $customerId: String!
  ) {
    carts_getCart(
      customerId: { customerId: $customerId }
    ) {
      state {
        cartId
        customerId
        items {
          productId
          quantity
        }
      }
    }
  }
`;
