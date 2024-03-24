import Box from '@mui/material/Box'
import CssBaseline from '@mui/material/CssBaseline'
import { styled } from '@mui/material/styles'
import * as React from 'react'
// @ts-ignore
import { ScrollToTopButton } from '@shortlink-org/ui-kit'

import PushNotificationLayout from 'components/PushNotificationLayout'
import Footer from 'components/Footer'
import Header from 'components/Header'

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),
  ...theme.mixins.toolbar,
}))

// @ts-ignore
export function Layout({ children }) {
  return (
    <PushNotificationLayout>
      <Box sx={{ display: 'flex', height: '100vh' }}>
        <CssBaseline />
        <Header />

        <Box component="main" sx={{ flexGrow: 1, p: 3, gridTemplateRows: 'auto 1fr' }}>
          <DrawerHeader />
          <Box
            pt={4}
            sx={{
              display: 'grid',
              height: '100%',
              gridTemplateRows: '1fr auto',
            }}
          >
            <div className="content-center max-w-7xl m-auto">{children}</div>

            <Footer className="content-center max-w-7xl m-auto" />

            <ScrollToTopButton />
          </Box>
        </Box>
      </Box>
    </PushNotificationLayout>
  )
}
