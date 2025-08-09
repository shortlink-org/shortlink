import { Metadata, Viewport } from 'next'
import { Roboto_Mono } from 'next/font/google'
import Script from 'next/script'
import { Organization, WithContext } from 'schema-dts'

import Providers from './providers'
import '../public/assets/styles.css'

const robotoMono = Roboto_Mono({
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-inter',
})

export const viewport: Viewport = {
  width: 'device-width',
  initialScale: 1,
  maximumScale: 1,
  themeColor: [
    { media: '(prefers-color-scheme: light)', color: 'cyan' },
    { media: '(prefers-color-scheme: dark)', color: 'black' },
  ],
  colorScheme: 'light',
}

export async function generateMetadata(): Promise<Metadata> {
  return {
    // https://nextjs.org/docs/app/api-reference/functions/generate-metadata#metadatabase
    metadataBase: new URL('https://shortlink.best/next'),
    alternates: {
      canonical: '/',
      languages: {
        en: '/',
        // ru: '/ru',
      },
    },
    title: {
      template: '%s | ShortLink',
      default: 'ShortLink',
    },
    description:
      "Shortlink is the simplest way to manage your projects. It's an online platform that lets you create, share, and track links for you.",
    openGraph: {
      type: 'website',
      title: 'ShortLink',
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
    manifest: '/next/manifest.json',
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

// TODO: research problem with faro
// initializeFaro({
//   url: process.env.NEXT_PUBLIC_FARO_URI,
//   app: {
//     name: process.env.NEXT_PUBLIC_SERVICE_NAME,
//     version: process.env.NEXT_PUBLIC_GIT_TAG,
//     environment: 'production',
//   },
//   instrumentations: [
//     // Mandatory, overwriting the instrumentations array would cause the default instrumentations to be omitted
//     ...getWebInstrumentations(),
//
//     // Initialization of the tracing package.
//     // This package is optional because it increases the bundle size noticeably. Only add it if you want tracing data.
//     new TracingInstrumentation(),
//   ],
// })

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" className={`${robotoMono.className} font-sans`} suppressHydrationWarning>
      <Script id="json-ld" type="application/ld+json" dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }} />

      <body className="bg-white text-black dark:bg-black dark:text-white h-screen w-screen">
        <Providers>{children}</Providers>
      </body>
    </html>
  )
}
