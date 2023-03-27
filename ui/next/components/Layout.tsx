// @ts-nocheck

import React from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import Box from '@mui/material/Box'
import Container from '@mui/material/Container'
import { styled } from '@mui/material/styles'
import Header from './Header'
import Footer from './Footer'
import PushNotificationLayout from 'components/PushNotificationLayout'

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'flex-end',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
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
            sx={{ display: 'grid', height: '100%', gridTemplateRows: '1fr auto' }}
          >
            <Container>{children}</Container>
            <Footer />
          </Box>
        </Box>
      </Box>
    </PushNotificationLayout>
  )
}
