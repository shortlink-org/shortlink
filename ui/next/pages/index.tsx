import { Card, CardTitle, P, H2, H3, CodeBox } from '@ory/themes'
import { AxiosError } from 'axios'
import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import Link from 'next/link'

import { DocsButton, MarginCard, createLogoutHandler } from '../pkg'
import ory from '../pkg/sdk'

const Home: NextPage = () => {
  const [session, setSession] = useState<string>(
    'No valid Ory Session was found.\nPlease sign in to receive one.'
  )
  const [hasSession, setHasSession] = useState<boolean>(false)
  const router = useRouter()
  const onLogout = createLogoutHandler()

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
            return router.push('/auth/login?aal=aal2')
          case 401:
            // do nothing, the user is not logged in
            return
        }

        // Something else happened!
        return Promise.reject(err)
      })
  })

  return (
    <div className={'container-fluid'}>
      <Head>
        <title>Ory NextJS Integration Example</title>
        <meta name="description" content="NextJS + React + Vercel + Ory" />
      </Head>

      <MarginCard wide>
        <CardTitle>Welcome to Ory!</CardTitle>
        <P>
          Welcome to the Ory Managed UI. This UI implements a run-of-the-mill
          user interface for all self-service flows (login, registration,
          recovery, verification, settings). The purpose of this UI is to help
          you get started quickly. In the long run, you probably want to
          implement your own custom user interface.
        </P>
        <div className="row">
          <div className="col-md-4 col-xs-12">
            <div className="box">
              <H3>Documentation</H3>
              <P>
                Here are some useful documentation pieces that help you get
                started.
              </P>
            </div>
          </div>
          <div className="col-md-8 col-xs-12">
            <div className="box">
              <H3>Session Information</H3>
              <P>
                Below you will find the decoded Ory Session if you are logged
                in.
              </P>
              <CodeBox data-testid="session-content" code={session} />
            </div>
          </div>
        </div>
      </MarginCard>

      <Card wide>
        <H2>Other User Interface Screens</H2>
        <div className={'row'}>
          <DocsButton
            unresponsive
            testid="login"
            href="/next/auth/login"
            disabled={hasSession}
            title={'Login'}
          />
          <DocsButton
            unresponsive
            testid="sign-up"
            href="/next/auth/registration"
            disabled={hasSession}
            title={'Sign Up'}
          />
          <DocsButton
            unresponsive
            testid="recover-account"
            href="/next/auth/recovery"
            disabled={hasSession}
            title="Recover Account"
          />
          <DocsButton
            unresponsive
            testid="verify-account"
            href="/next/auth/verification"
            title="Verify Account"
          />
          <DocsButton
            unresponsive
            testid="account-settings"
            href="/next/auth/settings"
            disabled={!hasSession}
            title={'Account Settings'}
          />
          <DocsButton
            unresponsive
            testid="logout"
            onClick={onLogout}
            disabled={!hasSession}
            title={'Logout'}
          />
        </div>
      </Card>
    </div>
  )
}

export default Home
