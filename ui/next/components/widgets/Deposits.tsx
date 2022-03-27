// @ts-nocheck

import React from 'react'
import Link from '@mui/material/Link'
import { styled } from '@mui/styles'
import Typography from '@mui/material/Typography'
import Title from './Title'

function preventDefault(event) {
  event.preventDefault()
}

const useStyles = styled({
  depositContext: {
    flex: 1,
  },
})

export default function Deposits() {
  const classes = {}
  return (
    <React.Fragment>
      <Title>Recent Deposits</Title>
      <Typography component="p" variant="h4">
        $3,024.00
      </Typography>
      <Typography color="textSecondary" className={classes.depositContext}>
        on 15 March, 2019
      </Typography>
      <div>
        <Link
          color="primary"
          href="#"
          onClick={preventDefault}
          underline="hover"
        >
          View balance
        </Link>
      </div>
    </React.Fragment>
  )
}
