import React, { useEffect, useState } from 'react'
import Avatar from '@material-ui/core/Avatar'
import Button from '@material-ui/core/Button'
import CssBaseline from '@material-ui/core/CssBaseline'
import TextField from '@material-ui/core/TextField'
import Link from '@material-ui/core/Link'
import Grid from '@material-ui/core/Grid'
import MuiAlert from '@material-ui/lab/Alert'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import { Configuration, PublicApi } from '@ory/kratos-client'

const useStyles = makeStyles(theme => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  csrf: {
    visibility: 'hidden',
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(3),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}))

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />
}

function Registration() {
  const classes = useStyles()

  const [formAction, setFormAction] = useState()
  const [firstName, setFirstName] = useState()
  const [lastName, setLastName] = useState()
  const [email, setEmail] = useState()
  const [password, setPassword] = useState()
  const [csrfToken, setCsrfToken] = useState()
  const [registerMessages, setRegisterMessages] = useState(undefined)

  const kratos = new PublicApi(
    new Configuration({ basePath: 'http://127.0.0.1:4433' }),
  )

  useEffect(() => {
    if (
      !new URL(document.location).searchParams.get('flow') &&
      new URL(document.location).href.indexOf('registration') !== -1
    ) {
      window.location.href =
        'http://127.0.0.1:4433/self-service/registration/browser'
    }
    const flowId = new URL(document.location).searchParams.get('flow')
    kratos
      .getSelfServiceRegistrationFlow(flowId)
      .then(({ status, data: flow }) => {
        if (status === 404 || status === 410 || status === 403) {
          return window.location.replace(
            'http://127.0.0.1:4433/self-service/registration/browser',
          )
        }
        if (status !== 200) {
          return Promise.reject(flow)
        }
        setFormAction(
          JSON.stringify(flow.methods.password.config.action).replaceAll(
            '"',
            '',
          ),
        )
        setCsrfToken(
          JSON.stringify(
            flow.methods.password.config.fields[0].value,
          ).replaceAll('"', ''),
        )
        setRegisterMessages(
          flow.methods.password.config.fields[1].messages[0].text,
        )
      })
      .catch(err => {
        console.log(err)
      })
  }, [csrfToken])

  return (
    <React.Fragment>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <div className={classes.paper}>
          <Avatar className={classes.avatar}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign up
          </Typography>
          <form
            className={classes.form}
            action={formAction}
            method="POST"
            noValidate
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
            <Grid container spacing={2}>
              <Grid item xs={12} sm={6}>
                <TextField
                  autoComplete="fname"
                  name="traits.name.first"
                  variant="outlined"
                  required
                  fullWidth
                  id="traits.name.first"
                  label="First Name"
                  autoFocus
                  onChange={event => {
                    setFirstName(event.target.value)
                  }}
                />
              </Grid>
              <Grid item xs={12} sm={6}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="traits.name.last"
                  label="Last Name"
                  name="traits.name.last"
                  autoComplete="lname"
                  onChange={event => {
                    setLastName(event.target.value)
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  id="traits.email"
                  label="Email Address"
                  type="email"
                  name="traits.email"
                  autoComplete="email"
                  onChange={event => {
                    setEmail(event.target.value)
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  variant="outlined"
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type="password"
                  id="password"
                  autoComplete="current-password"
                  onChange={event => {
                    setPassword(event.target.value)
                  }}
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              className={classes.submit}
            >
              Sign Up
            </Button>
            <Grid container justify="flex-end">
              <Grid item>
                <Link
                  href="http://127.0.0.1:3000/next/auth/login"
                  variant="body2"
                >
                  Already have an account? Sign in
                </Link>
              </Grid>
            </Grid>
          </form>
        </div>
        {registerMessages ? (
          <React.Fragment>
            <Alert severity="error" style={{ marginTop: '5%' }}>
              {registerMessages}
            </Alert>
          </React.Fragment>
        ) : (
          <React.Fragment />
        )}
      </Container>
    </React.Fragment>
  )
}

export default Registration
