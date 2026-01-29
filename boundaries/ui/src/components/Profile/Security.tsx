import * as React from 'react'
import { useState } from 'react'
import { toast } from 'sonner'
import { PasswordInput } from '@/components/common'
import { validatePassword, validatePasswordMatch } from '@/utils/validation'
import { Button, CircularProgress } from '@mui/material'

export default function Security() {
  const [oldPassword, setOldPassword] = useState('')
  const [newPassword, setNewPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    // Validate required fields
    if (!oldPassword || !newPassword || !confirmPassword) {
      toast.error('All fields are required', {
        description: 'Please fill in all password fields.',
      })
      return
    }

    // Validate password strength
    const passwordValidation = validatePassword(newPassword)
    if (!passwordValidation.isValid) {
      toast.error('Weak password', {
        description: passwordValidation.error || 'Password does not meet requirements.',
      })
      return
    }

    // Validate password match
    const matchValidation = validatePasswordMatch(newPassword, confirmPassword)
    if (!matchValidation.isValid) {
      toast.error('Passwords do not match', {
        description: 'Please make sure both passwords are the same.',
      })
      return
    }

    setLoading(true)

    try {
      // TODO: Integrate with Ory Kratos API to change password
      // For now, just simulate success
      await new Promise((resolve) => setTimeout(resolve, 1000))
      toast.success('Password changed', {
        description: 'Your password has been updated successfully.',
      })
      setOldPassword('')
      setNewPassword('')
      setConfirmPassword('')
    } catch (err: any) {
      toast.error('Failed to change password', {
        description: err.message || 'Please try again later.',
      })
    } finally {
      setLoading(false)
    }
  }

  return (
    <section className="mt-12" aria-labelledby="security-heading">
      <div className="md:grid md:grid-cols-3 md:gap-8">
        <div className="md:col-span-1">
          <div className="sticky top-6">
            <div className="flex items-center gap-3">
              <div
                className="flex-shrink-0 w-10 h-10 rounded-xl bg-gradient-to-br from-amber-400 to-orange-500 flex items-center justify-center"
                aria-hidden="true"
              >
                <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                  />
                </svg>
              </div>
              <h2 id="security-heading" className="text-xl font-semibold text-gray-900 dark:text-white">
                Password
              </h2>
            </div>
            <p className="mt-3 text-sm text-gray-500 dark:text-gray-400 leading-relaxed">
              Change your password to keep your account secure. We recommend using a strong, unique password.
            </p>
          </div>
        </div>

        <div className="mt-6 md:mt-0 md:col-span-2">
          <form onSubmit={handleSubmit} aria-label="Change password form">
            <div className="overflow-hidden rounded-2xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 shadow-sm transition-shadow hover:shadow-md">
              <div className="px-6 py-8 space-y-6">
                <PasswordInput
                  id="old_password"
                  name="old_password"
                  label="Current Password"
                  value={oldPassword}
                  onChange={(e) => setOldPassword(e.target.value)}
                  autoComplete="current-password"
                />

                <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                  <PasswordInput
                    id="new_password"
                    name="new_password"
                    label="New Password"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                    hint="At least 8 characters with uppercase, lowercase, and number"
                  />

                  <PasswordInput
                    id="confirm_new_password"
                    name="confirm_new_password"
                    label="Confirm New Password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                  />
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
                  {loading ? <CircularProgress size={18} color="inherit" aria-label="Saving..." /> : 'Change password'}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </section>
  )
}
