'use client'

import FileCopyIcon from '@mui/icons-material/FileCopy'
import CheckCircleIcon from '@mui/icons-material/CheckCircle'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import Grid from '@mui/material/Grid'
import IconButton from '@mui/material/IconButton'
import Link from '@mui/material/Link'
import TextField from '@mui/material/TextField'
import Typography from '@mui/material/Typography'
import React, { useState, useEffect } from 'react'
import { CopyToClipboard } from 'react-copy-to-clipboard'
import { useRouter } from 'next/navigation'

import { createLinkUseCase } from '@/application/link'
import { CreateLinkCommand } from '@/domain/link/link.types'
import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import { ErrorAlert, SuccessAlert } from '@/components/common'
import { validateUrl, validateEmail, validateEmailList, parseEmailList } from '@/utils/validation'

function Page() {
  const router = useRouter()
  const [form, setForm] = useState<CreateLinkCommand>({
    url: '',
    describe: '',
    allowed_emails: [],
  })

  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)
  const [createdHash, setCreatedHash] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)
  const [host, setHost] = useState<string>('')
  const [urlError, setUrlError] = useState<string | null>(null)
  const [copied, setCopied] = useState(false)
  const [allowedEmailsInput, setAllowedEmailsInput] = useState<string>('')
  const [allowedEmailsError, setAllowedEmailsError] = useState<string | null>(null)

  useEffect(() => {
    if (typeof window !== 'undefined') {
      setHost(window.location.host)
    }
  }, [])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setForm({ ...form, [name]: value })
    setError(null)
    setUrlError(null)

    // Real-time validation for URL
    if (name === 'url' && value) {
      const validation = validateUrl(value, false)
      if (!validation.isValid) {
        setUrlError(validation.error || null)
      }
    }
  }

  const handleAllowedEmailsChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const value = e.target.value
    setAllowedEmailsInput(value)
    setAllowedEmailsError(null)

    if (value.trim()) {
      const emails = parseEmailList(value)
      if (emails.length > 0) {
        const validation = validateEmailList(emails)
        if (!validation.isValid) {
          setAllowedEmailsError(validation.error || null)
        }
      }
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError(null)
    setSuccess(null)
    setCreatedHash(null)
    setCopied(false)

    // Validate URL
    const urlValidation = validateUrl(form.url, true)
    if (!urlValidation.isValid) {
      setError(urlValidation.error || 'Invalid URL')
      setUrlError(urlValidation.error || null)
      return
    }

    // Parse and validate allowed emails
    const emails = parseEmailList(allowedEmailsInput)
    if (emails.length > 0) {
      const emailValidation = validateEmailList(emails)
      if (!emailValidation.isValid) {
        setError(emailValidation.error || 'Invalid email addresses')
        setAllowedEmailsError(emailValidation.error || null)
        return
      }
    }

    setLoading(true)
    try {
      const result = await createLinkUseCase.execute({
        ...form,
        allowed_emails: emails.length > 0 ? emails : undefined,
      })

      if (result.kind === 'success') {
        setCreatedHash(result.link.hash)
        setSuccess('Link created successfully!')
        // Clear form after successful creation
        setForm({ url: '', describe: '', allowed_emails: [] })
        setAllowedEmailsInput('')
      } else {
        const errorMessage = result.error.detail || 'Failed to create link'
        setError(errorMessage)
        
        // Handle special actions from error
        if (result.error.action === 'LOGIN') {
          router.push('/auth/login')
        }
      }
    } catch (err: any) {
      console.error('An error occurred', err)
      setError(err.message || 'Could not create the link. Please try again later.')
    } finally {
      setLoading(false)
    }
  }

  const handleCopy = () => {
    setCopied(true)
    setSuccess('Link copied to clipboard!')
    setTimeout(() => setCopied(false), 2000)
  }

  const handleCreateAnother = () => {
    setCreatedHash(null)
    setSuccess(null)
    setForm({ url: '', describe: '', allowed_emails: [] })
    setError(null)
    setUrlError(null)
    setAllowedEmailsInput('')
    setAllowedEmailsError(null)
  }

  const shortUrl = createdHash && host ? `${host}/s/${createdHash}` : ''
  const describeMaxLength = 500

  return (
    <>
      {/*<NextSeo title="Add link" description="Add a new link" />*/}
      <div className="container mx-auto w-5/6 sm:w-2/3 h-full">
        <Header title="Add link" />

        {createdHash ? (
          // Success state - show created link
          <Card className="mt-6 bg-white dark:bg-gray-800 shadow-lg">
            <CardContent className="p-6">
              <div className="flex items-center gap-3 mb-4">
                <CheckCircleIcon className="text-green-500" fontSize="large" />
                <Typography variant="h5" className="text-gray-900 dark:text-white">
                  Link Created Successfully!
                </Typography>
              </div>

              <ErrorAlert error={error} />
              <SuccessAlert message={success} onClose={() => setSuccess(null)} />

              <div className="mt-4 p-4 bg-gray-50 dark:bg-gray-700 rounded-lg">
                <Typography variant="body2" className="text-gray-600 dark:text-gray-400 mb-2">
                  Your short link:
                </Typography>
                <div className="flex items-center gap-2">
                  <Link 
                    href={`/s/${createdHash}`} 
                    target="_blank" 
                    rel="noopener" 
                    className="text-indigo-600 dark:text-indigo-400 hover:underline font-mono text-lg break-all"
                  >
                    {shortUrl}
                  </Link>
                  <CopyToClipboard text={shortUrl} onCopy={handleCopy}>
                    <IconButton 
                      aria-label="copy link" 
                      color={copied ? 'success' : 'default'}
                      size="small"
                      title="Copy to clipboard"
                    >
                      <FileCopyIcon />
                    </IconButton>
                  </CopyToClipboard>
                </div>
                {copied && (
                  <Typography variant="caption" className="text-green-600 dark:text-green-400 mt-1">
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
                <Button
                  variant="outlined"
                  onClick={() => router.push('/links')}
                >
                  View All Links
                </Button>
              </div>
            </CardContent>
          </Card>
        ) : (
          // Form state
          <Card className="mt-6 bg-white dark:bg-gray-800 shadow-lg">
            <CardContent className="p-6">
              <Typography variant="h6" className="text-gray-900 dark:text-white mb-2">
                Create a Short Link
              </Typography>
              <Typography variant="body2" className="text-gray-600 dark:text-gray-400 mb-4">
                Enter a URL to create a short, easy-to-share link.
              </Typography>

              <Box
                component="form"
                onSubmit={handleSubmit}
                noValidate
                autoComplete="off"
                sx={{ width: '100%' }}
              >
                <ErrorAlert error={error} />
                <SuccessAlert message={success} onClose={() => setSuccess(null)} />

                <TextField
                  variant="outlined"
                  label="Your URL"
                  name="url"
                  required
                  fullWidth
                  value={form.url}
                  onChange={handleChange}
                  error={!!urlError}
                  helperText={urlError || 'Enter the full URL you want to shorten (e.g., https://example.com)'}
                  placeholder="https://example.com"
                  sx={{ mb: 2 }}
                  autoFocus
                />

                <TextField
                  variant="outlined"
                  label="Description (optional)"
                  name="describe"
                  fullWidth
                  multiline
                  rows={3}
                  value={form.describe ?? ''}
                  onChange={handleChange}
                  helperText={`${(form.describe ?? '').length}/${describeMaxLength} characters`}
                  inputProps={{ maxLength: describeMaxLength }}
                  sx={{ mb: 2 }}
                />

                <TextField
                  variant="outlined"
                  label="Allowed Emails (optional)"
                  name="allowed_emails"
                  fullWidth
                  multiline
                  rows={3}
                  value={allowedEmailsInput}
                  onChange={handleAllowedEmailsChange}
                  error={!!allowedEmailsError}
                  helperText={
                    allowedEmailsError ||
                    'Enter email addresses separated by commas or newlines. Leave empty for a public link (anyone can access). Maximum 100 emails.'
                  }
                  placeholder="user@example.com, another@example.com"
                  sx={{ mb: 3 }}
                />

                <Button
                  variant="contained"
                  type="submit"
                  disabled={loading || !!urlError || !!allowedEmailsError}
                  fullWidth
                  size="large"
                  sx={{
                    bgcolor: 'indigo.600',
                    '&:hover': { bgcolor: 'indigo.700' },
                    py: 1.5,
                  }}
                >
                  {loading ? 'Creating...' : 'Create Short Link'}
                </Button>
              </Box>
            </CardContent>
          </Card>
        )}
      </div>
    </>
  )
}

export default withAuthSync(Page)
