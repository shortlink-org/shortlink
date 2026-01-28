'use client'

import { useMemo, useState } from 'react'
import Link from 'next/link'
import Box from '@mui/material/Box'
import Button from '@mui/material/Button'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import Chip from '@mui/material/Chip'
import Divider from '@mui/material/Divider'
import LinearProgress from '@mui/material/LinearProgress'
import Stack from '@mui/material/Stack'
import Step from '@mui/material/Step'
import StepContent from '@mui/material/StepContent'
import StepLabel from '@mui/material/StepLabel'
import Stepper from '@mui/material/Stepper'
import TextField from '@mui/material/TextField'
import Typography from '@mui/material/Typography'

import withAuthSync from '@/components/Private'
import Header from '@/components/Page/Header'
import { ErrorAlert, SuccessAlert } from '@/components/common'
import { validateUrl } from '@/utils/validation'

type StepStatus = 'pending' | 'active' | 'done' | 'error'

type LogEntry = {
  time: string
  message: string
  type: 'info' | 'success' | 'error'
}

const steps = [
  {
    id: 'validate',
    title: 'Validate sitemap URL',
    description: 'Check that the URL is reachable and formatted correctly.',
  },
  {
    id: 'enqueue',
    title: 'Send parse request',
    description: 'Submit the sitemap to the service for processing.',
  },
  {
    id: 'parse',
    title: 'Parse sitemap',
    description: 'The service fetches and extracts links from the sitemap.',
  },
  {
    id: 'save',
    title: 'Save links & metadata',
    description: 'Links are stored, metadata enrichment starts asynchronously.',
  },
]

const formatTime = (date: Date) =>
  date.toLocaleTimeString('ru-RU', { hour: '2-digit', minute: '2-digit', second: '2-digit' })

