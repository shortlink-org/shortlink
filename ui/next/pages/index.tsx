import React from 'react'
import type { NextPage } from 'next'
import { NextSeo } from 'next-seo'
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
    />
    <Header />
    <Mobile />
    <Feature />
    <Testimonials />
    <Subscribe />
  </Layout>
)

export default Home
