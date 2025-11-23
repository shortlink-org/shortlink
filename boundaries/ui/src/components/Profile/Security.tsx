import * as React from 'react'
import { useState } from 'react'
import { ErrorAlert, SuccessAlert } from '@/components/common'
import { validatePassword, validatePasswordMatch } from '@/utils/validation'
import { Button, CircularProgress } from '@mui/material'

export default function Security() {
  const [oldPassword, setOldPassword] = useState('')
  const [newPassword, setNewPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError(null)
    setSuccess(false)

    // Validate required fields
    if (!oldPassword || !newPassword || !confirmPassword) {
      setError('All fields are required')
      return
    }

    // Validate password strength
    const passwordValidation = validatePassword(newPassword)
    if (!passwordValidation.isValid) {
      setError(passwordValidation.error || 'Invalid password')
      return
    }

    // Validate password match
    const matchValidation = validatePasswordMatch(newPassword, confirmPassword)
    if (!matchValidation.isValid) {
      setError(matchValidation.error || 'Passwords do not match')
      return
    }

    setLoading(true)

    try {
      // TODO: Integrate with Ory Kratos API to change password
      // For now, just simulate success
      await new Promise(resolve => setTimeout(resolve, 1000))
      setSuccess(true)
      setOldPassword('')
      setNewPassword('')
      setConfirmPassword('')
      setTimeout(() => setSuccess(false), 3000)
    } catch (err: any) {
      setError(err.message || 'Failed to change password')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="md:grid md:grid-cols-3 md:gap-6">
      <div className="md:col-span-1">
        <div className="px-4 sm:px-0">
          <h3 className="text-lg font-medium leading-6 text-gray-900 dark:text-gray-200">Password</h3>
          <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">Change your password to keep your account secure.</p>
        </div>
      </div>

      <div className="mt-5 md:mt-0 md:col-span-2">
        <form onSubmit={handleSubmit}>
          <div className="shadow overflow-hidden sm:rounded-md">
            <div className="px-4 py-5 bg-white dark:bg-gray-800 sm:p-6">
              <ErrorAlert error={error} onClose={() => setError(null)} />
              <SuccessAlert message={success ? 'Password changed successfully!' : null} onClose={() => setSuccess(false)} />
              
              <div className="grid grid-cols-6 gap-6">
                <div className="col-span-6">
                  <label htmlFor="old_password" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    Current Password
                  </label>
                  <input
                    type="password"
                    name="old_password"
                    id="old_password"
                    autoComplete="current-password"
                    value={oldPassword}
                    onChange={(e) => setOldPassword(e.target.value)}
                    className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white rounded-md"
                    required
                  />
                </div>

                <div className="col-span-6 sm:col-span-3">
                  <label htmlFor="new_password" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    New Password
                  </label>
                  <input
                    type="password"
                    name="new_password"
                    id="new_password"
                    autoComplete="new-password"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                    className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white rounded-md"
                    required
                  />
                  <p className="mt-1 text-xs text-gray-500">At least 8 characters with uppercase, lowercase, and number</p>
                </div>

                <div className="col-span-6 sm:col-span-3">
                  <label htmlFor="confirm_new_password" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    Confirm New Password
                  </label>
                  <input
                    type="password"
                    name="confirm_new_password"
                    id="confirm_new_password"
                    autoComplete="new-password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white rounded-md"
                    required
                  />
                </div>
              </div>
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
                {loading ? <CircularProgress size={16} color="inherit" /> : 'Change Password'}
              </Button>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}
