import React from 'react'
import { Layout } from 'components'
import PaymentMethod from 'components/Billing/PaymentMethod'
import Discounted from 'components/Billing/Discounted'
import withAuthSync from 'components/Private'

export function BillingContent() {
  return (
    <React.Fragment>
      <Discounted />

      <PaymentMethod />
    </React.Fragment>
  )
}

function Billing() {
  return <Layout content={BillingContent()} />
}

export default withAuthSync(Billing)
