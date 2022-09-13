import React from 'react'

import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import Box from '@mui/material/Box'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import CardHeader from '@mui/material/CardHeader'
import Grid from '@mui/material/Grid'
import StarIcon from '@mui/icons-material/StarBorder'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import { Layout } from 'components'
import Discounted from 'components/Billing/Discounted'
import { ArticleJsonLd, NextSeo } from "next-seo";

const tiers = [
  {
    title: 'Free',
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
    subheader: 'Most popular',
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

export function Pricing() {
  return (
    <Layout>
      <NextSeo
        title="Pricing"
        description="Pricing page for shortlink."
        openGraph={{
          title: "Pricing",
          description: "Pricing page for shortlink.",
          type: "article",
          article: {
            publishedTime: "2021-08-01T05:00:00.000Z",
            modifiedTime: "2021-08-01T05:00:00.000Z",
            section: "Pricing",
            authors: [
              "https://batazor.ru",
            ],
            tags: [ "shortlink", "pricing" ],
          }
        }}
      />
      <ArticleJsonLd
        url="https://architecture.ddns.net/next/about"
        title="Pricing"
        images={[
          'https://architecture.ddns.net/images/logo.png',
        ]}
        datePublished="2021-08-01T05:00:00.000Z"
        dateModified="2021-08-01T05:00:00.000Z"
        authorName={[
          {
            name: 'Login Viktor',
            url: 'https://batazor.ru',
          },
        ]}
        publisherName="Login Viktor"
        publisherLogo="https://architecture.ddns.net/images/logo.png"
        description="Pricing page for shortlink."
      />
      <Container
        disableGutters
        maxWidth="sm"
        component="main"
        sx={{ pt: 8, pb: 6 }}
      >
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
      </Container>

      <Discounted />

      <Container maxWidth="md" component="main">
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
              <Card>
                <CardHeader
                  title={tier.title}
                  subheader={tier.subheader}
                  titleTypographyProps={{ align: 'center' }}
                  action={tier.title === 'Pro' ? <StarIcon /> : null}
                  subheaderTypographyProps={{
                    align: 'center',
                  }}
                  sx={{
                    backgroundColor: (theme) =>
                      theme.palette.mode === 'light'
                        ? theme.palette.grey[200]
                        : theme.palette.grey[700],
                  }}
                />
                <CardContent>
                  <Box
                    sx={{
                      display: 'flex',
                      justifyContent: 'center',
                      alignItems: 'baseline',
                      mb: 2,
                    }}
                  >
                    <Typography
                      component="h2"
                      variant="h3"
                      color="text.primary"
                    >
                      ${tier.price}
                    </Typography>
                    <Typography variant="h6" color="text.secondary">
                      /mo
                    </Typography>
                  </Box>
                  <ul>
                    {tier.description.map((line) => (
                      <Typography
                        component="li"
                        variant="subtitle1"
                        align="center"
                        key={line}
                      >
                        {line}
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
      </Container>
    </Layout>
  )
}

export default Pricing
