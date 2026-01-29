import * as React from 'react'
import { useState } from 'react'
import { toast } from 'sonner'
import { validateUrl } from '@/utils/validation'
import { Button, CircularProgress } from '@mui/material'

export default function Profile() {
  const [website, setWebsite] = useState('')
  const [about, setAbout] = useState('')
  const [avatarFile, setAvatarFile] = useState<File | null>(null)
  const [coverFile, setCoverFile] = useState<File | null>(null)
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    // Validate URL if provided
    if (website) {
      const urlValidation = validateUrl(website, false)
      if (!urlValidation.isValid) {
        toast.error('Invalid URL', {
          description: urlValidation.error || 'Please enter a valid website URL',
        })
        return
      }
    }

    setLoading(true)

    try {
      // TODO: Integrate with API to update profile
      // For now, just simulate success
      await new Promise((resolve) => setTimeout(resolve, 1000))
      toast.success('Profile updated', {
        description: 'Your profile has been saved successfully.',
      })
    } catch (err: any) {
      toast.error('Failed to update profile', {
        description: err.message || 'Please try again later.',
      })
    } finally {
      setLoading(false)
    }
  }

  const handleAvatarChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      if (file.size > 10 * 1024 * 1024) {
        toast.error('File too large', {
          description: 'Avatar file size must be less than 10MB',
        })
        return
      }
      setAvatarFile(file)
      toast.success('Avatar selected', { duration: 2000 })
    }
  }

  const handleCoverChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      if (file.size > 10 * 1024 * 1024) {
        toast.error('File too large', {
          description: 'Cover photo file size must be less than 10MB',
        })
        return
      }
      setCoverFile(file)
      toast.success('Cover photo selected', { duration: 2000 })
    }
  }

  return (
    <section className="md:grid md:grid-cols-3 md:gap-8" aria-labelledby="profile-heading">
      <div className="md:col-span-1">
        <div className="sticky top-6">
          <h2 id="profile-heading" className="text-xl font-semibold text-gray-900 dark:text-white">
            Profile
          </h2>
          <p className="mt-2 text-sm text-gray-500 dark:text-gray-400 leading-relaxed">
            This information will be displayed publicly so be careful what you share.
          </p>
        </div>
      </div>
      <div className="mt-6 md:col-span-2 md:mt-0">
        <form onSubmit={handleSubmit} aria-label="Profile information form">
          <div className="overflow-hidden rounded-2xl border border-gray-200 dark:border-gray-700 bg-white dark:bg-gray-800 shadow-sm transition-shadow hover:shadow-md">
            <div className="space-y-8 px-6 py-8">
              <div>
                <label htmlFor="website" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                  Website
                </label>
                <div className="flex rounded-xl overflow-hidden ring-1 ring-gray-300 dark:ring-gray-600 focus-within:ring-2 focus-within:ring-indigo-500 transition-all">
                  <span className="inline-flex items-center border-r border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700/50 px-4 text-sm text-gray-500 dark:text-gray-400 font-medium">
                    https://
                  </span>
                  <input
                    type="text"
                    name="website"
                    id="website"
                    value={website}
                    onChange={(e) => setWebsite(e.target.value)}
                    className="block w-full flex-1 border-0 bg-transparent dark:bg-transparent dark:text-white py-3 px-4 text-sm placeholder:text-gray-400 focus:ring-0 focus:outline-none"
                    placeholder="example.com"
                  />
                </div>
              </div>

              <div>
                <label htmlFor="about" className="block text-sm font-semibold text-gray-900 dark:text-white mb-2">
                  About
                </label>
                <textarea
                  id="about"
                  name="about"
                  rows={4}
                  value={about}
                  onChange={(e) => setAbout(e.target.value)}
                  className="block w-full rounded-xl border-0 ring-1 ring-gray-300 dark:ring-gray-600 bg-white dark:bg-gray-700/50 dark:text-white py-3 px-4 text-sm placeholder:text-gray-400 focus:ring-2 focus:ring-indigo-500 transition-all resize-none"
                  placeholder="Tell us about yourself..."
                />
                <p className="mt-2 text-xs text-gray-500 dark:text-gray-400">Brief description for your profile. URLs are hyperlinked.</p>
              </div>

              <div>
                <label className="block text-sm font-semibold text-gray-900 dark:text-white mb-3">Avatar Photo</label>
                <div className="flex items-center gap-6">
                  <div className="relative">
                    <span className="inline-block h-20 w-20 overflow-hidden rounded-2xl bg-gradient-to-br from-gray-100 to-gray-200 dark:from-gray-700 dark:to-gray-600 ring-4 ring-white dark:ring-gray-800 shadow-lg">
                      {avatarFile ? (
                        <img src={URL.createObjectURL(avatarFile)} alt="Avatar preview" className="h-full w-full object-cover" />
                      ) : (
                        <svg className="h-full w-full text-gray-400 dark:text-gray-500 p-4" fill="currentColor" viewBox="0 0 24 24">
                          <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" />
                        </svg>
                      )}
                    </span>
                  </div>
                  <div className="flex flex-col gap-2">
                    <label
                      htmlFor="avatar-upload"
                      className="inline-flex items-center justify-center rounded-xl bg-white dark:bg-gray-700 px-4 py-2.5 text-sm font-semibold text-gray-700 dark:text-gray-200 ring-1 ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-600 cursor-pointer transition-all hover:ring-gray-400 dark:hover:ring-gray-500"
                    >
                      <input
                        id="avatar-upload"
                        name="avatar-upload"
                        type="file"
                        accept="image/*"
                        onChange={handleAvatarChange}
                        className="sr-only"
                      />
                      <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"
                        />
                      </svg>
                      Change photo
                    </label>
                    <span className="text-xs text-gray-500 dark:text-gray-400">JPG, PNG, GIF up to 10MB</span>
                  </div>
                </div>
              </div>

              <div>
                <label className="block text-sm font-semibold text-gray-900 dark:text-white mb-3">Cover Photo</label>
                <div className="group relative rounded-2xl border-2 border-dashed border-gray-300 dark:border-gray-600 hover:border-indigo-400 dark:hover:border-indigo-500 transition-all cursor-pointer overflow-hidden">
                  {coverFile ? (
                    <div className="relative">
                      <img src={URL.createObjectURL(coverFile)} alt="Cover preview" className="w-full h-48 object-cover" />
                      <div className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
                        <button
                          type="button"
                          onClick={() => setCoverFile(null)}
                          aria-label="Remove cover photo"
                          className="inline-flex items-center rounded-xl bg-red-500 px-4 py-2 text-sm font-semibold text-white hover:bg-red-600 transition-colors"
                        >
                          <svg className="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth={2}
                              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                            />
                          </svg>
                          Remove
                        </button>
                      </div>
                    </div>
                  ) : (
                    <label htmlFor="cover-upload" className="flex flex-col items-center justify-center py-12 px-6 cursor-pointer">
                      <div className="rounded-full bg-indigo-50 dark:bg-indigo-900/30 p-4 mb-4 group-hover:bg-indigo-100 dark:group-hover:bg-indigo-900/50 transition-colors">
                        <svg
                          className="h-8 w-8 text-indigo-500 dark:text-indigo-400"
                          stroke="currentColor"
                          fill="none"
                          viewBox="0 0 48 48"
                          aria-hidden="true"
                        >
                          <path
                            d="M28 8H12a4 4 0 00-4 4v20m32-12v8m0 0v8a4 4 0 01-4 4H12a4 4 0 01-4-4v-4m32-4l-3.172-3.172a4 4 0 00-5.656 0L28 28M8 32l9.172-9.172a4 4 0 015.656 0L28 28m0 0l4 4m4-24h8m-4-4v8m-12 4h.02"
                            strokeWidth={2}
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                        </svg>
                      </div>
                      <div className="text-center">
                        <span className="text-sm font-semibold text-indigo-600 dark:text-indigo-400">Upload a file</span>
                        <span className="text-sm text-gray-500 dark:text-gray-400"> or drag and drop</span>
                      </div>
                      <p className="text-xs text-gray-500 dark:text-gray-400 mt-2">PNG, JPG, GIF up to 10MB</p>
                      <input
                        id="cover-upload"
                        name="cover-upload"
                        type="file"
                        accept="image/*"
                        onChange={handleCoverChange}
                        className="sr-only"
                      />
                    </label>
                  )}
                </div>
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
                {loading ? <CircularProgress size={18} color="inherit" aria-label="Saving..." /> : 'Save changes'}
              </Button>
            </div>
          </div>
        </form>
      </div>
    </section>
  )
}
