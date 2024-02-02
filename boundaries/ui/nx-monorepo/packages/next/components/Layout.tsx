// @ts-nocheck

import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import CssBaseline from '@mui/material/CssBaseline'
import { styled } from '@mui/material/styles'
import * as React from 'react'
import {
  ScrollToTopButton, // @ts-ignore
} from '@shortlink-org/ui-kit'

import PushNotificationLayout from 'components/PushNotificationLayout'

import Footer from './Footer'
import Header from './Header'

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

        <Box
          component="main"
          sx={{ flexGrow: 1, p: 3, gridTemplateRows: 'auto 1fr' }}
        >
          <DrawerHeader />
          <Box
            pt={4}
            sx={{
              display: 'grid',
              height: '100%',
              gridTemplateRows: '1fr auto',
            }}
          >
            <Container>{children}</Container>
            <Footer />

            <ScrollToTopButton />
          </Box>
        </Box>
      </Box>
    </PushNotificationLayout>
  )
}
