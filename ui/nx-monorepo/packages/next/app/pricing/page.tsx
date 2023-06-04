'use client'

import React, { Fragment } from 'react'

import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import CardHeader from '@mui/material/CardHeader'
import Grid from '@mui/material/Grid'
import StarIcon from '@mui/icons-material/StarBorder'
import Typography from '@mui/material/Typography'
// @ts-ignore
import Discounted from 'components/Billing/Discounted'

const tiers = [
  {
    title: 'Free',
    subheader: 'Best option for personal use & for your next project.',
    price: '0',
    description: [
      '10 users included',
      '2 GB of storage',
      'Help center access',
      'Email support',
    ],
    buttonText: 'Sign up for free',
    buttonVariant: 'outlined',
  },
  {
    title: 'Pro',
    subheader: 'Most popular choice for small teams.',
    price: '15',
    description: [
      '20 users included',
      '10 GB of storage',
      'Help center access',
      'Priority email support',
    ],
    buttonText: 'Get started',
    buttonVariant: 'outlined',
  },
  {
    title: 'Enterprise',
    subheader: 'Best for large scale uses and extended redistribution rights.',
    price: '30',
    description: [
      '50 users included',
      '30 GB of storage',
      'Help center access',
      'Phone & email support',
    ],
    buttonText: 'Contact us',
    buttonVariant: 'outlined',
  },
]

function Pricing() {
  // <NextSeo
  //   title="Pricing Policy"
  //   description="Shortlink offers fair, upfront pricing for all of our services. We never charge hidden fees or use bait-and-switch tactics. You'll know exactly what you're getting before your start working on your project."
  //   openGraph={{
  //     title: 'Pricing Policy',
  //     description:
  //       "Shortlink offers fair, upfront pricing for all of our services. We never charge hidden fees or use bait-and-switch tactics. You'll know exactly what you're getting before your start working on your project.",
  //     type: 'article',
  //     article: {
  //       publishedTime: '2021-08-01T05:00:00.000Z',
  //       modifiedTime: '2021-08-01T05:00:00.000Z',
  //       section: 'Pricing',
  //       authors: ['https://batazor.ru'],
  //       tags: ['shortlink', 'pricing'],
  //     },
  //   }}
  // />
  //
  // <ArticleJsonLd
  //   url="https://shortlink.best/next/about"
  //   title="Pricing Policy"
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

  return (
    <Fragment>
      <div className="container mx-auto px-5 py-10 sm:max-w-sm">
        <Typography
          component="h1"
          variant="h2"
          align="center"
          color="text.primary"
          gutterBottom
        >
          Pricing
        </Typography>
        <Typography
          variant="h5"
          align="center"
          color="text.secondary"
          component="p"
        >
          Quickly build an effective pricing table for your potential customers
          with this layout. It&apos;s built with default Material-UI components
          with little customization.
        </Typography>

        <div className="flex flex-col justify-center text-xs text-gray-600 md:flex-row">
          <div className="flex items-center p-4">
            <svg
              viewBox="0 0 20 20"
              fill="currentColor"
              className="w-4 h-4 mr-1 text-green-600"
            >
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            No credit card required
          </div>

          <div className="flex items-center p-4">
            <svg
              viewBox="0 0 20 20"
              fill="currentColor"
              className="w-4 h-4 mr-1 text-green-600"
            >
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            14 days free
          </div>

          <div className="flex items-center p-4">
            <svg
              viewBox="0 0 20 20"
              fill="currentColor"
              className="w-4 h-4 mr-1 text-green-600"
            >
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
            Cancel anytime
          </div>
        </div>
      </div>

      <Discounted />

      <Grid container spacing={5} alignItems="flex-end">
        {tiers.map((tier) => (
          // Enterprise card is full width at sm breakpoint
          <Grid
            item
            key={tier.title}
            xs={12}
            sm={tier.title === 'Enterprise' ? 12 : 6}
            md={4}
          >
            <Card className="rounded-lg border border-gray-100 shadow dark:border-gray-600 xl:p-8 dark:bg-gray-800 dark:text-white">
              <CardHeader
                title={tier.title}
                subheader={tier.subheader}
                titleTypographyProps={{ align: 'center' }}
                action={tier.title === 'Pro' ? <StarIcon /> : null}
                subheaderTypographyProps={{
                  align: 'center',
                }}
                classes={{
                  title: 'mb-4 text-2xl font-semibold',
                  subheader: 'font-light sm:text-lg dark:text-white',
                }}
                sx={{
                  backgroundColor: (theme) =>
                    theme.palette.mode === 'light'
                      ? theme.palette.grey[200]
                      : theme.palette.grey[700],
                }}
              />
              <CardContent>
                <div className="flex justify-center items-baseline my-8">
                  <span className="mr-2 text-5xl font-extrabold">
                    ${tier.price}
                  </span>
                  <span className="text-gray-500 dark:text-gray-400">/mo</span>
                </div>

                <ul data-role="list" className="mb-8 space-y-4 text-left">
                  {tier.description.map((line) => (
                    <Typography
                      component="li"
                      variant="subtitle1"
                      align="center"
                      className="flex items-center space-x-3"
                      key={line}
                    >
                      <svg
                        className="flex-shrink-0 w-5 h-5 text-green-500 dark:text-green-400"
                        fill="currentColor"
                        viewBox="0 0 20 20"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          fillRule="evenodd"
                          d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                          clipRule="evenodd"
                        />
                      </svg>

                      <span>{line}</span>
                    </Typography>
                  ))}
                </ul>
              </CardContent>
              <CardActions>
                <Button
                  fullWidth
                  variant={tier.buttonVariant as 'outlined' | 'contained'}
                >
                  {tier.buttonText}
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Fragment>
  )
}

export default Pricing
