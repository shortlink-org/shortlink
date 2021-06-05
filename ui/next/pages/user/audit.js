import React from 'react'
import { Layout } from '../../components';
import Ready from "../../components/Landing/Ready";

export function BillingContent() {
  return (
    <Ready />
  )
}

export default function Billing() {
  return <Layout content={BillingContent()} />;
}