function Page() {
  const [url, setUrl] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [logs, setLogs] = useState<LogEntry[]>([])
  const [statuses, setStatuses] = useState<StepStatus[]>(() => steps.map(() => 'pending'))

  const activeStep = useMemo(() => {
    const current = statuses.findIndex(status => status === 'active')
    return current === -1 ? steps.length : current
  }, [statuses])

  const resetProcess = () => {
    setError(null)
    setSuccess(null)
    setLogs([])
    setStatuses(steps.map(() => 'pending'))
  }

  const pushLog = (message: string, type: LogEntry['type'] = 'info') => {
    setLogs(prev => [
      ...prev,
      {
        time: formatTime(new Date()),
        message,
        type,
      },
    ])
  }

  const setStepStatus = (index: number, status: StepStatus) => {
    setStatuses(prev => prev.map((value, idx) => (idx === index ? status : value)))
  }

  const markStepsAfter = (index: number, status: StepStatus) => {
    setStatuses(prev => prev.map((value, idx) => (idx > index ? status : value)))
  }

  const normalizeUrl = (input: string) => {
    if (!input) return input
    return input.startsWith('http://') || input.startsWith('https://') ? input : `https://${input}`
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()
    setError(null)
    setSuccess(null)
    setLogs([])
    setIsSubmitting(true)
    setStatuses(steps.map(() => 'pending'))

    setStepStatus(0, 'active')
    pushLog('Validating sitemap URL...')

    const validation = validateUrl(url, true)
    if (!validation.isValid) {
      setStepStatus(0, 'error')
      setError(validation.error || 'Invalid URL')
      pushLog(validation.error || 'Invalid URL', 'error')
      setIsSubmitting(false)
      return
    }

    const normalizedUrl = normalizeUrl(url)
    setStepStatus(0, 'done')
    setStepStatus(1, 'active')
    markStepsAfter(1, 'pending')
    pushLog(`URL validated: ${normalizedUrl}`, 'success')
    pushLog('Sending parse request to /api/sitemap...')

    try {
      const response = await fetch('/api/sitemap', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify({ url: normalizedUrl }),
      })

      if (!response.ok) {
        const payload = await response.json().catch(() => null)
        const message = payload?.error || 'Failed to start sitemap parsing'
        setStepStatus(1, 'error')
        setError(message)
        pushLog(message, 'error')
        setIsSubmitting(false)
        return
      }

      setStepStatus(1, 'done')
      setStepStatus(2, 'active')
      markStepsAfter(2, 'pending')
      pushLog('Request accepted. Parsing sitemap in background.', 'success')

      setSuccess('Sitemap accepted. Processing continues asynchronously.')
      setIsSubmitting(false)
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Network error'
      setStepStatus(1, 'error')
      setError(message)
      pushLog(message, 'error')
      setIsSubmitting(false)
    }
  }

  return (
    <Box className="container mx-auto px-6 pb-10">
      <Header title="Sitemap Parser" />

      <Stack spacing={3}>
        <Card variant="outlined">
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Запуск парсинга sitemap
            </Typography>
            <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
              Введите URL sitemap. Сервис обработает ссылки и запустит сохранение в фоне.
            </Typography>

            <ErrorAlert error={error} onClose={() => setError(null)} />
            <SuccessAlert message={success} onClose={() => setSuccess(null)} />

            <Box component="form" onSubmit={handleSubmit} sx={{ display: 'flex', gap: 2, flexWrap: 'wrap' }}>
              <TextField
                label="Sitemap URL"
                name="url"
                value={url}
                onChange={event => setUrl(event.target.value)}
                placeholder="https://example.com/sitemap.xml"
                fullWidth
                required
              />
              <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1} sx={{ width: { xs: '100%', sm: 'auto' } }}>
                <Button type="submit" variant="contained" disabled={isSubmitting}>
                  {isSubmitting ? 'Запуск...' : 'Запустить парсинг'}
                </Button>
                <Button type="button" variant="outlined" onClick={resetProcess} disabled={isSubmitting}>
                  Сбросить
                </Button>
              </Stack>
            </Box>
          </CardContent>
        </Card>

        <Card variant="outlined">
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Прогресс обработки
            </Typography>
            <Stepper activeStep={activeStep} orientation="vertical">
              {steps.map((step, index) => {
                const status = statuses[index]
                const isError = status === 'error'
                const isDone = status === 'done'

                return (
                  <Step key={step.id} completed={isDone} active={status === 'active'}>
                    <StepLabel error={isError}>{step.title}</StepLabel>
                    <StepContent>
                      <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
                        {step.description}
                      </Typography>
                      {status === 'active' && <LinearProgress sx={{ mb: 1 }} />}
                      {status === 'done' && <Chip label="Готово" size="small" color="success" />}
                      {status === 'error' && <Chip label="Ошибка" size="small" color="error" />}
                      {status === 'pending' && <Chip label="Ожидание" size="small" variant="outlined" />}
                    </StepContent>
                  </Step>
                )
              })}
            </Stepper>

            <Divider sx={{ my: 3 }} />

            <Typography variant="subtitle1" gutterBottom>
              Логи запуска
            </Typography>
            {logs.length === 0 ? (
              <Typography variant="body2" color="text.secondary">
                Здесь появится история действий после запуска.
              </Typography>
            ) : (
              <Stack spacing={1}>
                {logs.map((log, index) => (
                  <Box
                    key={`${log.time}-${index}`}
                    sx={{
                      display: 'flex',
                      gap: 2,
                      alignItems: 'flex-start',
                      p: 1.5,
                      borderRadius: 1.5,
                      bgcolor: log.type === 'error' ? 'error.light' : 'background.default',
                    }}
                  >
                    <Typography variant="caption" color="text.secondary" sx={{ minWidth: 72 }}>
                      {log.time}
                    </Typography>
                    <Typography variant="body2" color={log.type === 'error' ? 'error.contrastText' : 'text.primary'}>
                      {log.message}
                    </Typography>
                  </Box>
                ))}
              </Stack>
            )}

            <Box sx={{ mt: 3 }}>
              <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>
                После запуска обработка продолжается в фоне. Проверьте список ссылок, чтобы увидеть результат.
              </Typography>
              <Button component={Link} href="/admin/links" variant="text">
                Перейти к списку ссылок
              </Button>
            </Box>
          </CardContent>
        </Card>
      </Stack>
    </Box>
  )
}

export default withAuthSync(Page)
