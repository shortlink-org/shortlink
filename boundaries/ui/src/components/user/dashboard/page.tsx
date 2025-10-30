'use client'

import Paper from '@mui/material/Paper'

import Profile from '@/components/Dashboard/profile'
import withAuthSync from '@/components/Private'
import Orders from '@/components/widgets/Orders'
import Header from '@/components/Page/Header'

function Page() {
  return (
    <div className={'flex'}>
      {/*<NextSeo title="Dashboard" description="Dashboard page for your account." />*/}

      <main className={'flex-1'}>
        <Header title="Dashboard" />

        <Profile />

        <div className={'flex flex-col md:flex-row container'}>
          <div className="w-full">
            {/* Recent Orders */}
            <Paper sx={{ p: 2, my: 2, display: 'flex', flexDirection: 'column' }}>
              <Orders />
            </Paper>
          </div>
        </div>
      </main>
    </div>
  )
}

export default withAuthSync(Page)
