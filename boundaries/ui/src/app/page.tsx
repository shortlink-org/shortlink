'use client'

import { NextPage } from 'next'

import Feature from '@/components/Landing/feature'
import Header from '@/components/Landing/header'
import { LandingOpenSource } from '@/components/Landing/landing-open-source'
import { LandingStats } from '@/components/Landing/landing-stats'
import Mobile from '@/components/Landing/mobile'
import Subscribe from '@/components/Landing/subscribe'
import Testimonials from '@/components/Testimonials/Testimonials'

import styles from './page.module.css'

// <NextSeo
// title="Landing Page Service"
// description="Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!"
// openGraph={{
//   title: 'Landing',
//     description:
//   'Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!',
//     type: 'article',
//     article: {
//     publishedTime: '2021-08-01T05:00:00.000Z',
//       modifiedTime: '2021-08-01T05:00:00.000Z',
//       section: 'Landing',
//       authors: ['https://batazor.ru'],
//       tags: ['shortlink', 'landing'],
//   },
// }}
// />
// <SoftwareAppJsonLd
//   name="Shortlink"
//   price="Free"
//   priceCurrency="USD"
//   aggregateRating={{ ratingValue: '5', reviewCount: '8864' }}
//   operatingSystem="Web"
//   applicationCategory="Productivity"
// />

// @ts-ignore
const Home: NextPage = () => (
  <div className="mx-auto w-full min-w-0 max-w-7xl px-4 pt-5 sm:px-6 sm:pt-7 lg:px-8 lg:pt-8 xl:px-10 2xl:px-12">
    <Header />

    <div className="mt-6 flex w-full min-w-0 flex-col rounded-[2rem] border border-[var(--color-border)]/80 bg-[var(--color-surface)]/85 py-12 shadow-[0_24px_80px_-48px_rgb(15_23_42/0.35)] backdrop-blur-md dark:border-white/[0.08] dark:bg-[var(--color-surface)]/75 dark:shadow-[0_28px_90px_-52px_rgb(0_0_0/0.75)] sm:mt-8 sm:py-14 lg:mt-10 lg:py-16 xl:py-20">
      <div
        className={`${styles.mainStack} flex flex-col px-5 sm:px-7 lg:px-10 xl:px-12`}
      >
        <LandingStats />
        <Mobile />
        <Feature />
        <Testimonials />
        <LandingOpenSource />
        <Subscribe />
      </div>
    </div>
  </div>
)

// @ts-ignore
export default Home
