import { useEffect, useState } from 'react'
import ory from '../pkg/sdk'
import { AxiosError } from 'axios'
import { useRouter } from 'next/router'

export default function withAuthSync(Child: any) {
  return (props?: any) => {
    const [session, setSession] = useState<string>( // eslint-disable-line
      'No valid Ory Session was found.\nPlease sign in to receive one.',
    )
    const [hasSession, setHasSession] = useState<boolean>(false) // eslint-disable-line
    const router = useRouter()

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
    }, [])

    // If this is a token, we just render the component that was passed with all its props
    return <Child {...props} />
  }
}
