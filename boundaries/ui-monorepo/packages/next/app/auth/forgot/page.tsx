'use client'

// @ts-nocheck
import { RecoveryFlow, UpdateRecoveryFlowBody } from '@ory/client'
import { AxiosError } from 'axios'
import type { NextPage } from 'next'
import Link from 'next/link'
import { useRouter, useSearchParams } from 'next/navigation'
import React, { useEffect, useState } from 'react'

import { Flow } from 'components/ui/Flow'
import { handleFlowError } from 'pkg/errors'
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

const Page: NextPage = () => {
  const [flow, setFlow] = useState<RecoveryFlow>()

  // Get ?flow=... from the URL
  const router = useRouter()
  const searchParams = useSearchParams()
  const flowId = searchParams.get('flow')
  const returnTo = searchParams.get('return_to')

  useEffect(() => {
    // If the router is not ready yet, or we already have a flow, do nothing.
    if (flow) {
      return
    }

    // If ?flow=.. was in the URL, we fetch it
    if (flowId) {
      ory
        .getRecoveryFlow({ id: String(flowId) })
        .then(({ data }) => {
          setFlow(data)
        })
        .catch(handleFlowError(router, 'recovery', setFlow))
      return
    }

    // Otherwise, we initialize it
    ory
      .createBrowserRecoveryFlow({
        returnTo: String(returnTo || ''),
      })
      .then(({ data }) => {
        setFlow(data)
      })
      .catch(handleFlowError(router, 'recovery', setFlow))
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
  }, [flowId, router, returnTo, flow])

  const onSubmit = (values: UpdateRecoveryFlowBody) => {
    router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // his data when she/he reloads the page.
      .push(`/auth/forget?flow=${flow?.id}`)

    ory
      .updateRecoveryFlow({
        flow: String(flow?.id),
        updateRecoveryFlowBody: values,
      })
      .then(({ data }) => {
        // Form submission was successful, show the message to the user!
        setFlow(data)
      })
      .catch(handleFlowError(router, 'recovery', setFlow))
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 400:
            // Status code 400 implies the form validation had an error
            // @ts-ignore
            setFlow(err.response?.data)
            return
          default:
          // Otherwise, we nothitng - the error will be handled by the Flow component
        }

        throw err
      })
  }

  return (
    <>
      {/*<NextSeo title="Forgot Password" description="Forgot Password" />*/}

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
                <h3 className="text-2xl font-bold inline-block">Reset Password</h3>
                <p className="text-gray-500 whitespace-no-wrap">Forgotten Password? No prob!</p>
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
                Enter <span className="text-gray-400 font-light">your email below</span>
              </h3>

              <Flow onSubmit={onSubmit} flow={flow} />

              <div className="flex items-center justify-between">
                <Link href="/auth/login" legacyBehavior>
                  <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Log in
                  </p>
                </Link>

                <Link href="/auth/registration" legacyBehavior>
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

export default Page
