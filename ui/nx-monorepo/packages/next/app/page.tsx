'use client'

import React, { Fragment } from 'react'
import type { NextPage } from 'next'
import Header from 'components/Landing/header'
import Mobile from 'components/Landing/mobile'
import Feature from 'components/Landing/feature'
import Testimonials from 'components/Testimonials'
import Subscribe from 'components/Landing/subscribe'

const Home: NextPage = () => (
  <Fragment>
    <Header />
    <Mobile />
    <Feature />
    <Testimonials />
    <Subscribe />
  </Fragment>
)

export default Home
