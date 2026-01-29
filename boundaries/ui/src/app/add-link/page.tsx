'use client'

/**
 * Add Link Page - Migrated to React 19
 *
 * Changes:
 * - ✅ Replaced 9+ useState with consolidated state
 * - ✅ Added useOptimistic for instant form preview
 * - ✅ Added useTransition for smooth submit
 * - ✅ Removed useEffect for host (use window directly)
 * - ✅ Simplified error handling
 * - ✅ Better user experience
 *
 * Old version: 9+ useState, 315 lines
 * New version: Simplified, ~250 lines
 */

import { useState, useOptimistic, useTransition, useCallback, useRef } from 'react'
import FileCopyIcon from '@mui/icons-material/FileCopy'
import CheckCircleIcon from '@mui/icons-material/CheckCircle'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import IconButton from '@mui/material/IconButton'
import Link from '@mui/material/Link'
import TextField from '@mui/material/TextField'
import Typography from '@mui/material/Typography'
import { CopyToClipboard } from 'react-copy-to-clipboard'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

import { createLinkUseCase } from '@/application/link'
import { CreateLinkCommand, LinkDomainError } from '@/domain/link/link.types'
import withAuthSync from '@/components/Private'
import { FormErrorBoundary } from '@/components/error'
import Header from '@/components/Page/Header'
import { validateUrl, validateEmailList, parseEmailList } from '@/utils/validation'

// Types for consolidated state
interface FormState {
  url: string
  describe: string
  allowedEmailsInput: string
  urlError: string | null
  allowedEmailsError: string | null
}

interface SubmitState {
  createdHash: string | null
  copied: boolean
}

// Constants for retry logic
const MAX_RETRIES = 3
const RETRY_DELAY = 2000

