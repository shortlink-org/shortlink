import { Configuration, V0alpha2Api } from '@ory/client'
import { useState, useEffect, DependencyList } from 'react'
import { useRouter } from 'next/router'
import { AxiosError } from 'axios'

const KRATOS_PUBLIC_API =
  process.env.KRATOS_PUBLIC_API || 'http://127.0.0.1:4433'

const ory = new V0alpha2Api(
  new Configuration({
    basePath: KRATOS_PUBLIC_API,
    baseOptions: {
      withCredentials: true,
    },
  }),
)

export default ory

// AUTH ================================================================================================================
// Returns a function which will log the user out
export function useCreateLogoutHandler(deps?: DependencyList) {
  const [logoutToken, setLogoutToken] = useState<string>('')
  const router = useRouter()

  useEffect(() => {
    ory
      .createSelfServiceLogoutFlowUrlForBrowsers()
      .then(({ data }) => {
        setLogoutToken(data.logout_token)
      })
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 401:
            // do nothing, the user is not logged in
            return
        }

        // Something else happened!
        return Promise.reject(err)
      })
  }, deps)

  return () => {
    if (logoutToken) {
      ory
        .submitSelfServiceLogoutFlow(logoutToken)
        .then(() => router.push('/login'))
        .then(() => router.reload())
    }
  }
}
