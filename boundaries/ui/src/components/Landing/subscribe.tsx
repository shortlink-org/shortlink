'use client'

import { useState } from 'react'

function EmailIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={2}
        d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
      />
    </svg>
  )
}

function SendIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
    </svg>
  )
}

function CheckIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
    </svg>
  )
}

export default function Subscribe() {
  const [email, setEmail] = useState('')
  const [submitted, setSubmitted] = useState(false)
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!email) return

    setLoading(true)
    await new Promise((resolve) => setTimeout(resolve, 1000))
    setLoading(false)
    setSubmitted(true)
  }

  return (
    <section className="py-16 lg:py-24">
      <div className="rounded-3xl overflow-hidden bg-gradient-to-br from-indigo-600 via-purple-600 to-pink-500 dark:from-indigo-800 dark:via-purple-800 dark:to-pink-700">
        <div className="grid lg:grid-cols-2 gap-8 items-center">
          {/* Content */}
          <div className="p-8 lg:p-12">
            <span className="text-xs uppercase tracking-widest font-semibold text-indigo-200">Stay Updated</span>
            <h2 className="text-3xl lg:text-4xl font-bold text-white mt-2 mb-4">
              Subscribe to our{' '}
              <span className="text-transparent bg-clip-text bg-gradient-to-r from-yellow-300 to-orange-300">Newsletter</span>
            </h2>
            <p className="text-indigo-100 mb-8 leading-relaxed">
              Be the first to know about new features, updates, and tips for making the most of Shortlink.
            </p>

            {submitted ? (
              <div className="flex items-center gap-3 p-4 bg-white/10 backdrop-blur-sm rounded-xl">
                <CheckIcon className="w-8 h-8 text-green-400" />
                <div>
                  <p className="text-white font-semibold">You're subscribed!</p>
                  <p className="text-sm text-indigo-200">Check your inbox for a confirmation email.</p>
                </div>
              </div>
            ) : (
              <form onSubmit={handleSubmit} className="flex flex-col sm:flex-row gap-3">
                <div className="relative flex-1">
                  <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
                    <EmailIcon className="w-5 h-5 text-gray-400" />
                  </div>
                  <input
                    type="email"
                    placeholder="Enter your email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    className="w-full pl-12 pr-4 py-3 bg-white rounded-xl border-0 text-gray-900 placeholder-gray-400 focus:ring-2 focus:ring-white focus:outline-none"
                  />
                </div>
                <button
                  type="submit"
                  disabled={loading}
                  className="inline-flex items-center justify-center gap-2 px-6 py-3 bg-white text-indigo-600 font-semibold rounded-xl hover:bg-gray-100 hover:-translate-y-0.5 transition-all duration-200 disabled:opacity-70 disabled:cursor-not-allowed whitespace-nowrap"
                >
                  {loading ? 'Subscribing...' : 'Subscribe'}
                  <SendIcon className="w-4 h-4" />
                </button>
              </form>
            )}

            <p className="text-xs text-indigo-200 mt-4">No spam, unsubscribe at any time.</p>
          </div>

          {/* Illustration */}
          <div className="hidden lg:flex items-center justify-center p-8">
            <div className="relative">
              <div className="absolute inset-0 bg-white/10 rounded-full blur-3xl" />
              <img
                src="/assets/images/undraw_designer_re_5v95.svg"
                alt="Newsletter illustration"
                className="relative w-full max-w-sm h-auto drop-shadow-2xl"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
