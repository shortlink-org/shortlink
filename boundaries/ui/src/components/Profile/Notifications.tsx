import * as React from 'react'
import { useState } from 'react'
import { toast } from 'sonner'
import { Button, CircularProgress } from '@mui/material'

export default function Notifications() {
  const [emailLinkCreated, setEmailLinkCreated] = useState(true)
  const [emailLinkClicked, setEmailLinkClicked] = useState(false)
  const [emailWeeklyReport, setEmailWeeklyReport] = useState(true)
  const [emailSecurityAlerts, setEmailSecurityAlerts] = useState(true)
  const [pushNotifications, setPushNotifications] = useState('same')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)

    try {
      // TODO: Integrate with API to save notification preferences
      // For now, just simulate success
      await new Promise((resolve) => setTimeout(resolve, 1000))
      toast.success('Preferences saved', {
        description: 'Your notification preferences have been updated.',
      })
    } catch (err: any) {
      toast.error('Failed to save preferences', {
        description: err.message || 'Please try again later.',
      })
    } finally {
      setLoading(false)
    }
  }

  const CheckboxItem = ({
    id,
    name,
    label,
    description,
    checked,
    onChange,
  }: {
    id: string
    name: string
    label: string
    description: string
    checked: boolean
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  }) => (
    <label
      htmlFor={id}
      className="group flex items-start gap-4 p-4 rounded-xl cursor-pointer hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors -mx-4"
    >
      <div className="flex items-center h-6 mt-0.5">
        <input
          id={id}
          name={name}
          type="checkbox"
          checked={checked}
          onChange={onChange}
          className="h-5 w-5 rounded-md border-2 border-gray-300 dark:border-gray-600 text-indigo-600 focus:ring-indigo-500 focus:ring-offset-0 transition-all cursor-pointer"
        />
      </div>
      <div className="flex-1">
        <span className="block text-sm font-semibold text-gray-900 dark:text-white group-hover:text-indigo-600 dark:group-hover:text-indigo-400 transition-colors">
          {label}
        </span>
        <span className="block mt-1 text-sm text-gray-500 dark:text-gray-400">{description}</span>
      </div>
    </label>
  )

  const RadioItem = ({
    id,
    name,
    value,
    label,
    checked,
    onChange,
  }: {
    id: string
    name: string
    value: string
    label: string
    checked: boolean
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  }) => (
    <label
      htmlFor={id}
      className={`group flex items-center gap-4 p-4 rounded-xl cursor-pointer transition-all -mx-4 ${
        checked
          ? 'bg-indigo-50 dark:bg-indigo-900/20 ring-1 ring-indigo-200 dark:ring-indigo-800'
          : 'hover:bg-gray-50 dark:hover:bg-gray-700/50'
      }`}
    >
      <input
        id={id}
        name={name}
        type="radio"
        value={value}
        checked={checked}
        onChange={onChange}
        className="h-5 w-5 border-2 border-gray-300 dark:border-gray-600 text-indigo-600 focus:ring-indigo-500 focus:ring-offset-0 transition-all cursor-pointer"
      />
      <span
        className={`text-sm font-semibold transition-colors ${
          checked
            ? 'text-indigo-700 dark:text-indigo-300'
            : 'text-gray-900 dark:text-white group-hover:text-indigo-600 dark:group-hover:text-indigo-400'
        }`}
      >
        {label}
      </span>
    </label>
  )

  return (
    <section className="mt-12" aria-labelledby="notifications-heading">
      <div className="md:grid md:grid-cols-3 md:gap-8">
        <div className="md:col-span-1">
          <div className="sticky top-6">
            <div className="flex items-center gap-3">
              <div
                className="flex-shrink-0 w-10 h-10 rounded-xl bg-gradient-to-br from-indigo-400 to-purple-500 flex items-center justify-center"
                aria-hidden="true"
              >
                <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                  />
                </svg>
              </div>
              <h2 id="notifications-heading" className="text-xl font-semibold text-gray-900 dark:text-white">
                Notifications
              </h2>
            </div>
            <p className="mt-3 text-sm text-gray-500 dark:text-gray-400 leading-relaxed">
              Decide which communications you'd like to receive and how.
            </p>
          </div>
        </div>
        <div className="mt-6 md:mt-0 md:col-span-2">
          <form onSubmit={handleSubmit} aria-label="Notification preferences form">
            <div className="overflow-hidden rounded-2xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 shadow-sm transition-shadow hover:shadow-md">
              <div className="px-6 py-8 space-y-8">
                <fieldset>
                  <div className="flex items-center gap-3 mb-4">
                    <div className="flex-shrink-0 w-8 h-8 rounded-lg bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center">
                      <svg className="w-4 h-4 text-blue-600 dark:text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                        />
                      </svg>
                    </div>
                    <legend className="text-base font-semibold text-gray-900 dark:text-white">Email Notifications</legend>
                  </div>
                  <div className="space-y-1 pl-11">
                    <CheckboxItem
                      id="email_link_created"
                      name="email_link_created"
                      label="Link Created"
                      description="Get notified when you create a new short link."
                      checked={emailLinkCreated}
                      onChange={(e) => setEmailLinkCreated(e.target.checked)}
                    />
                    <CheckboxItem
                      id="email_link_clicked"
                      name="email_link_clicked"
                      label="Link Clicks"
                      description="Get notified when someone clicks on your links."
                      checked={emailLinkClicked}
                      onChange={(e) => setEmailLinkClicked(e.target.checked)}
                    />
                    <CheckboxItem
                      id="email_weekly_report"
                      name="email_weekly_report"
                      label="Weekly Report"
                      description="Receive a weekly summary of your link statistics."
                      checked={emailWeeklyReport}
                      onChange={(e) => setEmailWeeklyReport(e.target.checked)}
                    />
                    <CheckboxItem
                      id="email_security_alerts"
                      name="email_security_alerts"
                      label="Security Alerts"
                      description="Get notified about important security events on your account."
                      checked={emailSecurityAlerts}
                      onChange={(e) => setEmailSecurityAlerts(e.target.checked)}
                    />
                  </div>
                </fieldset>

                <div className="border-t border-gray-100 dark:border-gray-700 pt-8">
                  <fieldset>
                    <div className="flex items-center gap-3 mb-4">
                      <div className="flex-shrink-0 w-8 h-8 rounded-lg bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center">
                        <svg className="w-4 h-4 text-purple-600 dark:text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                          />
                        </svg>
                      </div>
                      <div>
                        <legend className="text-base font-semibold text-gray-900 dark:text-white">Push Notifications</legend>
                        <p className="text-sm text-gray-500 dark:text-gray-400 mt-0.5">Delivered via browser push notifications</p>
                      </div>
                    </div>
                    <div className="space-y-2 pl-11">
                      <RadioItem
                        id="push_everything"
                        name="push_notifications"
                        value="everything"
                        label="Everything"
                        checked={pushNotifications === 'everything'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                      />
                      <RadioItem
                        id="push_same"
                        name="push_notifications"
                        value="same"
                        label="Same as email"
                        checked={pushNotifications === 'same'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                      />
                      <RadioItem
                        id="push_nothing"
                        name="push_notifications"
                        value="nothing"
                        label="No push notifications"
                        checked={pushNotifications === 'nothing'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                      />
                    </div>
                  </fieldset>
                </div>
              </div>
              <div className="border-t border-gray-100 dark:border-gray-700 bg-gray-50/50 dark:bg-gray-800/50 px-6 py-4 flex justify-end">
                <Button
                  type="submit"
                  variant="contained"
                  disabled={loading}
                  aria-busy={loading}
                  sx={{
                    bgcolor: '#4f46e5',
                    borderRadius: '12px',
                    textTransform: 'none',
                    fontWeight: 600,
                    px: 4,
                    py: 1.25,
                    boxShadow: '0 4px 14px 0 rgba(79, 70, 229, 0.39)',
                    '&:hover': {
                      bgcolor: '#4338ca',
                      boxShadow: '0 6px 20px 0 rgba(79, 70, 229, 0.5)',
                    },
                    '&:disabled': {
                      bgcolor: '#a5b4fc',
                    },
                  }}
                >
                  {loading ? <CircularProgress size={18} color="inherit" aria-label="Saving..." /> : 'Save preferences'}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </section>
  )
}
