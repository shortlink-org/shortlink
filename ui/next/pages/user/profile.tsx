import React from 'react'
import { useSelector } from "react-redux"
import { Layout } from 'components'
import Welcome from 'components/Profile/Welcome'
import Profile from 'components/Profile/Profile'
import Personal from 'components/Profile/Personal'
import Notifications from 'components/Profile/Notifications'
import withAuthSync from 'components/Private'

export function ProfileContent() {
  // @ts-ignore
  const session = useSelector((state) => state.session)

  console.warn('session', session)

  return (
    <React.Fragment>
      <Welcome nickname={session.kratos.identity.traits.name.first} />

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
    </React.Fragment>
  )
}

export default withAuthSync(() => <Layout content={ProfileContent()} />)
