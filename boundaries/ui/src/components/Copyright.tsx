'use client'

import MuiLink from '@mui/material/Link'
import Typography from '@mui/material/Typography'
import React from 'react'

export default function Copyright() {
  const currentYear = new Date().getFullYear()

  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {'Copyright © '}
      <MuiLink color="inherit" href="/" underline="hover">
        Shortlink
      </MuiLink>{' '}
      {currentYear}.
    </Typography>
  )
}
