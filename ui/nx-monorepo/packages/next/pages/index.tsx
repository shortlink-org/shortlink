import type { NextPage } from 'next'
import { SoftwareAppJsonLd, NextSeo } from 'next-seo'

import { Layout } from 'components'
import Feature from 'components/Landing/feature'
import Header from 'components/Landing/header'
import Mobile from 'components/Landing/mobile'
import Subscribe from 'components/Landing/subscribe'
import Testimonials from 'components/Testimonials'

const Home: NextPage = () => (
  <Layout>
    <NextSeo
      title="Landing Page Service"
      description="Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!"
      openGraph={{
        title: 'Landing',
        description:
          'Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!',
        type: 'article',
        article: {
          publishedTime: '2021-08-01T05:00:00.000Z',
          modifiedTime: '2021-08-01T05:00:00.000Z',
          section: 'Landing',
          authors: ['https://batazor.ru'],
          tags: ['shortlink', 'landing'],
        },
      }}
    />
    <SoftwareAppJsonLd
      name="Shortlink"
      price="Free"
      priceCurrency="USD"
      aggregateRating={{ ratingValue: '5', reviewCount: '8864' }}
      operatingSystem="Web"
      applicationCategory="Productivity"
    />

    <Header />
    <Mobile />
    <Feature />
    <Testimonials />
    <Subscribe />
  </Layout>
)

export default Home
