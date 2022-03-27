// @ts-nocheck

import React from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import Box from '@mui/material/Box'
import { styled } from '@mui/material/styles'
import Header from './Header'
import Footer from './Footer'

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
    <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
      <CssBaseline />
      <Header />
      <DrawerHeader />
      <main>
      {/*  <div className={classes.appBarSpacer} />*/}
      {/*  <Container>{children}</Container>*/}
        <Box pt={4}>
          <Footer />
        </Box>
      </main>
    </Box>
  )
}
