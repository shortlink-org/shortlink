// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'
import { NextSeo } from "next-seo";

export function Intargrations() {
  return (
    <Layout>
      <NextSeo
        title="Intargrations"
        description="Intargrations page for your account."
      />
      <Ready />
    </Layout>
  )
}

export default withAuthSync(Intargrations)
