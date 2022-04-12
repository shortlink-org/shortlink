// @ts-nocheck
import React, { useEffect, useState } from 'react'
import type { NextPage } from 'next'
import { Layout } from 'components'
import { useRouter } from 'next/router'
import {
  SelfServiceLoginFlow,
  SubmitSelfServiceLoginFlowBody,
} from '@ory/client'

import ory, { useCreateLogoutHandler } from '../../pkg/sdk'
import { handleGetFlowError, handleFlowError } from '../../pkg/errors'
import { Flow } from '../../components/ui/Flow'

const SignIn: NextPage = () => {
  const [flow, setFlow] = useState<SelfServiceLoginFlow>()

  // Get ?flow=... from the URL
  const router = useRouter()
  const {
    return_to: returnTo,
    flow: flowId,
    // Refresh means we want to refresh the session. This is needed, for example, when we want to update the password
    // of a user.
    refresh,
    // AAL = Authorization Assurance Level. This implies that we want to upgrade the AAL, meaning that we want
    // to perform two-factor authentication/verification.
    aal,
  } = router.query

  // This might be confusing, but we want to show the user an option
  // to sign out if they are performing two-factor authentication!
  const onLogout = useCreateLogoutHandler([aal, refresh])

  useEffect(() => {
    // If the router is not ready yet, or we already have a flow, do nothing.
    if (!router.isReady || flow) {
      return
    }

    // If ?flow=.. was in the URL, we fetch it
    if (flowId) {
      ory
        .getSelfServiceLoginFlow(String(flowId))
        .then(({ data }) => {
          setFlow(data)
        })
        .catch(handleGetFlowError(router, 'login', setFlow))
      return
    }

    // Otherwise we initialize it
    ory
      .initializeSelfServiceLoginFlowForBrowsers(
        Boolean(refresh),
        aal ? String(aal) : undefined,
        returnTo ? String(returnTo) : undefined,
      )
      .then(({ data }) => {
        setFlow(data)
      })
      .catch(handleFlowError(router, 'login', setFlow))
  }, [flowId, router, router.isReady, aal, refresh, returnTo, flow])

  const onSubmit = (values: SubmitSelfServiceLoginFlowBody) =>
    router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // his data when she/he reloads the page.
      .push(`/auth/login?flow=${flow?.id}`, undefined, { shallow: true })
      .then(() =>
        ory
          .submitSelfServiceLoginFlow(String(flow?.id), undefined, values)
          // We logged in successfully! Let's bring the user home.
          .then((res) => {
            if (flow?.return_to) {
              window.location.href = flow?.return_to
              return
            }
            router.push('/')
          })
          .then(() => {})
          .catch(handleFlowError(router, 'login', setFlow))
          .catch((err: AxiosError) => {
            // If the previous handler did not catch the error it's most likely a form validation error
            if (err.response?.status === 400) {
              // Yup, it is!
              setFlow(err.response?.data)
              return
            }

            return Promise.reject(err)
          }),
      )

  return (
    <Layout>
      <div className="flex h-full p-4 rotate">
        <div className="sm:max-w-xl md:max-w-3xl w-full m-auto">
          <div className="flex items-stretch bg-white rounded-lg shadow-lg overflow-hidden border-t-4 border-indigo-500 sm:border-0">
            <div
              className="flex hidden overflow-hidden relative sm:block w-4/12 md:w-5/12 bg-gray-600 text-gray-300 py-4 bg-cover bg-center"
              style={{
                backgroundImage:
                  "url('https://images.unsplash.com/photo-1477346611705-65d1883cee1e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1350&q=80')",
              }}
            >
              <div className="flex-1 absolute bottom-0 text-white p-10">
                <h3 className="text-4xl font-bold inline-block">Login</h3>
                <p className="text-gray-500 whitespace-no-wrap">
                  Welcome back!
                </p>
              </div>
              <svg
                className="absolute animate h-full w-4/12 sm:w-2/12 right-0 inset-y-0 fill-current text-white"
                viewBox="0 0 100 100"
                xmlns="http://www.w3.org/2000/svg"
                preserveAspectRatio="none"
              >
                <polygon points="0,0 100,100 100,0" />
              </svg>
            </div>

            <div className="flex-1 p-6 sm:p-10 sm:py-12">
              <h3 className="text-xl text-gray-700 font-bold mb-6">
                Login{' '}
                <span className="text-gray-400 font-light">
                  to your account
                </span>
              </h3>

              <Flow onSubmit={onSubmit} flow={flow} />

              {/*<form*/}
              {/*  // action={flow.ui.action}*/}
              {/*  // method={flow.ui.method}*/}
              {/*  // onSubmit={handleSubmit}*/}
              {/*>*/}
              {/*  <TextField*/}
              {/*    name="csrf_token"*/}
              {/*    id="csrf_token"*/}
              {/*    type="hidden"*/}
              {/*    required*/}
              {/*    fullWidth*/}
              {/*    variant="outlined"*/}
              {/*    label="Csrf token"*/}
              {/*    // value={csrfToken}*/}
              {/*  />*/}

              {/*  <TextField*/}
              {/*    name="method"*/}
              {/*    id="method"*/}
              {/*    type="hidden"*/}
              {/*    required*/}
              {/*    fullWidth*/}
              {/*    variant="outlined"*/}
              {/*    label="method"*/}
              {/*    value="password"*/}
              {/*  />*/}

              {/*  <div className="py-2 space-y-6">*/}
              {/*    <SocialAuth />*/}

              {/*    <div className="flex flex-row items-center justify-center">*/}
              {/*      <hr className="w-28 border-gray-300 block" />*/}
              {/*      <label className="mx-2 text-sm text-gray-500">*/}
              {/*        Or continue with*/}
              {/*      </label>*/}
              {/*      <hr className="w-28 border-gray-300 block" />*/}
              {/*    </div>*/}
              {/*  </div>*/}

              {/*  <TextField*/}
              {/*    variant="outlined"*/}
              {/*    margin="normal"*/}
              {/*    required*/}
              {/*    fullWidth*/}
              {/*    id="password_identifier"*/}
              {/*    label="Email Address"*/}
              {/*    name="password_identifier"*/}
              {/*    type="email"*/}
              {/*    autoComplete="email"*/}
              {/*  />*/}
              {/*  <TextField*/}
              {/*    variant="outlined"*/}
              {/*    margin="normal"*/}
              {/*    required*/}
              {/*    fullWidth*/}
              {/*    name="password"*/}
              {/*    label="Password"*/}
              {/*    type="password"*/}
              {/*    id="password"*/}
              {/*    autoComplete="current-password"*/}
              {/*  />*/}
              {/*  <FormControlLabel*/}
              {/*    control={<Checkbox value="remember" color="primary" />}*/}
              {/*    label="Remember me"*/}
              {/*  />*/}

              {/*  <Button*/}
              {/*    type="submit"*/}
              {/*    fullWidth*/}
              {/*    variant="contained"*/}
              {/*  >*/}
              {/*    {(() => {*/}
              {/*      if (flow?.refresh) {*/}
              {/*        return 'Confirm Action'*/}
              {/*      } else if (flow?.requested_aal === 'aal2') {*/}
              {/*        return 'Two-Factor Authentication'*/}
              {/*      }*/}
              {/*      return 'Sign In'*/}
              {/*    })()}*/}
              {/*  </Button>*/}

              {/*  <div className="flex items-center justify-between">*/}
              {/*    <Link*/}
              {/*      href="/next/auth/forgot"*/}
              {/*      variant="body2"*/}
              {/*      underline="hover"*/}
              {/*    >*/}
              {/*      <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">*/}
              {/*        Forgot password?*/}
              {/*      </p>*/}
              {/*    </Link>*/}

              {/*    <Link*/}
              {/*      href="/next/auth/registration"*/}
              {/*      variant="body2"*/}
              {/*      underline="hover"*/}
              {/*    >*/}
              {/*      <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">*/}
              {/*        Don't have an account? Sign Up*/}
              {/*      </p>*/}
              {/*    </Link>*/}
              {/*  </div>*/}
              {/*</form>*/}
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default SignIn
