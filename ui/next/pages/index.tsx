import React from 'react';
import Feature from 'components/Landing/feature';
import Header from 'components/Landing/header';
import Mobile from 'components/Landing/mobile';
import Subscribe from 'components/Landing/subscribe';
import Testimonials from 'components/Testimonials';

import { Layout } from 'components';

export function ProfileContent() {
  return (
    <>
      <Header />
      <Mobile />
      <Feature />
      <Testimonials />
      <Subscribe />
    </>
  );
}

export default function Profile() {
  return <Layout content={ProfileContent()} />;
}
