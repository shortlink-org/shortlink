'use client'

import Container from '@mui/material/Container'
import Typography from '@mui/material/Typography'
// @ts-ignore
import { PriceTable } from '@shortlink-org/ui-kit'

import Discounted from 'components/Billing/Discounted'

const tiers = [
  {
    title: 'Free',
    subheader: 'Best option for personal use & for your next project.',
    price: 0,
    description: ['Consectetur adipiscing elit', '10 users included', '2 GB of storage', 'Help center access', 'Email support'],
    buttonText: 'Sign up for free',
    buttonVariant: 'outlined',
  },
  {
    title: 'Pro',
    subheader: 'Most popular choice for small teams.',
    price: 15,
    description: ['Consectetur adipiscing elit', '20 users included', '10 GB of storage', 'Help center access', 'Priority email support'],
    buttonText: 'Get started',
    buttonVariant: 'outlined',
  },
  {
    title: 'Enterprise',
    subheader: 'Best for large scale uses and extended redistribution rights.',
    price: 30,
    description: ['Consectetur adipiscing elit', '50 users included', '30 GB of storage', 'Help center access', 'Phone & email support'],
    buttonText: 'Contact us',
    buttonVariant: 'outlined',
  },
]


// <NextSeo
// title="Page Policy"
// description="Shortlink offers fair, upfront pricing for all of our services. We never charge hidden fees or use bait-and-switch tactics. You'll know exactly what you're getting before your start working on your project."
// openGraph={{
//   title: 'Pricing Policy',
//     description:
//   "Shortlink offers fair, upfront pricing for all of our services. We never charge hidden fees or use bait-and-switch tactics. You'll know exactly what you're getting before your start working on your project.",
//     type: 'article',
//     article: {
//     publishedTime: '2021-08-01T05:00:00.000Z',
//       modifiedTime: '2021-08-01T05:00:00.000Z',
//       section: 'Pricing',
//       authors: ['https://batazor.ru'],
//       tags: ['shortlink', 'pricing'],
//   },
// }}
// />
//
// <ArticleJsonLd
//   url="https://shortlink.best/next/about"
//   title="Page Policy"
//   images={['https://shortlink.best/images/logo.png']}
//   datePublished="2021-08-01T05:00:00.000Z"
//   dateModified="2021-08-01T05:00:00.000Z"
//   authorName={[
//     {
//       name: 'Login Viktor',
//       url: 'https://batazor.ru',
//     },
//   ]}
//   publisherName="Login Viktor"
//   publisherLogo="https://shortlink.best/images/logo.png"
//   description="Shortlink offers fair, upfront pricing for all of our services. We never charge hidden fees or use bait-and-switch tactics. You'll know exactly what you're getting before your start working on your project."
// />

function Page() {
  return (
    <>
      <Container disableGutters maxWidth="sm" component="main" sx={{ pt: 8, pb: 6 }}>
        <Typography component="h1" variant="h2" align="center" color="text.primary" gutterBottom>
          Page
        </Typography>
        <Typography variant="h5" align="center" color="text.secondary" component="p">
          Quickly build an effective pricing table for your potential customers with this layout. It&apos;s built with default Material-UI
          components with little customization.
        </Typography>

        <div className="flex flex-col justify-center text-xs text-gray-600 md:flex-row">
          <div className="flex items-center p-4">
            <svg viewBox="0 0 20 20" fill="currentColor" className="w-4 h-4 mr-1 text-green-600">
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            No credit card required
          </div>

          <div className="flex items-center p-4">
            <svg viewBox="0 0 20 20" fill="currentColor" className="w-4 h-4 mr-1 text-green-600">
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            14 days free
          </div>

          <div className="flex items-center p-4">
            <svg viewBox="0 0 20 20" fill="currentColor" className="w-4 h-4 mr-1 text-green-600">
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            Cancel anytime
          </div>
        </div>
      </Container>

      <Discounted />

      <PriceTable tiers={tiers} />
    </>
  )
}

// @ts-ignore
export default Page
