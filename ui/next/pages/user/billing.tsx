// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import PaymentMethod from 'components/Billing/PaymentMethod'
import Discounted from 'components/Billing/Discounted'
import withAuthSync from 'components/Private'
import { NextSeo } from "next-seo";

export function Billing() {
  return (
    <Layout>
      <NextSeo
        title="Billing"
        description="Billing page for your account."
      />
      <Discounted />

      <PaymentMethod />
    </Layout>
  )
}

export default withAuthSync(Billing)
