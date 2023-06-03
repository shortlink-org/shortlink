import React, { Fragment } from 'react'
import Header from '../../components/Page/Header'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <Fragment>
      <Header title="Admin users" />
      <section className="text-gray-600 body-font">
        {children}
      </section>
    </Fragment>
  )
}
