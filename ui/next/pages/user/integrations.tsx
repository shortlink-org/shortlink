// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'

export function IntargrationsContent() {
  return <Ready />
}

function Intargrations() {
  return <Layout content={IntargrationsContent()} />
}

export default withAuthSync(Intargrations)
