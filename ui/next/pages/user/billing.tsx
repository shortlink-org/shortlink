// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import PaymentMethod from 'components/Billing/PaymentMethod'
import Discounted from 'components/Billing/Discounted'
import withAuthSync from 'components/Private'

export function Billing() {
  return (
    <Layout>
      <Discounted />

      <PaymentMethod />
    </Layout>
  )
}

export default withAuthSync(Billing)
