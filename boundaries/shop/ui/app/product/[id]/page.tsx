import type { Metadata } from 'next';
import { notFound } from 'next/navigation';

import { GridTileImage } from 'components/grid/tile';
import Footer from 'components/layout/footer';
import { Gallery } from 'components/product/gallery';
import { ProductProvider } from 'components/product/product-context';
import { ProductDescription } from 'components/product/product-description';
import { HIDDEN_PRODUCT_TAG } from 'lib/constants';
import { getProduct, getProductRecommendations } from 'lib/shopify';
import { Image } from 'lib/shopify/types';
import Link from 'next/link';
import { Suspense } from 'react';

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#experimental_ppr
export const experimental_ppr = true

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#dynamic
export const dynamic = 'force-dynamic'

export async function generateMetadata({
  params
}: {
  params: { id: number };
}): Promise<Metadata> {
  console.warn('product 1', params);

  const product = await getProduct(params.id);

  if (!product) return notFound();

  console.warn('product 3', product);

  return {
    title: product.name,
    description: product.description,
  };
}

export default async function ProductPage({ params }: { params: { id: number } }) {
  const product = await getProduct(params.id);

  if (!product) return notFound();

  const productJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'Product',
    name: product.name,
    description: product.description,
    offers: {
      '@type': 'AggregateOffer',
      priceCurrency: product.price,
      highPrice: product.price,
      lowPrice: product.price,
    }
  };

  product.images = [{}, {}, {}, {}, {}];

  return (
    <ProductProvider>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{
          __html: JSON.stringify(productJsonLd)
        }}
      />
      <div className="mx-auto max-w-screen-2xl px-4">
        <div
          className='flex flex-col rounded-lg border border-neutral-200 bg-white p-8 md:p-12 lg:flex-row lg:gap-8 dark:border-neutral-800 dark:bg-black'>
          <div className='h-full w-full basis-full lg:basis-4/6'>
            <Suspense
              fallback={
                <div className='relative aspect-square h-full max-h-[550px] w-full overflow-hidden' />
              }
            >
              <Gallery
                images={product.images.slice(0, 5).map((image: Image) => ({
                  src: "https://picsum.photos/200",
                  altText: image.altText,
                }))}
              />
            </Suspense>
          </div>

          <div className='basis-full lg:basis-2/6'>
            <Suspense fallback={null}>
              <ProductDescription product={product} />
            </Suspense>
          </div>
        </div>
        {/*<RelatedProducts id={product.id} />*/}
      </div>
      <Footer />
    </ProductProvider>
  );
}

async function RelatedProducts({ id }: { id: number }) {
  const relatedProducts = await getProductRecommendations(id)

  if (!relatedProducts.length) return null

  return (
    <div className='py-8'>
      <h2 className='mb-4 text-2xl font-bold'>Related Products</h2>
      <ul className="flex w-full gap-4 overflow-x-auto pt-1">
        {relatedProducts.map((product) => (
          <li
            key={product.name}
            className="aspect-square w-full flex-none min-[475px]:w-1/2 sm:w-1/3 md:w-1/4 lg:w-1/5"
          >
            <Link
              className="relative h-full w-full"
              href={`/product/${product.name}`}
              prefetch={true}
            >
              <GridTileImage
                alt={product.name}
                label={{
                  title: product.name,
                  amount: product.price,
                  // currencyCode: product.priceRange.maxVariantPrice.currencyCode
                }}
                src={"https://picsum.photos/200"}
                fill
                sizes="(min-width: 1024px) 20vw, (min-width: 768px) 25vw, (min-width: 640px) 33vw, (min-width: 475px) 50vw, 100vw"
              />
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
