import { NextPage, Metadata } from 'next'

import Home from '../components/Home/Home'

export const metadata: Metadata = {
  title: 'Welcome | Routing by project',
}

const HomePage: NextPage = () => <Home />

export default HomePage
