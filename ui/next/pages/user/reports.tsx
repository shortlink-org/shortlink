// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'

export function Reports() {
  return (
    <Layout>
      <Ready />
    </Layout>
  )
}

export default withAuthSync(Reports)
