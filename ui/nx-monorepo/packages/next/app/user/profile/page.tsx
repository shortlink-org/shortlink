'use client'

// @ts-nocheck
import React, { Fragment, useEffect, useState } from 'react'
import get from 'lodash/get'
import Welcome from 'components/Profile/Welcome'
import Profile from 'components/Profile/Profile'
import Personal from 'components/Profile/Personal'
import Notifications from 'components/Profile/Notifications'
import withAuthSync from 'components/Private'
import ory from '../../../pkg/sdk'
import { AxiosError } from 'axios'
// @ts-ignore
import { Header } from '@shortlink-org/ui-kit'

function ProfileContent() {
  const [session, setSession] = useState<string>(
    'No valid Ory Session was found.\nPlease sign in to receive one.',
  )
  const [hasSession, setHasSession] = useState<boolean>(false) // eslint-disable-line

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
            // it's second factor
            return router.push('/login?aal=aal2')
          case 401:
            // do nothing, the user is not logged in
            return
        }

        // Something else happened!
        Promise.reject(err)
      })
  }, [])

  return (
    <Fragment>
      {/*<NextSeo*/}
      {/*  title="Profile"*/}
      {/*  description="Profile page for your account."*/}
      {/*  openGraph={{*/}
      {/*    title: 'Profile',*/}
      {/*    description: 'Profile page for your account.',*/}
      {/*    type: 'profile',*/}
      {/*    profile: {*/}
      {/*      firstName: 'John',*/}
      {/*      lastName: 'Doe',*/}
      {/*      username: 'johndoe',*/}
      {/*    },*/}
      {/*  }}*/}
      {/*/>*/}

      <Header title="Profile" />

      <Welcome
        nickname={get(session, 'kratos.identity.traits.name.first', 'default')}
      />

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
    </Fragment>
  )
}

export default withAuthSync(ProfileContent)
