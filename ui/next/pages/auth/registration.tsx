// @ts-nocheck
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
import FormControlLabel from '@mui/material/FormControlLabel'
import Checkbox from '@mui/material/Checkbox'
import Link from '@mui/material/Link'
import Grid from '@mui/material/Grid'
import type { NextPage } from 'next'
import { Layout } from 'components'
import {
  SelfServiceRegistrationFlow,
  SubmitSelfServiceRegistrationFlowBody
} from '@ory/client'
import { AxiosError } from 'axios'
import { useRouter, NextRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { Flow } from '../../components/ui/Flow'

import ory, { createLogoutHandler } from '../../pkg/sdk'
import { handleGetFlowError, handleFlowError } from '../../pkg/errors'

// Renders the registration page
const SignUp: NextPage = () => {
  const router = useRouter()

  // The "flow" represents a registration process and contains
  // information about the form we need to render (e.g. username + password)
  const [flow, setFlow] = useState<SelfServiceRegistrationFlow>()

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
        .getSelfServiceRegistrationFlow(String(flowId))
        .then(({ data }) => {
          // We received the flow - let's use its data and render the form!
          setFlow(data)
        })
        .catch(handleFlowError(router, 'registration', setFlow))
      return
    }

    // Otherwise we initialize it
    ory
      .initializeSelfServiceRegistrationFlowForBrowsers(
        returnTo ? String(returnTo) : undefined
      )
      .then(({ data }) => {
        setFlow(data)
      })
      .catch(handleFlowError(router, 'registration', setFlow))
  }, [flowId, router, router.isReady, returnTo, flow])

  const onSubmit = (values: SubmitSelfServiceRegistrationFlowBody) =>
    router
      // On submission, add the flow ID to the URL but do not navigate. This prevents the user loosing
      // his data when she/he reloads the page.
      .push(`/auth/registration?flow=${flow?.id}`, undefined, { shallow: true })
      .then(() =>
        ory
          .submitSelfServiceRegistrationFlow(String(flow?.id), values)
          .then(({ data }) => {
            // If we ended up here, it means we are successfully signed up!
            //
            // You can do cool stuff here, like having access to the identity which just signed up:
            console.log('This is the user session: ', data, data.identity)

            // For now however we just want to redirect home!
            return router.push(flow?.return_to || '/').then(() => {})
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

              {/*<form*/}
              {/*  // action={flow.ui.action}*/}
              {/*  // method={flow.ui.method}*/}
              {/*  // onSubmit={this.handleSubmit}*/}
              {/*>*/}
              {/*  {*/}
              {/*    // <TextField*/}
              {/*    //   name="csrf_token"*/}
              {/*    //   id="csrf_token"*/}
              {/*    //   type="hidden"*/}
              {/*    //   required*/}
              {/*    //   fullWidth*/}
              {/*    //   variant="outlined"*/}
              {/*    //   label="Csrf token"*/}
              {/*    //   value={csrfToken}*/}
              {/*    // />*/}
              {/*  }*/}

              {/*  <Grid container spacing={2}>*/}
              {/*    <Grid item xs={12} sm={6}>*/}
              {/*      <TextField*/}
              {/*        autoComplete="fname"*/}
              {/*        name="traits.name.first"*/}
              {/*        variant="outlined"*/}
              {/*        required*/}
              {/*        fullWidth*/}
              {/*        id="traits.name.first"*/}
              {/*        label="First Name"*/}
              {/*        autoFocus*/}
              {/*      />*/}
              {/*    </Grid>*/}
              {/*    <Grid item xs={12} sm={6}>*/}
              {/*      <TextField*/}
              {/*        variant="outlined"*/}
              {/*        required*/}
              {/*        fullWidth*/}
              {/*        id="traits.name.last"*/}
              {/*        label="Last Name"*/}
              {/*        name="traits.name.last"*/}
              {/*        autoComplete="lname"*/}
              {/*      />*/}
              {/*    </Grid>*/}
              {/*    <Grid item xs={12}>*/}
              {/*      <TextField*/}
              {/*        variant="outlined"*/}
              {/*        required*/}
              {/*        fullWidth*/}
              {/*        id="traits.email"*/}
              {/*        label="Email Address"*/}
              {/*        name="traits.email"*/}
              {/*        type="email"*/}
              {/*        autoComplete="email"*/}
              {/*      />*/}
              {/*    </Grid>*/}
              {/*    <Grid item xs={12}>*/}
              {/*      <TextField*/}
              {/*        variant="outlined"*/}
              {/*        required*/}
              {/*        fullWidth*/}
              {/*        name="password"*/}
              {/*        label="Password"*/}
              {/*        type="password"*/}
              {/*        id="password"*/}
              {/*        autoComplete="current-password"*/}
              {/*      />*/}
              {/*    </Grid>*/}
              {/*    <Grid item xs={12}>*/}
              {/*      <FormControlLabel*/}
              {/*        control={*/}
              {/*          <Checkbox value="allowExtraEmails" color="primary" />*/}
              {/*        }*/}
              {/*        label="I want to receive inspiration, marketing promotions and updates via email."*/}
              {/*      />*/}
              {/*    </Grid>*/}
              {/*  </Grid>*/}
              {/*  <Button*/}
              {/*    type="submit"*/}
              {/*    fullWidth*/}
              {/*    variant="contained"*/}
              {/*    color="primary"*/}
              {/*  >*/}
              {/*    Sign Up*/}
              {/*  </Button>*/}
              {/*  <Grid container justifyContent="flex-end">*/}
              {/*    <Grid item>*/}
              {/*      <Link*/}
              {/*        href="/next/auth/login"*/}
              {/*        variant="body2"*/}
              {/*        underline="hover"*/}
              {/*      >*/}
              {/*        <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">*/}
              {/*          Already have an account? Log in*/}
              {/*        </p>*/}
              {/*      </Link>*/}
              {/*    </Grid>*/}
              {/*  </Grid>*/}
              {/*</form>*/}
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}

export default SignUp
