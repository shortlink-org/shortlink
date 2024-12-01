'use client'

import Grid from '@mui/material/Grid'
import Paper from '@mui/material/Paper'

import Profile from 'components/Dashboard/profile'
import withAuthSync from 'components/Private'
import Orders from 'components/widgets/Orders'
import Header from 'components/Page/Header'

function Page() {
  return (
    <div className={'flex'}>
      {/*<NextSeo title="Dashboard" description="Dashboard page for your account." />*/}

      <main className={'flex-1'}>
        <Header title="Dashboard" />

        <Profile />

        <div className={'flex flex-col md:flex-row container'}>
          <Grid container spacing={3}>
            {/* Recent Orders */}
            <Grid item xs={12}>
              <Paper sx={{ p: 2, my: 2, display: 'flex', flexDirection: 'column' }}>
                <Orders />
              </Paper>
            </Grid>
          </Grid>
        </div>
      </main>
    </div>
  )
}

export default withAuthSync(Page)
