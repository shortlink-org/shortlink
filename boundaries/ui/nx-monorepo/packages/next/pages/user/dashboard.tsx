// @ts-nocheck
import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import CssBaseline from '@mui/material/CssBaseline'
import Grid from '@mui/material/Grid'
import Paper from '@mui/material/Paper'
import { NextSeo } from 'next-seo'

import { Layout } from 'components'
import Profile from 'components/Dashboard/profile'
import withAuthSync from 'components/Private'
import Orders from 'components/widgets/Orders'

import Header from '../../components/Page/Header'

function Dashboard() {
  return (
    <Layout>
      <NextSeo
        title="Dashboard"
        description="Dashboard page for your account."
      />

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
    </Layout>
  )
}

export default withAuthSync(Dashboard)
