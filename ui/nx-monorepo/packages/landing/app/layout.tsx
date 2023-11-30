import { Metadata } from 'next'
// eslint-disable-next-line camelcase
import { Roboto_Mono } from 'next/font/google'
import Script from 'next/script'
import * as React from 'react'
import { Organization, WithContext } from 'schema-dts'

import { Providers } from './providers'
import 'public/assets/styles.css'

const robotoMono = Roboto_Mono({
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-inter',
})

export async function generateMetadata(): Promise<Metadata> {
  return {
    // https://nextjs.org/docs/app/api-reference/functions/generate-metadata#metadatabase
    metadataBase: new URL('https://shortlink.best/'),
    alternates: {
      canonical: '/',
      languages: {
        // example:
        // en: '/',
        // ru: '/ru',
      },
    },
    title: {
      template: '%s | Routing by project',
      default: 'ShortLink | Routing by project',
    },
    description:
      "Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you.",
    openGraph: {
      type: 'website',
      title: 'ShortLink | Routing by project',
      description:
        "Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you.",
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
  return (
    <html
      lang="en"
      className={`${robotoMono.className} font-sans`}
      suppressHydrationWarning
    >
      <Script
        id="json-ld"
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />

      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
