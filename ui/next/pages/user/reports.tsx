// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'

export function ReportsContent() {
  return <Ready />
}

function Reports() {
  return <Layout content={ReportsContent()} />
}

export default withAuthSync(Reports)
