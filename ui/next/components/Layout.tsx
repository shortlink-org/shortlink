// @ts-nocheck

import React from 'react'
import CssBaseline from '@mui/material/CssBaseline'
import Container from '@mui/material/Container'
import Grid from '@mui/material/Grid'
import { makeStyles } from '@mui/styles'
import Box from '@mui/material/Box'
import Header from './Header'
import Footer from './Footer'
import { colors } from '@mui/material'

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    background: colors.grey[100],
  },
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
    gridTemplateRows: '64px 1fr auto',
    display: 'grid',
    paddingTop: theme.spacing(4),
  },
  appBarSpacer: theme.mixins.toolbar,
}))

// @ts-ignore
export function Layout({ children }) {
  const classes = useStyles()

  return (
    <Grid className={classes.root}>
      <CssBaseline />
      <Header />
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container>{children}</Container>
        <Box pt={4}>
          <Footer />
        </Box>
      </main>
    </Grid>
  )
}
