// @ts-nocheck
import type { NextPage } from 'next'
import { Layout } from 'components'
import { RegistrationFlow, UpdateRegistrationFlowBody } from '@ory/client'
import { AxiosError } from 'axios'
import Grid from '@mui/material/Grid'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import Link from 'next/link'
import { Flow } from '../../components/ui/Flow'

import ory from '../../pkg/sdk'
import { handleFlowError } from '../../pkg/errors'
import { BreadcrumbJsonLd, NextSeo } from 'next-seo'

// Renders the registration page
const SignUp: NextPage = () => {
  const router = useRouter()

  // The "flow" represents a registration process and contains
  // information about the form we need to render (e.g. username + password)
  const [flow, setFlow] = useState<RegistrationFlow>()

  // Get ?flow=... from the URL
  const { flow: flowId, return_to: returnTo } = router.query

  // In this effect we either initiate a new registration flow, or we fetch an existing registration flow.
  useEffect(() => {
    // If the router is not ready yet, or we already have a flow, do nothing.
    if (!router.isReady || flow) {
      return
    }

    // If ?flow=.. was in the URL, we fetch it
    if (flowId) {
      ory
        .getRegistrationFlow({ id: String(flowId) })
        .then(({ data }) => {
          // We received the flow - let's use its data and render the form!
          setFlow(data)
        })
        .catch(handleFlowError(router, 'registration', setFlow))
      return
    }

    // Otherwise we initialize it
    ory
      .createBrowserRegistrationFlow({
        returnTo: returnTo ? String(returnTo) : undefined,
      })
      .then(({ data }) => {
        setFlow(data)
      })
      .catch(handleFlowError(router, 'registration', setFlow))
  }, [flowId, router, router.isReady, returnTo, flow])

  const onSubmit = async (values: UpdateRegistrationFlowBody) => {
    await router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // his data when she/he reloads the page.
      .push(`/auth/registration?flow=${flow?.id}`, undefined, { shallow: true })

    ory
      .updateRegistrationFlow({
        flow: String(flow?.id),
        updateRegistrationFlowBody: values,
      })
      .then(async ({ data }) => {
        // If we ended up here, it means we are successfully signed up!
        //
        // You can do cool stuff here, like having access to the identity which just signed up:
        console.log('This is the user session: ', data, data.identity)

        // continue_with is a list of actions that the user might need to take before the registration is complete.
        // It could, for example, contain a link to the verification form.
        if (data.continue_with) {
          // eslint-disable-next-line no-restricted-syntax
          for (const item of data.continue_with) {
            switch (item.action) {
              case 'show_verification_ui':
                // eslint-disable-next-line no-await-in-loop
                await router.push(`/auth/verification?flow=${item.flow.id}`)
                return
            }
          }
        }

        // If continue_with did not contain anything, we can just return to the home page.
        await router.push(flow?.return_to || '/')
      })
      .catch(handleFlowError(router, 'registration', setFlow))
      .catch((err: AxiosError) => {
        // If the previous handler did not catch the error it's most likely a form validation error
        if (err.response?.status === 400) {
          // Yup, it is!
          setFlow(err.response?.data)
          return
        }

        return Promise.reject(err)
      })
  }

  return (
    <Layout>
      <NextSeo title="Registration" description="Registration a new account" />
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
                <h3 className="text-4xl font-bold inline-block">Register</h3>
                <p className="text-gray-500 whitespace-no-wrap">
                  Signup for an Account
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
                Register{' '}
                <span className="text-gray-400 font-light">for an account</span>
              </h3>

              <Flow onSubmit={onSubmit} flow={flow} />

              <Grid container justifyContent="flex-end">
                <Grid item>
                  <Link
                    href="/auth/login"
                    variant="body2"
                    underline="hover"
                    legacyBehavior
                  >
                    <p className="cursor-pointer no-underline hover:underline mt-4 text-sm font-medium text-indigo-600 hover:text-indigo-500">
                      Already have an account? Log in
                    </p>
                  </Link>
                </Grid>
              </Grid>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default SignUp
