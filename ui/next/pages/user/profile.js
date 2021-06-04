import React from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import { Layout } from '../../components';

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    '& > *': {
      margin: theme.spacing(1),
      width: theme.spacing(16),
      height: theme.spacing(16),
    },
  },
}))

export function ProfileContent() {
  const classes = useStyles()

  return (
    <Grid container direction="column" justify="space-around" alignItems="center" className={classes.root}>
      <Paper />
    </Grid>
  )
}

export default function Profile() {
  return <Layout content={ProfileContent()} />;
}
