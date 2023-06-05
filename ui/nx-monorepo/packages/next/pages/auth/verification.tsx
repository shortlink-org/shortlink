// @ts-nocheck
import React, { useEffect, useState } from 'react'
import type { NextPage } from 'next'
import Link from 'next/link'
import { Layout } from 'components'
import { useRouter } from 'next/router'
import { VerificationFlow, UpdateVerificationFlowBody } from '@ory/client'

import ory from '../../pkg/sdk'
import { handleGetFlowError, handleFlowError } from '../../pkg/errors'
import { Flow } from '../../components/ui/Flow'
import { BreadcrumbJsonLd, NextSeo } from 'next-seo'

const Verification: NextPage = () => {
  const [flow, setFlow] = useState<VerificationFlow>()

  // Get ?flow=... from the URL
  const router = useRouter()
  const { flow: flowId, return_to: returnTo } = router.query

  useEffect(() => {
    // If the router is not ready yet, or we already have a flow, do nothing.
    if (!router.isReady || flow) {
      return
    }

    // If ?flow=.. was in the URL, we fetch it
    if (flowId) {
      ory
        .getVerificationFlow({ id: String(flowId) })
        .then(({ data }) => {
          setFlow(data)
        })
        .catch((err: AxiosError) => {
          switch (err.response?.status) {
            case 410:
            // Status code 410 means the request has expired - so let's load a fresh flow!
            case 403:
              // Status code 403 implies some other issue (e.g. CSRF) - let's reload!
              return router.push('/verification')
          }

          throw err
        })
      return
    }

    // Otherwise we initialize it
    ory
      .createBrowserVerificationFlow({
        returnTo: returnTo ? String(returnTo) : undefined,
      })
      .then(({ data }) => {
        setFlow(data)
      })
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 400:
            // Status code 400 implies the user is already signed in
            return router.push('/')
        }

        throw err
      })
  }, [flowId, router, router.isReady, returnTo, flow])

  const onSubmit = async (values: UpdateVerificationFlowBody) => {
    await router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // their data when they reload the page.
      .push(`/verification?flow=${flow?.id}`, undefined, { shallow: true })

    ory
      .updateVerificationFlow({
        flow: String(flow?.id),
        updateVerificationFlowBody: values,
      })
      .then(({ data }) => {
        // Form submission was successful, show the message to the user!
        setFlow(data)
      })
      .catch((err: AxiosError) => {
        switch (err.response?.status) {
          case 400:
            // Status code 400 implies the form validation had an error
            setFlow(err.response?.data)
            return
          case 410:
            // eslint-disable-next-line no-case-declarations
            const newFlowID = err.response.data.use_flow_id
            router
              // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
              // their data when they reload the page.
              .push(`/verification?flow=${newFlowID}`, undefined, {
                shallow: true,
              })

            ory
              .getVerificationFlow({ id: newFlowID })
              .then(({ data }) => setFlow(data))
            return
        }

        throw err
      })
  }

  return (
    <Layout>
      <NextSeo title="Verification" description="Verify your account" />
      <BreadcrumbJsonLd
        itemListElements={[
          {
            position: 1,
            name: 'Login page',
            item: 'https://shortlink.best/next/auth/login',
          },
          {
            position: 2,
            name: 'Forgot Password',
            item: 'https://shortlink.best/next/auth/forgot',
          },
          {
            position: 3,
            name: 'Registration page',
            item: 'https://shortlink.best/next/auth/registration',
          },
        ]}
      />

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
                <p className="text-gray-500 whitespace-no-wrap">
                  Verification page for your account
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
                Verification{' '}
                <span className="text-gray-400 font-light">
                  to your account
                </span>
              </h3>

              <Flow key="verification" onSubmit={onSubmit} flow={flow} />

              <div className="flex items-center justify-between">
                <Link href="/auth/forgot" variant="body2" legacyBehavior>
                  <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Forgot password?
                  </p>
                </Link>

                <Link href="/auth/registration" variant="body2" legacyBehavior>
                  <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Don't have an account? Sign Up
                  </p>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default Verification
