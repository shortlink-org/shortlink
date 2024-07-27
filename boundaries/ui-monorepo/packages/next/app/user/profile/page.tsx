'use client'

import { AxiosError } from 'axios'
import get from 'lodash/get'
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
// @ts-ignore
import { Header } from '@shortlink-org/ui-kit'

import ory from 'pkg/sdk'
import withAuthSync from 'components/Private'
import Notifications from 'components/Profile/Notifications'
import Personal from 'components/Profile/Personal'
import Profile from 'components/Profile/Profile'
import Welcome from 'components/Profile/Welcome'

// <NextSeo
// title="Profile"
// description="Profile page for your account."
// openGraph={{
//   title: 'Profile',
//     description: 'Profile page for your account.',
//     type: 'profile',
//     profile: {
//     firstName: 'John',
//       lastName: 'Doe',
//       username: 'johndoe',
//   },
// }}
// />

function ProfileContent() {
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
    <>
      <Header title="Profile" />

      <Welcome nickname={get(session, 'kratos.identity.traits.name.first', 'default')} />

      <Profile />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200" />
        </div>
      </div>

      <Personal session={session} />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200" />
        </div>
      </div>

      <Notifications />
    </>
  )
}

export default withAuthSync(ProfileContent)
