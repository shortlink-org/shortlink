import React from 'react'
import { Metadata } from 'next'
import { Organization, WithContext } from 'schema-dts'
import Script from 'next/script'
// eslint-disable-next-line camelcase
import { Roboto_Mono } from 'next/font/google'

import { Providers } from './providers'
import 'public/assets/styles.css'
import 'react-toastify/dist/ReactToastify.css'

const robotoMono = Roboto_Mono({
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-inter',
})

export async function generateMetadata(): Promise<Metadata> {
  return {
    title: 'Landing Page Service | Shortlink',
    description:
      'Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!',
    openGraph: {
      type: 'website',
      title: 'Landing Page Service | Shortlink',
      description: 'Shortlink is your go-to source for all things URL. We offer a wide range of services, including shortening, tracking, and protecting links. Visit our website today to learn more!',
      locale: 'en_IE',
      url: 'https://shortlink.best/',
      siteName: 'ShortLink',
      images: [
        {
          url: 'https://shortlink.best/images/logo.png',
          width: 600,
          height: 600,
          alt: 'ShortLink service',
        },
      ],
    },
    twitter: {
      site: '@shortlink',
      title: 'ShortLink',
      description: 'ShortLink service',
      images: 'https://shortlink.best/images/logo.png',
    },
    viewport: 'initial-scale=1, width=device-width',
    manifest: '/manifest.json',
    icons: ['/favicon.ico'],
  }
}

const jsonLd: WithContext<Organization> = {
  '@context': 'https://schema.org',
  '@type': 'Organization',
  name: 'ShortLink',
  url: 'https://shortlink.best/',
  logo: 'https://shortlink.best/images/logo.png',
  description: 'ShortLink service',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const GA_ID = process.env.NEXT_PUBLIC_GOOGLE_ANALYTICS_ID

  return (
    <html lang="en" className={`${robotoMono.className} font-sans`}>
      <Script
        id="json-ld"
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <Script
        id="google-analytics-script"
        src={`https://www.googletagmanager.com/gtag/js?id=${GA_ID}`}
        strategy="afterInteractive"
      />
      <Script id="google-analytics" strategy="afterInteractive">
        {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', '${GA_ID}');
        `}
      </Script>

      <body className="bg-white dark:bg-gray-800 text-black dark:text-white">
        {/* @ts-ignore */}
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
