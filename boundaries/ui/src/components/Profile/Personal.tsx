/**
 * Personal Information Form - Migrated to React 19 with useOptimistic
 *
 * Changes:
 * - Added useOptimistic for instant UI updates
 * - Added useTransition for async coordination
 * - Replaced manual loading state with transition
 * - Form values update instantly (optimistic)
 * - Automatic rollback on error
 * - Toast notifications for feedback
 * - Accessibility improvements
 */

import * as React from 'react'
import { useState, useOptimistic, useTransition } from 'react'
import { Session } from '@ory/client'
import { toast } from 'sonner'
import { validateEmail, validateRequired } from '@/utils/validation'
import { Button, CircularProgress } from '@mui/material'

interface PersonalProps {
  session: Session
  firstName: string
  lastName: string
  email: string
}

type FormData = {
  firstName: string
  lastName: string
  email: string
  country: string
  streetAddress: string
  city: string
  region: string
  postalCode: string
}

export default function Personal({ session, firstName: initialFirstName, lastName: initialLastName, email: initialEmail }: PersonalProps) {
  // Initialize form data
  const initialData: FormData = {
    firstName: initialFirstName,
    lastName: initialLastName,
    email: initialEmail,
    country: '',
    streetAddress: '',
    city: '',
    region: '',
    postalCode: '',
  }

  // useOptimistic for instant UI updates
  const [optimisticData, setOptimisticData] = useOptimistic(initialData)
  const [isPending, startTransition] = useTransition()
  const [country, setCountry] = useState(initialData.country)
  const [streetAddress, setStreetAddress] = useState(initialData.streetAddress)
  const [city, setCity] = useState(initialData.city)
  const [region, setRegion] = useState(initialData.region)
  const [postalCode, setPostalCode] = useState(initialData.postalCode)

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const formData = new FormData(e.currentTarget)
    const firstName = formData.get('first_name') as string
    const email = formData.get('email_address') as string

    // Validate required fields
    const firstNameValidation = validateRequired(firstName, 'First name')
    const emailValidation = validateEmail(email, true)

    if (!firstNameValidation.isValid) {
      toast.error('Validation error', {
        description: firstNameValidation.error || 'First name is required',
      })
      return
    }

    if (!emailValidation.isValid) {
      toast.error('Validation error', {
        description: emailValidation.error || 'Valid email is required',
      })
      return
    }

    const newData: FormData = {
      firstName: firstName,
      lastName: formData.get('last_name') as string,
      email: email,
      country: formData.get('country') as string,
      streetAddress: formData.get('street_address') as string,
      city: formData.get('city') as string,
      region: formData.get('region') as string,
      postalCode: formData.get('postal_code') as string,
    }

    // Start transition with optimistic update
    startTransition(async () => {
      // Optimistically update UI immediately
      setOptimisticData(newData)

      try {
        // TODO: Integrate with Ory Kratos API to update profile
        // For now, just simulate success
        await new Promise((resolve) => setTimeout(resolve, 1000))
        toast.success('Personal information updated', {
          description: 'Your changes have been saved successfully.',
        })
      } catch (err: any) {
        // If error, optimistic data will automatically revert to initialData
        toast.error('Failed to update', {
          description: err.message || 'Please try again later.',
        })
      }
    })
  }

  const inputClassName =
    'block w-full rounded-xl border-0 ring-1 ring-gray-300 dark:ring-gray-600 bg-white dark:bg-gray-700/50 dark:text-white py-3 px-4 text-sm placeholder:text-gray-400 focus:ring-2 focus:ring-indigo-500 transition-all'

  return (
    <section className="mt-12" aria-labelledby="personal-heading">
      <div className="md:grid md:grid-cols-3 md:gap-8">
        <div className="md:col-span-1">
          <div className="sticky top-6">
            <h2 id="personal-heading" className="text-xl font-semibold text-gray-900 dark:text-white">
              Personal Information
            </h2>
            <p className="mt-2 text-sm text-gray-500 dark:text-gray-400 leading-relaxed">
              Update your personal information and contact details.
            </p>
          </div>
        </div>

        <div className="mt-6 md:mt-0 md:col-span-2">
          <form onSubmit={handleSubmit} aria-label="Personal information form">
            <div className="overflow-hidden rounded-2xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 shadow-sm transition-shadow hover:shadow-md">
              <div className="px-6 py-8" style={{ opacity: isPending ? 0.7 : 1, transition: 'opacity 0.2s' }}>
                <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                  <div>
                    <label htmlFor="first_name" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      First name{' '}
                      <span className="text-red-500" aria-label="required">
                        *
                      </span>
                    </label>
                    <input
                      type="text"
                      name="first_name"
                      id="first_name"
                      autoComplete="given-name"
                      defaultValue={optimisticData.firstName}
                      className={inputClassName}
                      required
                      aria-required="true"
                    />
                  </div>

                  <div>
                    <label htmlFor="last_name" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      Last name
                    </label>
                    <input
                      type="text"
                      name="last_name"
                      id="last_name"
                      autoComplete="family-name"
                      defaultValue={optimisticData.lastName}
                      className={inputClassName}
                    />
                  </div>

                  <div className="sm:col-span-2">
                    <label htmlFor="email_address" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      Email address{' '}
                      <span className="text-red-500" aria-label="required">
                        *
                      </span>
                    </label>
                    <div className="relative">
                      <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none" aria-hidden="true">
                        <svg className="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={1.5}
                            d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                          />
                        </svg>
                      </div>
                      <input
                        type="email"
                        name="email_address"
                        id="email_address"
                        autoComplete="email"
                        defaultValue={optimisticData.email}
                        className={`${inputClassName} pl-12`}
                        required
                        aria-required="true"
                      />
                    </div>
                  </div>

                  <div className="sm:col-span-2">
                    <label htmlFor="country" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      Country / Region
                    </label>
                    <select
                      id="country"
                      name="country"
                      autoComplete="country"
                      value={country}
                      onChange={(e) => setCountry(e.target.value)}
                      className={inputClassName}
                    >
                      <option value="">Select a country</option>
                      <option value="US">United States</option>
                      <option value="CA">Canada</option>
                      <option value="MX">Mexico</option>
                      <option value="GB">United Kingdom</option>
                      <option value="DE">Germany</option>
                      <option value="FR">France</option>
                      <option value="RU">Russia</option>
                    </select>
                  </div>

                  <div className="sm:col-span-2">
                    <label htmlFor="street_address" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      Street address
                    </label>
                    <input
                      type="text"
                      name="street_address"
                      id="street_address"
                      autoComplete="street-address"
                      value={streetAddress}
                      onChange={(e) => setStreetAddress(e.target.value)}
                      className={inputClassName}
                    />
                  </div>

                  <div>
                    <label htmlFor="city" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      City
                    </label>
                    <input
                      type="text"
                      name="city"
                      id="city"
                      autoComplete="address-level2"
                      value={city}
                      onChange={(e) => setCity(e.target.value)}
                      className={inputClassName}
                    />
                  </div>

                  <div>
                    <label htmlFor="region" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      State / Province
                    </label>
                    <input
                      type="text"
                      name="region"
                      id="region"
                      autoComplete="address-level1"
                      value={region}
                      onChange={(e) => setRegion(e.target.value)}
                      className={inputClassName}
                    />
                  </div>

                  <div>
                    <label htmlFor="postal_code" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                      ZIP / Postal code
                    </label>
                    <input
                      type="text"
                      name="postal_code"
                      id="postal_code"
                      autoComplete="postal-code"
                      value={postalCode}
                      onChange={(e) => setPostalCode(e.target.value)}
                      className={inputClassName}
                    />
                  </div>
                </div>
              </div>
              <div className="border-t border-gray-100 dark:border-gray-700 bg-gray-50/50 dark:bg-gray-800/50 px-6 py-4 flex justify-end">
                <Button
                  type="submit"
                  variant="contained"
                  disabled={isPending}
                  aria-busy={isPending}
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
                  {isPending ? <CircularProgress size={18} color="inherit" aria-label="Saving..." /> : 'Save changes'}
                </Button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </section>
  )
}
