import React, { useEffect, useState } from 'react'
import Avatar from '@material-ui/core/Avatar'
import Button from '@material-ui/core/Button'
import CssBaseline from '@material-ui/core/CssBaseline'
import TextField from '@material-ui/core/TextField'
import Link from '@material-ui/core/Link'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import AlertUI from '@material-ui/lab/Alert'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import { Configuration, PublicApi } from '@ory/kratos-client'

const useStyles = makeStyles(theme => ({
  root: {
    height: '100vh',
  },
  image: {
    backgroundImage: 'url(https://source.unsplash.com/random)',
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'light'
        ? theme.palette.grey[50]
        : theme.palette.grey[900],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  },
  paper: {
    margin: theme.spacing(8, 4),
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
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 5),
  },
}))

function Alert(props) {
  return <AlertUI elevation={6} variant="filled" {...props} />
}

function Login() {
  const classes = useStyles()

  const [formUsernamePasswordAction, setFormUsernamePasswordAction] = useState()
  const [csrfUsernamePasswordToken, setCsrfUsernamePasswordToken] = useState()
  const [loginMessages, setLoginMessages] = useState()

  const kratos = new PublicApi(
    new Configuration({ basePath: 'http://127.0.0.1:4433' }),
  )

  useEffect(() => {
    if (!new URL(document.location).searchParams.get('flow')) {
      window.location.href = 'http://127.0.0.1:4433/self-service/login/browser'
    }
    const flowId = new URL(document.location).searchParams.get('flow')
    kratos
      .getSelfServiceLoginFlow(flowId)
      .then(({ status, data: flow }) => {
        if ([401, 403, 404].includes(status)) {
          return window.location.replace(
            'http://127.0.0.1:4433/self-service/login/browser',
          )
        }
        if (status !== 200) {
          return Promise.reject(flow)
        }
        setFormUsernamePasswordAction(
          JSON.stringify(flow.methods.password.config.action).replaceAll(
            '"',
            '',
          ),
        )
        setCsrfUsernamePasswordToken(
          JSON.stringify(
            flow.methods.password.config.fields[2].value,
          ).replaceAll('"', ''),
        )
        setLoginMessages(flow.methods.password.config.messages[0].text)
      })
      .catch(err => {
        console.log(err)
      })
  }, [loginMessages])

  return (
    <React.Fragment>
      <Grid container component="main" className={classes.root}>
        <CssBaseline />
        <Grid item xs={false} sm={4} md={7} className={classes.image} />
        <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
          <div className={classes.paper}>
            <Avatar className={classes.avatar}>
              <LockOutlinedIcon />
            </Avatar>
            <Typography component="h1" variant="h5">
              Sign in
            </Typography>
            <form
              className={classes.form}
              action={formUsernamePasswordAction}
              method="POST"
            >
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                id="identifier"
                type="text"
                label="Email Address"
                name="identifier"
                autoComplete="email"
                autoFocus
                required
              />
              <TextField
                variant="outlined"
                margin="normal"
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="current-password"
                required
              />
              <Button
                type="submit"
                fullWidth
                variant="contained"
                color="primary"
                className={classes.submit}
              >
                Sign In
              </Button>
              <TextField
                name="csrf_token"
                id="csrf_token"
                type="hidden"
                required
                fullWidth
                variant="outlined"
                label="Csrf token"
                value={csrfUsernamePasswordToken}
                className={classes.csrf}
              />
            </form>

            <Grid container>
              <Grid item xs>
                <Link href="http://127.0.0.1:3000/recovery" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="http://127.0.0.1:3000/register" variant="body2">
                  Don't have an account? Sign Up
                </Link>
              </Grid>
            </Grid>
          </div>
          {loginMessages ? (
            <React.Fragment>
              <Alert severity="error" style={{ margin: 'auto', width: '75%' }}>
                {loginMessages}
              </Alert>
            </React.Fragment>
          ) : (
            <React.Fragment />
          )}
        </Grid>
      </Grid>
    </React.Fragment>
  )
}

export default Login
