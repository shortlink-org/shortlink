'use client'

// @ts-nocheck
import React, { Fragment } from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import Container from '@mui/material/Container'
import Grid from '@mui/material/Grid'
import Paper from '@mui/material/Paper'
import Box from '@mui/material/Box'
import Orders from 'components/widgets/Orders'
import Profile from 'components/Dashboard/profile'
import withAuthSync from 'components/Private'
import Header from '../../../components/Page/Header'

function Dashboard() {
  return (
    <Fragment>
      {/*<NextSeo*/}
      {/*  title="Dashboard"*/}
      {/*  description="Dashboard page for your account."*/}
      {/*/>*/}

      <Box sx={{ display: 'flex' }}>
        <CssBaseline />

        <Box
          component="main"
          sx={{
            backgroundColor: (theme) =>
              theme.palette.mode === 'light'
                ? theme.palette.grey[100]
                : theme.palette.grey[900],
            flexGrow: 1,
            overflow: 'auto',
          }}
        >
          <Header title="Dashboard" />

          <Profile />

          <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
            <Grid container spacing={3}>
              {/* Recent Orders */}
              <Grid item xs={12}>
                <Paper
                  sx={{ p: 2, my: 2, display: 'flex', flexDirection: 'column' }}
                >
                  <Orders />
                </Paper>
              </Grid>
            </Grid>
          </Container>
        </Box>
      </Box>
    </Fragment>
  )
}

export default withAuthSync(Dashboard)
