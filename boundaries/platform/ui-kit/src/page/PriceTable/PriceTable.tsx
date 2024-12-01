import { OverridableStringUnion } from '@mui/types'
import Button, { ButtonPropsVariantOverrides } from '@mui/material/Button'
import Typography from '@mui/material/Typography'
import React, { useState } from 'react'
import Balancer from 'react-wrap-balancer'

import PricingToggle from './PricingToggle'

type Tier = {
  title: string
  subheader: string
  price: number
  description: string[]
  buttonVariant: OverridableStringUnion<
    'text' | 'outlined' | 'contained',
    ButtonPropsVariantOverrides
  >
  buttonText: string
}

export type TiersProps = {
  tiers: Tier[]
}

export const PriceTable: React.FC<TiersProps> = ({ tiers }) => {
  const [isAnnual, setIsAnnual] = useState<boolean>(true)

  return (
    <div className="w-full max-w-6xl mx-auto px-4 md:px-6 py-24">
      <PricingToggle isAnnual={isAnnual} setIsAnnual={setIsAnnual} />

      <div className="max-w-sm mx-auto grid gap-6 lg:grid-cols-3 items-start lg:max-w-none">
        {tiers.map((tier) => (
          // Enterprise card is full width at sm breakpoint
          <div
            className={`relative flex flex-col h-full p-6 rounded-2xl bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-900 shadow shadow-slate-950/5 ${
              tier.title === 'Pro' && 'bg-green-50 dark:bg-green-900'
            }`}
            key={tier.title}
          >
            {tier.title === 'Pro' && (
              <div className="absolute top-0 right-0 mr-6 -mt-4">
                <div className="inline-flex items-center text-xs font-semibold py-1.5 px-3 bg-emerald-500 text-white rounded-full shadow-sm shadow-slate-950/5">
                  Most Popular
                </div>
              </div>
            )}
            <div className="mb-5">
              <div className="text-slate-900 dark:text-slate-200 font-semibold mb-1">
                {tier.title}
              </div>
              <div className="inline-flex items-baseline mb-2">
                <span className="text-slate-900 dark:text-slate-200 font-bold text-3xl">
                  $
                </span>
                <span className="text-slate-900 dark:text-slate-200 font-bold text-4xl">
                  {isAnnual ? tier.price : tier.price * 1.2}
                </span>
                <span className="text-slate-500 font-medium">/mo</span>
              </div>
              <div className="text-sm text-slate-500 mb-5">
                <Balancer>{tier.subheader}</Balancer>
              </div>

              <Button fullWidth variant={tier.buttonVariant}>
                {tier.buttonText}
              </Button>
            </div>

            <div className="text-slate-900 dark:text-slate-200 font-medium mb-3">
              Includes:
            </div>
            <ul className="text-slate-600 dark:text-slate-400 text-sm space-y-3 grow">
              {tier.description.map((line: string) => (
                <Typography
                  component="li"
                  variant="subtitle1"
                  align="center"
                  className="flex items-center"
                  key={line}
                >
                  <svg
                    className="w-3 h-3 fill-emerald-500 mr-3 shrink-0"
                    viewBox="0 0 12 12"
                    xmlns="http://www.w3.org/2000/svg"
                  >
                    <path d="M10.28 2.28L3.989 8.575 1.695 6.28A1 1 0 00.28 7.695l3 3a1 1 0 001.414 0l7-7A1 1 0 0010.28 2.28z" />
                  </svg>

                  <span>
                    <Balancer>{line}</Balancer>
                  </span>
                </Typography>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </div>
  )
}

export default PriceTable
