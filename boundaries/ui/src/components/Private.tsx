import React from 'react'
import { AxiosError } from 'axios'
import { useRouter } from 'next/navigation'
import { useEffect, useState } from 'react'

import { FrontendApi } from '@ory/client'

export default function withAuthSync<P extends object>(Child: React.ComponentType<P>) {
  return function WrappedComponent(props: P) {
    const [session, setSession] = useState<string>('No valid Ory Session was found.\nPlease sign in to receive one.')
    const [hasSession, setHasSession] = useState<boolean>(false)
    const router = useRouter()

    useEffect(() => {
      const ory = new FrontendApi()
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
              // it's a second factor
              return router.push('/auth/login?aal=aal2')
            case 401:
              // do nothing, the user is not logged in
              return
            default:
            // Otherwise, we nothitng - the error will be handled by the Flow component
          }

          // Something else happened!
          return Promise.reject(err)
        })
    }, [router])

    // If this is a token, we just render the component that was passed with all its props
    return React.createElement(Child, props)
  }
}
