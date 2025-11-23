import * as React from 'react'
import { useState } from 'react'
import { ErrorAlert, SuccessAlert } from '@/components/common'
import { validateUrl } from '@/utils/validation'
import { Button, CircularProgress } from '@mui/material'

export default function Profile() {
  const [website, setWebsite] = useState('')
  const [about, setAbout] = useState('')
  const [avatarFile, setAvatarFile] = useState<File | null>(null)
  const [coverFile, setCoverFile] = useState<File | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    // Validate URL if provided
    if (website) {
      const urlValidation = validateUrl(website, false)
      if (!urlValidation.isValid) {
        setError(urlValidation.error || 'Invalid URL')
        return
      }
    }

    setLoading(true)
    setError(null)
    setSuccess(false)

    try {
      // TODO: Integrate with API to update profile
      // For now, just simulate success
      await new Promise(resolve => setTimeout(resolve, 1000))
      setSuccess(true)
    } catch (err: any) {
      setError(err.message || 'Failed to update profile')
    } finally {
      setLoading(false)
    }
  }

  const handleAvatarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      if (file.size > 10 * 1024 * 1024) {
        setError('Avatar file size must be less than 10MB')
        return
      }
      setAvatarFile(file)
    }
  }

  const handleCoverChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      if (file.size > 10 * 1024 * 1024) {
        setError('Cover photo file size must be less than 10MB')
        return
      }
      setCoverFile(file)
    }
  }

  return (
    <div className="md:grid md:grid-cols-3 md:gap-6">
      <div className="md:col-span-1">
        <div className="px-4 sm:px-0">
          <h3 className="text-lg font-medium leading-6 text-gray-900 dark:text-gray-200">Profile</h3>
          <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">
            This information will be displayed publicly so be careful what you share.
          </p>
        </div>
      </div>
      <div className="mt-5 md:col-span-2 md:mt-0">
        <form onSubmit={handleSubmit}>
          <div className="shadow sm:overflow-hidden sm:rounded-md">
            <div className="space-y-6 bg-white dark:bg-gray-800 px-4 py-5 sm:p-6">
              <ErrorAlert error={error} onClose={() => setError(null)} />
              <SuccessAlert message={success ? 'Profile updated successfully!' : null} onClose={() => setSuccess(false)} />

              <div className="grid grid-cols-3 gap-6">
                <div className="col-span-3 sm:col-span-2">
                  <label htmlFor="website" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    Website
                  </label>
                  <div className="mt-1 flex rounded-md shadow-sm">
                    <span className="inline-flex items-center rounded-l-md border border-r-0 border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 px-3 text-sm text-gray-500 dark:text-gray-400">
                      https://
                    </span>
                    <input
                      type="text"
                      name="website"
                      id="website"
                      value={website}
                      onChange={(e) => setWebsite(e.target.value)}
                      className="block w-full flex-1 rounded-none rounded-r-md border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                      placeholder="example.com"
                    />
                  </div>
                </div>
              </div>

              <div>
                <label htmlFor="about" className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                  About
                </label>
                <div className="mt-1">
                  <textarea
                    id="about"
                    name="about"
                    rows={3}
                    value={about}
                    onChange={(e) => setAbout(e.target.value)}
                    className="mt-1 block w-full rounded-md border-gray-300 dark:border-gray-600 dark:bg-gray-700 dark:text-white shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                    placeholder="Tell us about yourself..."
                  />
                </div>
                <p className="mt-2 text-sm text-gray-500 dark:text-gray-400">Brief description for your profile. URLs are hyperlinked.</p>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Photo</label>
                <div className="mt-1 flex items-center">
                  <span className="inline-block h-12 w-12 overflow-hidden rounded-full bg-gray-100 dark:bg-gray-700">
                    {avatarFile ? (
                      <img
                        src={URL.createObjectURL(avatarFile)}
                        alt="Avatar preview"
                        className="h-full w-full object-cover"
                      />
                    ) : (
                      <svg className="h-full w-full text-gray-300 dark:text-gray-600" fill="currentColor" viewBox="0 0 24 24">
                        <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" />
                      </svg>
                    )}
                  </span>
                  <label
                    htmlFor="avatar-upload"
                    className="ml-5 rounded-md border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 py-2 px-3 text-sm font-medium leading-4 text-gray-700 dark:text-gray-300 shadow-sm hover:bg-gray-50 dark:hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 cursor-pointer"
                  >
                    <input
                      id="avatar-upload"
                      name="avatar-upload"
                      type="file"
                      accept="image/*"
                      onChange={handleAvatarChange}
                      className="sr-only"
                    />
                    Change
                  </label>
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">Cover photo</label>
                <div className="mt-1 flex justify-center rounded-md border-2 border-dashed border-gray-300 dark:border-gray-600 px-6 pt-5 pb-6">
                  <div className="space-y-1 text-center">
                    {coverFile ? (
                      <div className="mt-2">
                        <img
                          src={URL.createObjectURL(coverFile)}
                          alt="Cover preview"
                          className="mx-auto h-32 w-full object-cover rounded-md"
                        />
                        <button
                          type="button"
                          onClick={() => setCoverFile(null)}
                          className="mt-2 text-sm text-red-600 hover:text-red-800"
                        >
                          Remove
                        </button>
                      </div>
                    ) : (
                      <>
                        <svg className="mx-auto h-12 w-12 text-gray-400" stroke="currentColor" fill="none" viewBox="0 0 48 48" aria-hidden="true">
                          <path
                            d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
                            strokeWidth={2}
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                        </svg>
                        <div className="flex text-sm text-gray-600 dark:text-gray-400">
                          <label
                            htmlFor="cover-upload"
                            className="relative cursor-pointer rounded-md bg-white dark:bg-gray-800 font-medium text-indigo-600 dark:text-indigo-400 focus-within:outline-none focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 hover:text-indigo-500"
                          >
                            <span>Upload a file</span>
                            <input
                              id="cover-upload"
                              name="cover-upload"
                              type="file"
                              accept="image/*"
                              onChange={handleCoverChange}
                              className="sr-only"
                            />
                          </label>
                          <p className="pl-1">or drag and drop</p>
                        </div>
                        <p className="text-xs text-gray-500 dark:text-gray-400">PNG, JPG, GIF up to 10MB</p>
                      </>
                    )}
                  </div>
                </div>
              </div>
            </div>
            <div className="bg-gray-50 dark:bg-gray-700 px-4 py-3 text-right sm:px-6">
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
  )
}
