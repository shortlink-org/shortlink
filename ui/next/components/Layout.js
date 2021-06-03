import React from 'react';
import Container from '@material-ui/core/Container'
import Grid from '@material-ui/core/Grid'
import {makeStyles} from "@material-ui/core/styles"
import Box from '@material-ui/core/Box'
import Header from './Header';
import Copyright from './Copyright';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex'
  },
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
    gridTemplateRows: '64px auto 72px',
    display: 'grid',
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  appBarSpacer: theme.mixins.toolbar,
}));

export function Layout(props) {
  const classes = useStyles()

  return (
    (
      <Grid className={classes.root}>
        <Header />
        <main className={classes.content}>
          <div className={classes.appBarSpacer} />
          <Container>
            {props.content}
          </Container>
          <Box pt={4}>
            <Copyright />
          </Box>
        </main>
      </Grid>
    )
  )
}

