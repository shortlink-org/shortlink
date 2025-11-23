import * as React from 'react'
import { useState } from 'react'
import { ErrorAlert, SuccessAlert } from '@/components/common'
import { Button, CircularProgress } from '@mui/material'

export default function Notifications() {
  const [emailLinkCreated, setEmailLinkCreated] = useState(true)
  const [emailLinkClicked, setEmailLinkClicked] = useState(false)
  const [emailWeeklyReport, setEmailWeeklyReport] = useState(true)
  const [emailSecurityAlerts, setEmailSecurityAlerts] = useState(true)
  const [pushNotifications, setPushNotifications] = useState('same')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      // TODO: Integrate with API to save notification preferences
      // For now, just simulate success
      await new Promise(resolve => setTimeout(resolve, 1000))
      setSuccess(true)
      setTimeout(() => setSuccess(false), 3000)
    } catch (err: any) {
      setError(err.message || 'Failed to update notification preferences')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="mt-10 sm:mt-0">
      <div className="md:grid md:grid-cols-3 md:gap-6">
        <div className="md:col-span-1">
          <div className="px-4 sm:px-0">
            <h3 className="text-lg font-medium leading-6 text-gray-900 dark:text-gray-200">Notifications</h3>
            <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">
              Decide which communications you'd like to receive and how.
            </p>
          </div>
        </div>
        <div className="mt-5 md:mt-0 md:col-span-2">
          <form onSubmit={handleSubmit}>
            <div className="shadow overflow-hidden sm:rounded-md">
              <div className="px-4 py-5 bg-white dark:bg-gray-800 space-y-6 sm:p-6">
                <ErrorAlert error={error} onClose={() => setError(null)} />
                <SuccessAlert message={success ? 'Notification preferences updated successfully!' : null} onClose={() => setSuccess(false)} />

                <fieldset>
                  <legend className="text-base font-medium text-gray-900 dark:text-gray-200">By Email</legend>
                  <div className="mt-4 space-y-4">
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          id="email_link_created"
                          name="email_link_created"
                          type="checkbox"
                          checked={emailLinkCreated}
                          onChange={(e) => setEmailLinkCreated(e.target.checked)}
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label htmlFor="email_link_created" className="font-medium text-gray-700 dark:text-gray-300">
                          Link Created
                        </label>
                        <p className="text-gray-500 dark:text-gray-400">Get notified when you create a new short link.</p>
                      </div>
                    </div>
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          id="email_link_clicked"
                          name="email_link_clicked"
                          type="checkbox"
                          checked={emailLinkClicked}
                          onChange={(e) => setEmailLinkClicked(e.target.checked)}
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label htmlFor="email_link_clicked" className="font-medium text-gray-700 dark:text-gray-300">
                          Link Clicks
                        </label>
                        <p className="text-gray-500 dark:text-gray-400">Get notified when someone clicks on your links.</p>
                      </div>
                    </div>
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          id="email_weekly_report"
                          name="email_weekly_report"
                          type="checkbox"
                          checked={emailWeeklyReport}
                          onChange={(e) => setEmailWeeklyReport(e.target.checked)}
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label htmlFor="email_weekly_report" className="font-medium text-gray-700 dark:text-gray-300">
                          Weekly Report
                        </label>
                        <p className="text-gray-500 dark:text-gray-400">Receive a weekly summary of your link statistics.</p>
                      </div>
                    </div>
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          id="email_security_alerts"
                          name="email_security_alerts"
                          type="checkbox"
                          checked={emailSecurityAlerts}
                          onChange={(e) => setEmailSecurityAlerts(e.target.checked)}
                          className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600 rounded"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label htmlFor="email_security_alerts" className="font-medium text-gray-700 dark:text-gray-300">
                          Security Alerts
                        </label>
                        <p className="text-gray-500 dark:text-gray-400">Get notified about important security events on your account.</p>
                      </div>
                    </div>
                  </div>
                </fieldset>
                <fieldset>
                  <div>
                    <legend className="text-base font-medium text-gray-900 dark:text-gray-200">Push Notifications</legend>
                    <p className="text-sm text-gray-500 dark:text-gray-400">These are delivered via browser push notifications.</p>
                  </div>
                  <div className="mt-4 space-y-4">
                    <div className="flex items-center">
                      <input
                        id="push_everything"
                        name="push_notifications"
                        type="radio"
                        value="everything"
                        checked={pushNotifications === 'everything'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                        className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600"
                      />
                      <label htmlFor="push_everything" className="ml-3 block text-sm font-medium text-gray-700 dark:text-gray-300">
                        Everything
                      </label>
                    </div>
                    <div className="flex items-center">
                      <input
                        id="push_same"
                        name="push_notifications"
                        type="radio"
                        value="same"
                        checked={pushNotifications === 'same'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                        className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600"
                      />
                      <label htmlFor="push_same" className="ml-3 block text-sm font-medium text-gray-700 dark:text-gray-300">
                        Same as email
                      </label>
                    </div>
                    <div className="flex items-center">
                      <input
                        id="push_nothing"
                        name="push_notifications"
                        type="radio"
                        value="nothing"
                        checked={pushNotifications === 'nothing'}
                        onChange={(e) => setPushNotifications(e.target.value)}
                        className="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 dark:border-gray-600"
                      />
                      <label htmlFor="push_nothing" className="ml-3 block text-sm font-medium text-gray-700 dark:text-gray-300">
                        No push notifications
                      </label>
                    </div>
                  </div>
                </fieldset>
              </div>
              <div className="px-4 py-3 bg-gray-50 dark:bg-gray-700 text-right sm:px-6">
                <Button
                  type="submit"
                  variant="contained"
                  disabled={loading}
                  sx={{
                    bgcolor: 'indigo.600',
                    '&:hover': { bgcolor: 'indigo.700' },
                  }}
                >
                  {loading ? <CircularProgress size={16} color="inherit" /> : 'Save'}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  )
}
