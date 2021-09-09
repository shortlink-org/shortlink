// @ts-nocheck
import React from 'react'
import clsx from 'clsx'
import { makeStyles } from '@material-ui/core/styles'
import CssBaseline from '@material-ui/core/CssBaseline'
import Container from '@material-ui/core/Container'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import { Layout } from 'components'
import Chart from 'components/widgets/Chart'
import Deposits from 'components/widgets/Deposits'
import Orders from 'components/widgets/Orders'
import Profile from 'components/Dashboard/profile'
import withAuthSync from 'components/Private'

const useStyles = makeStyles((theme) => ({
  title: {
    flexGrow: 1,
  },
  content: {
    flexGrow: 1,
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
}))

function Dashboard() {
  const classes = useStyles()
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight)

  return (
    <Layout>
      <div className={classes.root}>
        <CssBaseline />

        <main className={classes.content}>
          <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={3}>
              <Profile />

              <Grid item xs={12} md={12} lg={9}>
                <div className="bg-white p-6 shadow-lg rounded-lg flex justify-between items-center">
                  <div className="flex">
                    <div className="mr-4">
                      <img
                        className="shadow sm:w-12 sm:h-12 w-14 h-14 rounded-full"
                        src="http://tailwindtemplates.io/wp-content/uploads/2019/03/link.jpg"
                        alt="Avatar"
                      />
                    </div>
                    <div>
                      <h1 className="text-xl font-medium text-gray-700">Link</h1>
                      <p className="text-gray-600">UX Designer at Hyrule</p>
                    </div>
                  </div>
                  <button className="bg-blue-500 hover:opacity-75 text-white rounded-full px-8 py-2">
                    Follow
                  </button>
                </div>
              </Grid>

              {/* Chart */}
              <Grid item xs={12} md={8} lg={9}>
                <Paper className={fixedHeightPaper}>
                  <Chart />
                </Paper>
              </Grid>

              {/* Recent Deposits */}
              <Grid item xs={12} md={4} lg={3}>
                <Paper className={fixedHeightPaper}>
                  <Deposits />
                </Paper>
              </Grid>
              {/* Recent Orders */}
              <Grid item xs={12}>
                <Paper className={classes.paper}>
                  <Orders />
                </Paper>
              </Grid>
            </Grid>
          </Container>
        </main>
      </div>
    </Layout>
  )
}

export default withAuthSync(Dashboard)
