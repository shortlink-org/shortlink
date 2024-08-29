export const getProductQuery = /* GraphQL */ `
  query Goods_goods_retrieve($id: Int!) {
    goods_goods_retrieve(id: $id) {
        created_at
        description
        id
        name
        price
        updated_at
    }
  }
`;

export const getProductsQuery = /* GraphQL */ `
  query Goods_goods_retrieve($id: Int!) {
    goods_goods_retrieve(id: $id) {
        created_at
        description
        id
        name
        price
        updated_at
    }
  }
`;

export const getProductRecommendationsQuery = /* GraphQL */ `
  query Goods_goods_retrieve($id: Int!) {
    goods_goods_retrieve(id: $id) {
        created_at
        description
        id
        name
        price
        updated_at
    }
  }
`;
