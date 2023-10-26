import { NextPage, Metadata, Viewport } from 'next'

import Home from '../components/Home/Home'

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

export const metadata: Metadata = {
  title: 'Welcome | Routing by project',
}

const HomePage: NextPage = () => <Home />

export default HomePage
