import React from 'react'
import type { NextPage } from 'next'
import { SoftwareAppJsonLd, NextSeo } from "next-seo";
import Header from 'components/Landing/header'
import Mobile from 'components/Landing/mobile'
import Feature from 'components/Landing/feature'
import Testimonials from 'components/Testimonials'
import Subscribe from 'components/Landing/subscribe'

import { Layout } from 'components'

const Home: NextPage = () => (
  <Layout>
    <NextSeo
      title="Landing"
      description="Landing page for shortlink."
      openGraph={{
        title: "Landing",
        description: "Landing page for shortlink.",
        type: "article",
        article: {
          publishedTime: "2021-08-01T05:00:00.000Z",
          modifiedTime: "2021-08-01T05:00:00.000Z",
          section: "Landing",
          authors: [
            "https://batazor.ru",
          ],
          tags: [ "shortlink", "landing" ],
        }
      }}
    />
    <SoftwareAppJsonLd
      name={"Shortlink"}
      price={"Free"}
      priceCurrency={"USD"}
      aggregateRating={{ ratingValue: '5', reviewCount: '8864' }}
      operatingSystem={"Web"}
      applicationCategory={"Productivity"}
    />

    <Header />
    <Mobile />
    <Feature />
    <Testimonials />
    <Subscribe />
  </Layout>
)

export default Home
