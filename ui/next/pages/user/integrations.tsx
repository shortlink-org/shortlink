// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'

export function Intargrations() {
  return (
    <Layout>
      <Ready />
    </Layout>
  )
}

export default withAuthSync(Intargrations)
