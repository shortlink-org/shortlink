// @ts-nocheck
import React from 'react'
import get from 'lodash/get'
import { useSelector } from 'react-redux'
import { Layout } from 'components'
import Welcome from 'components/Profile/Welcome'
import Profile from 'components/Profile/Profile'
import Personal from 'components/Profile/Personal'
import Notifications from 'components/Profile/Notifications'
import withAuthSync from 'components/Private'

export function Profile() {
  const session = useSelector((state) => state.session)

  return (
    <Layout>
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
    </Layout>
  )
}

export default withAuthSync(() => Profile)
