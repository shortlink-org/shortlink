import { CartProvider } from 'components/cart/cart-context';
import { Navbar } from 'components/layout/navbar';
import { GeistSans } from 'geist/font/sans';
import { getCart } from 'lib/shopify';
import { ensureStartsWith } from 'lib/utils';
import { cookies } from 'next/headers';
import { ReactNode } from 'react';
import { Toaster } from 'sonner';
import './globals.css';

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#experimental_ppr
export const experimental_ppr = true

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#dynamic
export const dynamic = 'force-dynamic'

const { TWITTER_CREATOR, TWITTER_SITE, SITE_NAME } = process.env;
const baseUrl = process.env.NEXT_PUBLIC_VERCEL_URL
  ? `https://${process.env.NEXT_PUBLIC_VERCEL_URL}`
  : 'http://localhost:3000';
const twitterCreator = TWITTER_CREATOR ? ensureStartsWith(TWITTER_CREATOR, '@') : undefined;
const twitterSite = TWITTER_SITE ? ensureStartsWith(TWITTER_SITE, 'https://') : undefined;

export const metadata = {
  metadataBase: new URL(baseUrl),
  title: {
    default: SITE_NAME!,
    template: `%s | ${SITE_NAME}`
  },
  robots: {
    follow: true,
    index: true
  },
  ...(twitterCreator &&
    twitterSite && {
      twitter: {
        card: 'summary_large_image',
        creator: twitterCreator,
        site: twitterSite
      }
    })
};

export default async function RootLayout({ children }: { children: ReactNode }) {
  const cartId = cookies().get('cartId')?.value;
  // Don't await the fetch, pass the Promise to the context provider
  // const cart = getCart(cartId);

  const mockCart = {
    id: 'mockCartId',
    checkoutUrl: 'http://localhost:3000/checkout',
    totalQuantity: 1,
    lines: [
      {
        id: 'mockLineId',
        quantity: 1,
        cost: {
          totalAmount: {
            amount: '100.00',
            currencyCode: 'USD'
          }
        },
        merchandise: {
          id: 'mockMerchandiseId',
          title: 'Mock Product',
          selectedOptions: [],
          product: {
            id: 'mockProductId',
            handle: 'mock-product',
            title: 'Mock Product',
            featuredImage: null
          }
        }
      }
    ],
    cost: {
      subtotalAmount: { amount: '100.00', currencyCode: 'USD' },
      totalAmount: { amount: '100.00', currencyCode: 'USD' },
      totalTaxAmount: { amount: '0', currencyCode: 'USD' }
    }
  };

  return (
    <html lang="en" className={GeistSans.variable}>
      <body className="bg-neutral-50 text-black selection:bg-teal-300 dark:bg-neutral-900 dark:text-white dark:selection:bg-pink-500 dark:selection:text-white">
        {/*<CartProvider cartPromise={Promise.resolve(mockCart)}>*/}
          <Navbar />
          {/*<main>*/}
          {/*  {children}*/}
          {/*  <Toaster closeButton />*/}
          {/*</main>*/}
        {/*</CartProvider>*/}
      </body>
    </html>
  );
}
