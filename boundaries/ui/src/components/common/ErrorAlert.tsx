import { Alert, AlertProps } from '@mui/material'
import React from 'react'

interface ErrorAlertProps {
  error: string | null
  onClose?: () => void
  severity?: AlertProps['severity']
  sx?: AlertProps['sx']
}

export default function ErrorAlert({ error, onClose, severity = 'error', sx }: ErrorAlertProps) {
  if (!error) return null

  return (
    <Alert severity={severity} onClose={onClose} sx={{ mb: 2, ...sx }}>
      {error}
    </Alert>
  )
}

