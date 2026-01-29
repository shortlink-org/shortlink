'use client'

function AppleIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 305 305">
      <path d="M40.74 112.12c-25.79 44.74-9.4 112.65 19.12 153.82C74.09 286.52 88.5 305 108.24 305c.37 0 .74 0 1.13-.02 9.27-.37 15.97-3.23 22.45-5.99 7.27-3.1 14.8-6.3 26.6-6.3 11.22 0 18.39 3.1 25.31 6.1 6.83 2.95 13.87 6 24.26 5.81 22.23-.41 35.88-20.35 47.92-37.94a168.18 168.18 0 0021-43l.09-.28a2.5 2.5 0 00-1.33-3.06l-.18-.08c-3.92-1.6-38.26-16.84-38.62-58.36-.34-33.74 25.76-51.6 31-54.84l.24-.15a2.5 2.5 0 00.7-3.51c-18-26.37-45.62-30.34-56.73-30.82a50.04 50.04 0 00-4.95-.24c-13.06 0-25.56 4.93-35.61 8.9-6.94 2.73-12.93 5.09-17.06 5.09-4.64 0-10.67-2.4-17.65-5.16-9.33-3.7-19.9-7.9-31.1-7.9l-.79.01c-26.03.38-50.62 15.27-64.18 38.86z" />
      <path d="M212.1 0c-15.76.64-34.67 10.35-45.97 23.58-9.6 11.13-19 29.68-16.52 48.38a2.5 2.5 0 002.29 2.17c1.06.08 2.15.12 3.23.12 15.41 0 32.04-8.52 43.4-22.25 11.94-14.5 17.99-33.1 16.16-49.77A2.52 2.52 0 00212.1 0z" />
    </svg>
  )
}

function PlayStoreIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 512 512">
      <path d="M99.617 8.057a50.191 50.191 0 00-38.815-6.713l230.932 230.933 74.846-74.846L99.617 8.057zM32.139 20.116c-6.441 8.563-10.148 19.077-10.148 30.199v411.358c0 11.123 3.708 21.636 10.148 30.199l235.877-235.877L32.139 20.116zM464.261 212.087l-67.266-37.637-81.544 81.544 81.548 81.548 67.273-37.64c16.117-9.03 25.738-25.442 25.738-43.908s-9.621-34.877-25.749-43.907zM291.733 279.711L60.815 510.629c3.786.891 7.639 1.371 11.492 1.371a50.275 50.275 0 0027.31-8.07l266.965-149.372-74.849-74.847z" />
    </svg>
  )
}

function PhoneIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth={2}
        d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"
      />
    </svg>
  )
}

export default function MobileApps() {
  return (
    <section className="py-16 lg:py-24">
      <div className="rounded-3xl overflow-hidden bg-gradient-to-br from-slate-50 to-gray-100 dark:from-gray-800 dark:to-gray-900 border border-gray-200 dark:border-gray-700">
        <div className="grid lg:grid-cols-2 gap-8 items-center">
          {/* Content */}
          <div className="p-8 lg:p-12">
            <div className="flex items-center gap-2 mb-4">
              <PhoneIcon className="w-5 h-5 text-indigo-600 dark:text-indigo-400" />
              <span className="text-xs uppercase tracking-widest font-semibold text-indigo-600 dark:text-indigo-400">Mobile Apps</span>
            </div>

            <h2 className="text-3xl lg:text-4xl font-bold text-gray-900 dark:text-white mb-4">Shortlink on the go</h2>

            <p className="text-gray-600 dark:text-gray-400 mb-8 leading-relaxed">
              Access your links anywhere, anytime. Our mobile apps sync seamlessly across all your devices — iOS, Android, and desktop.
            </p>

            <div className="flex flex-col sm:flex-row gap-4">
              <button
                type="button"
                className="group inline-flex items-center gap-3 px-6 py-3 bg-gray-900 dark:bg-white text-white dark:text-gray-900 rounded-xl shadow-lg hover:shadow-xl hover:-translate-y-0.5 transition-all duration-200"
              >
                <AppleIcon className="w-7 h-7" />
                <span className="text-left">
                  <span className="block text-xs text-gray-400 dark:text-gray-500 leading-none">Download on the</span>
                  <span className="block font-bold text-base leading-tight">App Store</span>
                </span>
              </button>

              <button
                type="button"
                className="group inline-flex items-center gap-3 px-6 py-3 bg-gray-900 dark:bg-white text-white dark:text-gray-900 rounded-xl shadow-lg hover:shadow-xl hover:-translate-y-0.5 transition-all duration-200"
              >
                <PlayStoreIcon className="w-7 h-7" />
                <span className="text-left">
                  <span className="block text-xs text-gray-400 dark:text-gray-500 leading-none">GET IT ON</span>
                  <span className="block font-bold text-base leading-tight">Google Play</span>
                </span>
              </button>
            </div>
          </div>

          {/* Visual */}
          <div className="hidden lg:flex items-center justify-center p-8 relative">
            <div className="relative">
              {/* Background decoration */}
              <div className="absolute -inset-8 bg-gradient-to-br from-indigo-500/20 to-purple-500/20 rounded-full blur-3xl" />

              {/* Phone mockups */}
              <div className="relative flex items-center justify-center gap-4">
                <div className="w-48 h-96 bg-gradient-to-br from-gray-800 to-gray-900 rounded-[2.5rem] p-2 shadow-2xl transform -rotate-6">
                  <div className="w-full h-full bg-gradient-to-br from-indigo-600 to-purple-600 rounded-[2rem] flex items-center justify-center">
                    <span className="text-4xl font-bold text-white">S</span>
                  </div>
                </div>
                <div className="w-48 h-96 bg-gradient-to-br from-gray-800 to-gray-900 rounded-[2.5rem] p-2 shadow-2xl transform rotate-6 -ml-12">
                  <div className="w-full h-full bg-gradient-to-br from-purple-600 to-pink-600 rounded-[2rem] flex items-center justify-center">
                    <span className="text-4xl font-bold text-white">L</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
