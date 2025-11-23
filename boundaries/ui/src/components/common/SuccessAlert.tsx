import { Alert, AlertProps } from '@mui/material'
import React, { useEffect } from 'react'

interface SuccessAlertProps {
  message: string | null
  onClose?: () => void
  autoHideDuration?: number
  sx?: AlertProps['sx']
}

export default function SuccessAlert({ 
  message, 
  onClose, 
  autoHideDuration = 3000,
  sx 
}: SuccessAlertProps) {
  useEffect(() => {
    if (message && autoHideDuration > 0) {
      const timer = setTimeout(() => {
        onClose?.()
      }, autoHideDuration)
      return () => clearTimeout(timer)
    }
  }, [message, autoHideDuration, onClose])

  if (!message) return null

  return (
    <Alert severity="success" onClose={onClose} sx={{ mb: 2, ...sx }}>
      {message}
    </Alert>
  )
}

