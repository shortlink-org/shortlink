// @ts-nocheck

import React from 'react'
import CssBaseline from '@mui/material/CssBaseline'
// import Container from '@mui/material/Container'
import Grid from '@mui/material/Grid'
// import Box from '@mui/material/Box'
import Header from './Header'
// import Footer from './Footer'

// @ts-ignore
export function Layout({ children }) {
  return (
    <Grid>
      <CssBaseline />
      <Header />
      {/*<main className={classes.content}>*/}
      {/*  <div className={classes.appBarSpacer} />*/}
      {/*  <Container>{children}</Container>*/}
      {/*  <Box pt={4}>*/}
      {/*    <Footer />*/}
      {/*  </Box>*/}
      {/*</main>*/}
    </Grid>
  )
}
