// @ts-nocheck
import React, { useEffect, useState } from 'react'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Checkbox from '@material-ui/core/Checkbox'
import Link from '@material-ui/core/Link'
import { makeStyles } from '@material-ui/core/styles'
import { Layout } from 'components'
import SocialAuth from 'components/widgets/oAuthServices'
import { Configuration, PublicApi } from '@ory/kratos-client'

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(1),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(1, 0, 2),
  },
  csrf: {
    visibility: 'hidden',
  },
}))

export default function SignIn() {
  const classes = useStyles()

  const kratos = new PublicApi(
    new Configuration({ basePath: 'http://127.0.0.1:4433' }),
  )

  const [kratosState, setKratos] = useState()
  const [csrfToken, setCsrfToken] = useState()

  useEffect(() => {
    if (
      !new URL(document.location).searchParams.get('flow') &&
      new URL(document.location).href.indexOf('login') !== -1
    ) {
      window.location.href = 'http://127.0.0.1:4433/self-service/login/browser'
    }

    // @ts-ignore
    const flowId = new URL(document.location).searchParams.get('flow')

    // @ts-ignore
    kratos
      .getSelfServiceLoginFlow(flowId)
      .then(({ status, data: flow }) => {
        if (status === 404 || status === 410 || status === 403) {
          window.location.replace(
            'http://127.0.0.1:4433/self-service/registration/browser',
          )
        }
        if (status !== 200) {
          return Promise.reject(flow)
        }

        // @ts-ignore
        setKratos(flow)
        // @ts-ignore
        setCsrfToken(flow.ui.nodes[0].attributes.value)
      })}, [csrfToken])

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
                Login{' '}
                <span className="text-gray-400 font-light">to your account</span>
              </h3>

              <form
                className={classes.form}
                action={kratosState && kratosState.ui.action}
                method={kratosState && kratosState.ui.method}
              >
                <TextField
                  name="csrf_token"
                  id="csrf_token"
                  type="hidden"
                  required
                  fullWidth
                  variant="outlined"
                  label="Csrf token"
                  value={csrfToken}
                  className={classes.csrf}
                />

                <TextField
                  name="method"
                  id="method"
                  type="hidden"
                  required
                  fullWidth
                  variant="outlined"
                  label="method"
                  value="password"
                  className={classes.csrf}
                />

                <div className="py-2 space-y-6">
                  <SocialAuth />

                  <div className="flex flex-row items-center justify-center">
                    <hr className="w-28 border-gray-300 block" />
                    <label className="mx-2 text-sm text-gray-500">
                      Or continue with
                    </label>
                    <hr className="w-28 border-gray-300 block" />
                  </div>
                </div>

                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  id="password_identifier"
                  label="Email Address"
                  name="password_identifier"
                  type="email"
                  autoComplete="email"
                />
                <TextField
                  variant="outlined"
                  margin="normal"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                />
                <FormControlLabel
                  control={<Checkbox value="remember" color="primary" />}
                  label="Remember me"
                />

                <Button
                  type="submit"
                  fullWidth
                  variant="contained"
                  color="primary"
                  className={classes.submit}
                >
                  Log In
                </Button>

                <div className="flex items-center justify-between">
                  <Link href="/next/auth/forgot" variant="body2">
                    <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">
                      Forgot password?
                    </p>
                  </Link>

                  <Link href="/next/auth/registration" variant="body2">
                    <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">
                      Don't have an account? Sign Up
                    </p>
                  </Link>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
}
