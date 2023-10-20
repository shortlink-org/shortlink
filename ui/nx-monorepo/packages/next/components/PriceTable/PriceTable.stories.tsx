// @ts-nocheck
import { BeakerIcon } from '@heroicons/react/20/solid'
import { Meta, StoryFn } from '@storybook/react'

import PriceTable, { TiersProps } from './PriceTable'

const meta: Meta<TiersProps> = {
  title: 'Page/PriceTable',
  component: PriceTable,
}

export default meta

export const Default = {
  args: {
    tiers: [
      {
        title: 'Free',
        subheader: 'Best option for personal use & for your next project.',
        price: 0,
        description: [
          'Consectetur adipiscing elit',
          '10 users included',
          '2 GB of storage',
          'Help center access',
          'Email support',
        ],
        buttonText: 'Sign up for free',
        buttonVariant: 'outlined',
      },
      {
        title: 'Pro',
        subheader: 'Most popular choice for small teams.',
        price: 15,
        description: [
          'Consectetur adipiscing elit',
          '20 users included',
          '10 GB of storage',
          'Help center access',
          'Priority email support',
        ],
        buttonText: 'Get started',
        buttonVariant: 'outlined',
      },
      {
        title: 'Enterprise',
        subheader:
          'Best for large scale uses and extended redistribution rights.',
        price: 30,
        description: [
          'Consectetur adipiscing elit',
          '50 users included',
          '30 GB of storage',
          'Help center access',
          'Phone & email support',
        ],
        buttonText: 'Contact us',
        buttonVariant: 'outlined',
      },
    ],
  },
}
