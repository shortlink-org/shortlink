// @ts-nocheck

import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import Container from '@material-ui/core/Container'
import Grid from '@material-ui/core/Grid'
import { makeStyles } from '@material-ui/core/styles'
import Box from '@material-ui/core/Box'
import Header from './Header'
import Footer from './Footer'
import {colors} from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    background: colors.grey[100]
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
export function Layout(props) {
  const classes = useStyles()

  return (
    <Grid className={classes.root}>
      <CssBaseline />
      <Header />
      <main className={classes.content}>
        <div className={classes.appBarSpacer} />
        <Container>{props.content}</Container>
        <Box pt={4}>
          <Footer />
        </Box>
      </main>
    </Grid>
  )
}
