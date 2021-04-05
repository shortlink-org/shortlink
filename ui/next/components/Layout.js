import React from 'react';
import Container from '@material-ui/core/Container'
import Grid from '@material-ui/core/Grid'
import Header from './Header';
import Copyright from './Copyright';
import {makeStyles} from "@material-ui/core/styles";

const useStyles = makeStyles((theme) => ({
  root: {
    height: '100vh',
    display: 'grid',
    gridTemplateRows: 'auto 1fr auto',
  },
}));

export function Layout(props) {
  const classes = useStyles()

  return (
    (
      <Grid className={classes.root}>
        <Header />
        <Container >
          {props.content}
        </Container>
        <Copyright />
      </Grid>
    )
  )
}

