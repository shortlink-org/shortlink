'use client'

import { AxiosError } from 'axios'
import { Session } from '@ory/client'
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { CircularProgress, Alert, Box } from '@mui/material'

import ory from '@/pkg/sdk'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import Notifications from '@/components/Profile/Notifications'
import Personal from '@/components/Profile/Personal'
import Profile from '@/components/Profile/Profile'
import Welcome from '@/components/Profile/Welcome'

function ProfileContent() {
  const router = useRouter()

  const [session, setSession] = useState<Session | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    setLoading(true)
    setError(null)
    
    ory
      .toSession()
      .then(({ data }) => {
        setSession(data)
        setLoading(false)
      })
      .catch((err: AxiosError) => {
        setLoading(false)
        
        switch (err.response?.status) {
          case 403:
          case 422:
            // Session needs second factor authentication
            return router.push('/login?aal=aal2')
          case 401:
            // User is not logged in
            setError('Please sign in to view your profile')
            return
          default:
            setError(err.message || 'Failed to load profile')
        }
      })
  }, [router])

  if (loading) {
    return (
      <>
        <Header title="Profile" />
        <Box display="flex" justifyContent="center" alignItems="center" minHeight="400px">
          <CircularProgress />
        </Box>
      </>
    )
  }

  if (error) {
    return (
      <>
        <Header title="Profile" />
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      </>
    )
  }

  if (!session) {
    return null
  }

  const firstName = session.identity?.traits?.name?.first || 'User'
  const lastName = session.identity?.traits?.name?.last || ''
  const email = session.identity?.traits?.email || ''

  return (
    <>
      <Header title="Profile" />

      <Welcome nickname={firstName} />

      <Profile />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200 dark:border-gray-700" />
        </div>
      </div>

      <Personal 
        session={session}
        firstName={firstName}
        lastName={lastName}
        email={email}
      />

      <div className="hidden sm:block" aria-hidden="true">
        <div className="py-5">
          <div className="border-t border-gray-200 dark:border-gray-700" />
        </div>
      </div>

      <Notifications />
    </>
  )
}

export default withAuthSync(ProfileContent)
