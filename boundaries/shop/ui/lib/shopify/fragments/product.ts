import imageFragment from './image';
import seoFragment from './seo';

const productFragment = /* GraphQL */ `
  fragment product on Product {
    id
    name
    price
    description
    created_at
    updated_at
  }
  // ${imageFragment}
  // ${seoFragment}
`;

export default productFragment;
