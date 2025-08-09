'use client'

// @ts-ignore
import { Header } from '@shortlink-org/ui-kit'
import { AxiosError } from 'axios'
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'

import { Layout } from 'components'
import withAuthSync from 'components/Private'
import Security from 'components/Profile/Security'
import ory from 'pkg/sdk'

// <NextSeo
// title="Security"
// description="Security page for your account."
// openGraph={{
//   title: 'Security',
//     description: 'Security page for your account.',
//     type: 'website',
// }}
// />

function SecurityContent() {
  const router = useRouter()

  const [session, setSession] = useState<string>('No valid Ory Session was found.\nPlease sign in to receive one.')
  const [hasSession, setHasSession] = useState<boolean>(false)

  useEffect(() => {
    ory
      .toSession()
      .then(({ data }) => {
        setSession(JSON.stringify(data, null, 2))
        setHasSession(true)
      })
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 403:
          // This is a legacy error code thrown. See code 422 for
          // more details.
          case 422:
            // This status code is returned when we are trying to
            // validate a session which has not yet completed
            // its second factor
            return router.push('/login?aal=aal2')
          case 401:
            // do nothing, the user is not logged in
            return
          default:
          // Otherwise, we nothitng - the error will be handled by the Flow component
        }

        // Something else happened!
        Promise.reject(err)
      })
  }, [router])

  return (
    <Layout>
      <Header title="Security" />

      <Security session={session} />
    </Layout>
  )
}

export default withAuthSync(SecurityContent)
