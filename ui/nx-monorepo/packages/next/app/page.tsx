'use client'

import React from 'react'
import type { NextPage } from 'next'
import Header from 'components/Landing/header'
import Mobile from 'components/Landing/mobile'
import Feature from 'components/Landing/feature'
import Testimonials from 'components/Testimonials'
import Subscribe from 'components/Landing/subscribe'

import { Layout } from 'components'

const Home: NextPage = () => (
  <Layout>
    <Header />
    <Mobile />
    <Feature />
    <Testimonials />
    <Subscribe />
  </Layout>
)

export default Home
