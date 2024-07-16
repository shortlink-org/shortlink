'use client'

// @ts-nocheck
import { LoginFlow, UpdateLoginFlowBody } from '@ory/client'
import type { NextPage } from 'next'
import Link from 'next/link'
import { useRouter, useSearchParams } from 'next/navigation'
import React, { useEffect, useState } from 'react'
import { AxiosError } from 'axios'

import { Flow } from 'components/ui/Flow'
import { handleGetFlowError, handleFlowError } from 'pkg/errors'
import ory from 'pkg/sdk'

// <BreadcrumbJsonLd
// itemListElements={[
//     {
//       position: 1,
//       name: 'Login page',
//       item: 'https://shortlink.best/next/auth/login',
//     },
// {
//   position: 2,
//     name: 'Forgot Password',
//   item: 'https://shortlink.best/next/auth/forgot',
// },
// {
//   position: 3,
//     name: 'Registration page',
//   item: 'https://shortlink.best/next/auth/registration',
// },
// ]}
// />

const SignIn: NextPage = () => {
  const [flow, setFlow] = useState<LoginFlow>()

  // Get ?flow=... from the URL
  const router = useRouter()
  const searchParams = useSearchParams()
  const flowId = searchParams.get('flow')
  const returnTo = searchParams.get('return_to')
  // Refresh means we want to refresh the session. This is needed, for example, when we want to update the password
  // of a user.
  const refresh = searchParams.get('refresh')
  // AAL = Authorization Assurance Level. This implies that we want to upgrade the AAL, meaning that we want
  // to perform two-factor authentication/verification.
  const aal = searchParams.get('aal')

  useEffect(() => {
    // If the router is not ready yet, or we already have a flow, do nothing.
    if (flow) {
      return
    }

    // If ?flow=.. was in the URL, we fetch it
    if (flowId) {
      ory // @ts-ignore
        .getLoginFlow(flowId)
        .then(({ data }) => {
          setFlow(data)
        })
        .catch(handleGetFlowError(router, 'login', setFlow))
      return
    }

    // Otherwise we initialize it
    ory
      .createBrowserLoginFlow({
        refresh: Boolean(refresh),
        aal: aal ? String(aal) : undefined,
        returnTo: returnTo ? String(returnTo) : undefined,
      })
      .then(({ data }) => {
        setFlow(data)
      })
      .catch(handleFlowError(router, 'login', setFlow))
  }, [flowId, router, aal, refresh, returnTo, flow])

  const onSubmit = (values: UpdateLoginFlowBody) => {
    router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // his data when she/he reloads the page.
      .push(`/auth/login?flow=${flow?.id}`)

    ory
      .updateLoginFlow({
        flow: String(flow?.id),
        updateLoginFlowBody: values,
      })
      // We logged in successfully! Let's bring the user home.
      .then(() => {
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
          // @ts-ignore
          setFlow(err.response?.data)
          return
        }

        return Promise.reject(err)
      })
  }

  return (
    <>
      {/*<NextSeo title="Login" description="Login to your account" />*/}

      <div className="flex h-full p-4 rotate">
        <div className="sm:max-w-xl md:max-w-3xl w-full m-auto">
          <div className="flex items-stretch bg-white dark:bg-gray-800 rounded-lg shadow-lg overflow-hidden border-t-4 border-indigo-500 sm:border-0">
            <div
              className="flex hidden overflow-hidden relative sm:block w-4/12 md:w-5/12 bg-gray-600 text-gray-300 py-4 bg-cover bg-center"
              style={{
                backgroundImage:
                  "url('https://images.unsplash.com/photo-1477346611705-65d1883cee1e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1350&q=80')",
              }}
            >
              <div className="flex-1 absolute bottom-0 text-white p-10">
                <h3 className="text-4xl font-bold inline-block">Login</h3>
                <p className="text-gray-500 whitespace-no-wrap">Welcome back!</p>
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
                Login <span className="text-gray-400 font-light">to your account</span>
              </h3>

              <Flow key="login" onSubmit={onSubmit} flow={flow} />

              <div className="flex items-center justify-between">
                <Link href="/auth/forgot">
                  <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Forgot password?
                  </p>
                </Link>

                <Link href="/auth/registration">
                  <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Don't have an account? Sign Up
                  </p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default SignIn
