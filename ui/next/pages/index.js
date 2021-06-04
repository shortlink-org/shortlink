import React from 'react'
import Feature from '../components/Landing/feature'
import Header from '../components/Landing/header'
import Testimonials from '../components/Testimonials'

import { Layout } from '../components';

export function ProfileContent() {
  return (
    <React.Fragment>
      <Header />
      <Feature />
      <br />
      <br />
      <br />
      <br />
      <Testimonials />
    </React.Fragment>
  )
}

export default function Profile() {
  return <Layout content={ProfileContent()} />;
}
