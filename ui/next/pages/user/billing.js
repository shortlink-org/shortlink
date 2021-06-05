import React from 'react'
import { Layout } from '../../components';
import PaymentMethod from '../../components/Billing/PaymentMethod';
import Discounted from '../../components/Billing/Discounted';
import Ready from '../../components/Landing/Ready';

export function BillingContent() {
  return (
    <React.Fragment>
      <Discounted />

      <PaymentMethod />
    </React.Fragment>
  )
}

export default function Billing() {
  return <Layout content={BillingContent()} />;
}
