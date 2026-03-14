'use client'

import Paper from '@mui/material/Paper'

import Profile from '@/components/Dashboard/profile'
import withAuthSync from '@/components/Private'
import Orders from '@/components/widgets/Orders'
import Header from '@/components/Page/Header'
import PageSection from '@/components/Page/Section'

function Page() {
  return (
    <div className={'flex'}>
      {/*<NextSeo title="Dashboard" description="Dashboard page for your account." />*/}

      <main className={'flex-1'}>
        <Header title="Dashboard" description="Get a quick overview of your workspace activity and recent orders." />

        <PageSection className="space-y-6 pb-10">
          <Profile />

          <div className="flex flex-col md:flex-row">
            <div className="w-full">
              <Paper sx={{ p: 2, display: 'flex', flexDirection: 'column' }}>
                <Orders />
              </Paper>
            </div>
          </div>
        </PageSection>
      </main>
    </div>
  )
}

export default withAuthSync(Page)
