import React, {useEffect} from 'react'
import Button from '@material-ui/core/Button'
import TextField from '@material-ui/core/TextField'
import FormControlLabel from '@material-ui/core/FormControlLabel'
import Checkbox from '@material-ui/core/Checkbox'
import Link from '@material-ui/core/Link'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import { makeStyles } from '@material-ui/core/styles'
import { Layout } from 'components'
import SocialAuth from 'components/widgets/oAuthServices'
import {useDispatch} from "react-redux"
import {fetchLinkList} from "../../store"
import {loginAuth} from "../../store/actions/session"

const useStyles = makeStyles(theme => ({
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
}))

export function SignInPageContent() {
  const classes = useStyles()

  // const dispatch = useDispatch()
  // useEffect(() => {
  //   dispatch(loginAuth())
  // }, [dispatch])

  return (
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

            <form className={classes.form} noValidate>
              <div className="py-2 space-y-6">
                <SocialAuth />

                <div className="flex flex-row items-center justify-center">
                  <hr className="w-28 border-gray-300 block" />
                  <label className="mx-2 text-sm text-gray-500">Or continue with</label>
                  <hr className="w-28 border-gray-300 block" />
                </div>
              </div>

              <TextField
                variant="outlined"
                margin="normal"
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                autoFocus
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
                <Link href={"/next/auth/forgot"} variant="body2">
                  <p className="text-sm font-medium text-indigo-600 hover:text-indigo-500">
                    Forgot password?
                  </p>
                </Link>

                <Link href={"/next/auth/registration"} variant="body2">
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
  )
}

export default function SignIn() {
  return <Layout content={SignInPageContent()} />
}