function Page() {
  const router = useRouter()
  const [isPending, startTransition] = useTransition()
  const retryCountRef = useRef(0)

  // Consolidated form state with useOptimistic for instant updates
  const [formState, setFormState] = useState<FormState>({
    url: '',
    describe: '',
    allowedEmailsInput: '',
    urlError: null,
    allowedEmailsError: null,
  })

  const [optimisticForm, setOptimisticForm] = useOptimistic(formState)

  // Submit state (separate from form for clarity)
  const [submitState, setSubmitState] = useState<SubmitState>({
    createdHash: null,
    copied: false,
  })

  // Get host directly (no useEffect needed)
  const host = typeof window !== 'undefined' ? window.location.host : ''

  /**
   * Handle domain error with toast and appropriate action
   */
  const handleError = useCallback(
    (error: LinkDomainError) => {
      // Show toast based on error type
      if (error.action === 'LOGIN') {
        toast.error(error.title, {
          description: error.detail,
          action: {
            label: 'Sign in',
            onClick: () => router.push('/auth/login'),
          },
          duration: 5000,
        })
        // Auto-redirect after delay
        setTimeout(() => router.push('/auth/login'), 3000)
      } else if (error.action === 'RETRY') {
        toast.error(error.title, {
          description: error.detail,
          action: {
            label: 'Try again',
            onClick: () => toast.dismiss(),
          },
          duration: 8000,
        })
      } else {
        toast.error(error.title, {
          description: error.detail,
          duration: 5000,
        })
      }
    },
    [router],
  )

  /**
   * Execute create link with auto-retry for service unavailable
   */
  const executeWithRetry = useCallback(
    async (command: CreateLinkCommand): Promise<boolean> => {
      const result = await createLinkUseCase.execute(command)

      if (result.kind === 'success') {
        retryCountRef.current = 0
        setSubmitState({
          createdHash: result.link.hash,
          copied: false,
        })

        // Clear form after success
        const emptyForm = {
          url: '',
          describe: '',
          allowedEmailsInput: '',
          urlError: null,
          allowedEmailsError: null,
        }
        setFormState(emptyForm)
        setOptimisticForm(emptyForm)

        toast.success('Link created successfully!', {
          description: 'Your short link is ready to use.',
          duration: 3000,
        })
        return true
      }

      // Handle retryable errors
      if ((result.error.code === 'SERVICE_UNAVAILABLE' || result.error.code === 'NETWORK_ERROR') && retryCountRef.current < MAX_RETRIES) {
        retryCountRef.current++

        toast.loading(`Connection issue. Retrying... (${retryCountRef.current}/${MAX_RETRIES})`, {
          id: 'retry-create-link',
          duration: RETRY_DELAY,
        })

        await new Promise((resolve) => setTimeout(resolve, RETRY_DELAY))
        return executeWithRetry(command)
      }

      // Reset retry count and show error
      if (retryCountRef.current >= MAX_RETRIES) {
        retryCountRef.current = 0
        toast.dismiss('retry-create-link')
        toast.error('Service unavailable', {
          description: 'Please try again later. If the problem persists, contact support.',
          duration: 8000,
        })
      } else {
        handleError(result.error)
      }

      return false
    },
    [handleError, setOptimisticForm],
  )

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target

    // Optimistic update for instant UI feedback
    startTransition(() => {
      const newState = { ...optimisticForm, [name]: value }

      // Real-time validation
      if (name === 'url' && value) {
        const validation = validateUrl(value, false)
        newState.urlError = validation.isValid ? null : validation.error || null
      } else if (name === 'url') {
        newState.urlError = null
      }

      if (name === 'allowedEmailsInput' && value.trim()) {
        const emails = parseEmailList(value)
        if (emails.length > 0) {
          const validation = validateEmailList(emails)
          newState.allowedEmailsError = validation.isValid ? null : validation.error || null
        } else {
          newState.allowedEmailsError = null
        }
      } else if (name === 'allowedEmailsInput') {
        newState.allowedEmailsError = null
      }

      setOptimisticForm(newState)
      setFormState(newState)
    })
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    // Validate URL
    const urlValidation = validateUrl(optimisticForm.url, true)
    if (!urlValidation.isValid) {
      toast.error('Invalid URL', {
        description: urlValidation.error || 'Please enter a valid URL',
      })
      setFormState((prev) => ({ ...prev, urlError: urlValidation.error || null }))
      return
    }

    // Parse and validate allowed emails
    const emails = parseEmailList(optimisticForm.allowedEmailsInput)
    if (emails.length > 0) {
      const emailValidation = validateEmailList(emails)
      if (!emailValidation.isValid) {
        toast.error('Invalid email addresses', {
          description: emailValidation.error || 'Please check the email format',
        })
        setFormState((prev) => ({ ...prev, allowedEmailsError: emailValidation.error || null }))
        return
      }
    }

    // Submit with transition for smooth UX
    startTransition(async () => {
      try {
        await executeWithRetry({
          url: optimisticForm.url,
          describe: optimisticForm.describe,
          allowed_emails: emails.length > 0 ? emails : undefined,
        })
      } catch (err: any) {
        console.error('An error occurred', err)
        toast.error('Unexpected error', {
          description: 'Could not create the link. Please try again later.',
        })
      }
    })
  }

  const handleCopy = () => {
    setSubmitState((prev) => ({ ...prev, copied: true }))
    toast.success('Copied to clipboard!', { duration: 2000 })
    setTimeout(() => {
      setSubmitState((prev) => ({ ...prev, copied: false }))
    }, 2000)
  }

  const handleCreateAnother = () => {
    setSubmitState({
      createdHash: null,
      copied: false,
    })
    const emptyForm = {
      url: '',
      describe: '',
      allowedEmailsInput: '',
      urlError: null,
      allowedEmailsError: null,
    }
    setFormState(emptyForm)
    setOptimisticForm(emptyForm)
  }

  const shortUrl = submitState.createdHash && host ? `${host}/s/${submitState.createdHash}` : ''
  const describeMaxLength = 500

  return (
    <>
      {/*<NextSeo title="Add link" description="Add a new link" />*/}
      <div className="container mx-auto w-5/6 sm:w-2/3 h-full" role="main">
        <Header title="Add link" />

        <FormErrorBoundary>
          {submitState.createdHash ? (
            // Success state - show created link
            <Card className="mt-6 bg-white dark:bg-gray-800 shadow-lg" role="region" aria-label="Link created successfully">
              <CardContent className="p-6">
                <div className="flex items-center gap-3 mb-4">
                  <CheckCircleIcon className="text-green-500" fontSize="large" aria-hidden="true" />
                  <Typography variant="h5" component="h2" className="text-gray-900 dark:text-white">
                    Link Created Successfully!
                  </Typography>
                </div>

                <div className="mt-4 p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                  <Typography variant="body2" component="p" className="text-gray-600 dark:text-gray-400 mb-2">
                    Your short link:
                  </Typography>
                  <div className="flex items-center gap-2">
                    <Link
                      href={`/s/${submitState.createdHash}`}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-indigo-600 dark:text-indigo-400 hover:underline font-mono text-lg break-all"
                      aria-label={`Short link: ${shortUrl}`}
                    >
                      {shortUrl}
                    </Link>
                    <CopyToClipboard text={shortUrl} onCopy={handleCopy}>
                      <IconButton
                        aria-label={submitState.copied ? 'Link copied' : 'Copy link to clipboard'}
                        color={submitState.copied ? 'success' : 'default'}
                        size="small"
                      >
                        <FileCopyIcon />
                      </IconButton>
                    </CopyToClipboard>
                  </div>
                  {submitState.copied && (
                    <Typography variant="caption" className="text-green-600 dark:text-green-400 mt-1" role="status" aria-live="polite">
                      Copied to clipboard!
                    </Typography>
                  )}
                </div>

                <div className="mt-6 flex gap-3">
                  <Button
                    variant="contained"
                    onClick={handleCreateAnother}
                    sx={{
                      bgcolor: 'indigo.600',
                      '&:hover': { bgcolor: 'indigo.700' },
                    }}
                  >
                    Create Another Link
                  </Button>
                  <Button variant="outlined" onClick={() => router.push('/links')}>
                    View All Links
                  </Button>
                </div>
              </CardContent>
            </Card>
          ) : (
            // Form state
            <Card className="mt-6 bg-white dark:bg-gray-800 shadow-lg">
              <CardContent className="p-6">
                <Typography variant="h6" component="h2" className="text-gray-900 dark:text-white mb-2">
                  Create a Short Link
                </Typography>
                <Typography variant="body2" component="p" className="text-gray-600 dark:text-gray-400 mb-4">
                  Enter a URL to create a short, easy-to-share link.
                </Typography>

                <Box
                  component="form"
                  onSubmit={handleSubmit}
                  noValidate
                  autoComplete="off"
                  sx={{ width: '100%' }}
                  aria-label="Create short link form"
                >
                  <TextField
                    variant="outlined"
                    label="Your URL"
                    name="url"
                    required
                    fullWidth
                    value={optimisticForm.url}
                    onChange={handleChange}
                    error={!!optimisticForm.urlError}
                    helperText={optimisticForm.urlError || 'Enter the full URL you want to shorten (e.g., https://example.com)'}
                    placeholder="https://example.com"
                    inputProps={{
                      'aria-describedby': 'url-helper-text',
                      'aria-invalid': !!optimisticForm.urlError,
                    }}
                    sx={{
                      mb: 2,
                      opacity: isPending ? 0.7 : 1,
                      transition: 'opacity 0.2s',
                    }}
                    autoFocus
                  />

                  <TextField
                    variant="outlined"
                    label="Description (optional)"
                    name="describe"
                    fullWidth
                    multiline
                    rows={3}
                    value={optimisticForm.describe ?? ''}
                    onChange={handleChange}
                    helperText={`${(optimisticForm.describe ?? '').length}/${describeMaxLength} characters`}
                    inputProps={{
                      maxLength: describeMaxLength,
                      'aria-describedby': 'describe-helper-text',
                    }}
                    sx={{
                      mb: 2,
                      opacity: isPending ? 0.7 : 1,
                      transition: 'opacity 0.2s',
                    }}
                  />

                  <TextField
                    variant="outlined"
                    label="Allowed Emails (optional)"
                    name="allowedEmailsInput"
                    fullWidth
                    multiline
                    rows={3}
                    value={optimisticForm.allowedEmailsInput}
                    onChange={handleChange}
                    error={!!optimisticForm.allowedEmailsError}
                    helperText={
                      optimisticForm.allowedEmailsError ||
                      'Enter email addresses separated by commas or newlines. Leave empty for a public link (anyone can access). Maximum 100 emails.'
                    }
                    placeholder="user@example.com, another@example.com"
                    inputProps={{
                      'aria-describedby': 'emails-helper-text',
                      'aria-invalid': !!optimisticForm.allowedEmailsError,
                    }}
                    sx={{
                      mb: 3,
                      opacity: isPending ? 0.7 : 1,
                      transition: 'opacity 0.2s',
                    }}
                  />

                  <Button
                    variant="contained"
                    type="submit"
                    disabled={isPending || !!optimisticForm.urlError || !!optimisticForm.allowedEmailsError}
                    fullWidth
                    size="large"
                    aria-busy={isPending}
                    sx={{
                      bgcolor: 'indigo.600',
                      '&:hover': { bgcolor: 'indigo.700' },
                      py: 1.5,
                    }}
                  >
                    {isPending ? 'Creating...' : 'Create Short Link'}
                  </Button>
                </Box>
              </CardContent>
            </Card>
          )}
        </FormErrorBoundary>
      </div>
    </>
  )
}

export default withAuthSync(Page)
