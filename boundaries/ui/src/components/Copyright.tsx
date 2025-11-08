'use client'

import MuiLink from '@mui/material/Link'
import Typography from '@mui/material/Typography'
import React, { useEffect, useState } from 'react'

export default function Copyright() {
  const [currentYear, setCurrentYear] = useState<number | null>(null)

  useEffect(() => {
    setCurrentYear(new Date().getFullYear())
  }, [])

  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {'Copyright © '}
      <MuiLink color="inherit" href="/next" underline="hover">
        Shortlink
      </MuiLink>{' '}
      {currentYear ?? '...'}.
    </Typography>
  )
}
