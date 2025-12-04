import * as React from 'react'
import Balancer from 'react-wrap-balancer'

export default function Feature() {
  return (
    <>
      {/* Main Features Section */}
      <section className="container mx-auto px-4 sm:px-6 py-16 lg:py-24">
        <h2 className="text-4xl lg:text-5xl font-bold text-center bg-gradient-to-r from-indigo-600 to-purple-600 dark:from-indigo-400 dark:to-purple-400 bg-clip-text text-transparent pb-4 mb-16">
          Features
        </h2>

        {/* Feature 1 - URL Shortener */}
        <div className="flex items-center flex-wrap mb-16 lg:mb-24 group">
          <div className="w-full lg:w-1/2 lg:pr-12 mb-8 lg:mb-0">
            <div className="bg-gradient-to-br from-indigo-50 to-purple-50 dark:from-indigo-950 dark:to-purple-950 p-8 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-indigo-100 dark:border-indigo-800">
              <div className="inline-block p-3 bg-indigo-600 dark:bg-indigo-500 rounded-xl mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                </svg>
              </div>
              <h4 className="text-3xl lg:text-4xl text-gray-900 dark:text-white font-bold mb-4">URL Shortener</h4>
              <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed">
                Free custom URL Shortener with many features that gives you better quality for links shortening. We do not display ads during direct redirecting to the original url.
              </p>
            </div>
          </div>
          <div className="w-full lg:w-1/2 flex justify-center">
            <div className="relative group-hover:scale-105 transition-transform duration-300">
              <div className="absolute inset-0 bg-gradient-to-r from-indigo-400 to-purple-400 rounded-2xl blur-2xl opacity-30 group-hover:opacity-50 transition-opacity" />
              <img 
                src="https://www.dropbox.com/s/mimcvn6zxtoruis/health.svg?raw=1" 
                alt="Monitoring" 
                className="relative rounded-2xl shadow-xl"
              />
            </div>
          </div>
        </div>

        {/* Feature 2 - Reporting */}
        <div className="flex items-center flex-wrap mb-16 lg:mb-24 flex-col-reverse lg:flex-row group">
          <div className="w-full lg:w-1/2 flex justify-center mb-8 lg:mb-0">
            <div className="relative group-hover:scale-105 transition-transform duration-300">
              <div className="absolute inset-0 bg-gradient-to-r from-purple-400 to-pink-400 rounded-2xl blur-2xl opacity-30 group-hover:opacity-50 transition-opacity" />
              <img 
                src="https://www.dropbox.com/s/hllo2ueo8zgi2tt/report.svg?raw=1" 
                alt="Reporting"
                className="relative rounded-2xl shadow-xl"
              />
            </div>
          </div>
          <div className="w-full lg:w-1/2 lg:pl-12">
            <div className="bg-gradient-to-br from-purple-50 to-pink-50 dark:from-purple-950 dark:to-pink-950 p-8 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-purple-100 dark:border-purple-800">
              <div className="inline-block p-3 bg-purple-600 dark:bg-purple-500 rounded-xl mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
              </div>
              <h4 className="text-3xl lg:text-4xl text-gray-900 dark:text-white font-bold mb-4">Reporting</h4>
              <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed">
                Our shortlink service can generate a comprehensive report on your vitals depending on your settings either daily, weekly, monthly, quarterly or yearly.
              </p>
            </div>
          </div>
        </div>

        {/* Feature 3 - Syncing */}
        <div className="flex items-center flex-wrap group">
          <div className="w-full lg:w-1/2 lg:pr-12 mb-8 lg:mb-0">
            <div className="bg-gradient-to-br from-pink-50 to-rose-50 dark:from-pink-950 dark:to-rose-950 p-8 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-pink-100 dark:border-pink-800">
              <div className="inline-block p-3 bg-pink-600 dark:bg-pink-500 rounded-xl mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                </svg>
              </div>
              <h4 className="text-3xl lg:text-4xl text-gray-900 dark:text-white font-bold mb-4">Syncing</h4>
              <p className="text-lg text-gray-700 dark:text-gray-300 leading-relaxed">
                Our shortlink service allows you to sync data across all your mobile devices whether iOS, Android or Windows OS and also to your laptop whether MacOS, GNU/Linux or Windows OS. This allows you to access your data from anywhere at any time.
              </p>
            </div>
          </div>
          <div className="w-full lg:w-1/2 flex justify-center">
            <div className="relative group-hover:scale-105 transition-transform duration-300">
              <div className="absolute inset-0 bg-gradient-to-r from-pink-400 to-rose-400 rounded-2xl blur-2xl opacity-30 group-hover:opacity-50 transition-opacity" />
              <img 
                src="https://www.dropbox.com/s/v0x0ywlvgmw04z6/sync.svg?raw=1" 
                alt="Syncing"
                className="relative rounded-2xl shadow-xl"
              />
            </div>
          </div>
        </div>
      </section>

      {/* Why Choose Us Section */}
      <div className="py-16 lg:py-24 bg-gradient-to-br from-slate-50 via-white to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-indigo-950">
        <section className="mx-auto container px-4 sm:px-6">
          <div className="text-center mb-16">
            <p className="uppercase text-sm font-semibold text-indigo-600 dark:text-indigo-400 tracking-wide mb-3">
              in few easy steps
            </p>
            <h1 className="text-3xl sm:text-4xl lg:text-5xl font-extrabold bg-gradient-to-r from-gray-900 to-gray-700 dark:from-white dark:to-gray-300 bg-clip-text text-transparent leading-tight max-w-4xl mx-auto">
              <Balancer>Create Beautiful Short Links &amp; Use a Powerful Link Management</Balancer>
            </h1>
          </div>

          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8 px-4">
            {/* Card 1 */}
            <div className="group flex flex-col items-center p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
              <div className="relative mb-6">
                <div className="absolute inset-0 bg-indigo-200 dark:bg-indigo-800 rounded-2xl rotate-6 group-hover:rotate-12 transition-transform" />
                <div className="relative bg-gradient-to-br from-indigo-600 to-indigo-700 dark:from-indigo-500 dark:to-indigo-600 rounded-2xl w-20 h-20 flex items-center justify-center shadow-xl">
                  <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG1.svg" alt="drawer" className="w-10 h-10" />
                </div>
              </div>
              <h4 className="text-lg font-bold text-center text-gray-900 dark:text-white leading-snug">
                We're a friendly, open source project that respects your privacy
              </h4>
            </div>

            {/* Card 2 */}
            <div className="group flex flex-col items-center p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
              <div className="relative mb-6">
                <div className="absolute inset-0 bg-purple-200 dark:bg-purple-800 rounded-2xl rotate-6 group-hover:rotate-12 transition-transform" />
                <div className="relative bg-gradient-to-br from-purple-600 to-purple-700 dark:from-purple-500 dark:to-purple-600 rounded-2xl w-20 h-20 flex items-center justify-center shadow-xl">
                  <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG2.svg" alt="check" className="w-10 h-10" />
                </div>
              </div>
              <h4 className="text-lg font-bold text-center text-gray-900 dark:text-white leading-snug">
                Coded by Developers for Developers
              </h4>
            </div>

            {/* Card 3 */}
            <div className="group flex flex-col items-center p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
              <div className="relative mb-6">
                <div className="absolute inset-0 bg-pink-200 dark:bg-pink-800 rounded-2xl rotate-6 group-hover:rotate-12 transition-transform" />
                <div className="relative bg-gradient-to-br from-pink-600 to-pink-700 dark:from-pink-500 dark:to-pink-600 rounded-2xl w-20 h-20 flex items-center justify-center shadow-xl">
                  <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG4.svg" alt="monitor" className="w-10 h-10" />
                </div>
              </div>
              <h4 className="text-lg font-bold text-center text-gray-900 dark:text-white leading-snug">
                We use cutting edge technologies to provide you with the best possible experience
              </h4>
            </div>

            {/* Card 4 */}
            <div className="group flex flex-col items-center p-6 bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
              <div className="relative mb-6">
                <div className="absolute inset-0 bg-rose-200 dark:bg-rose-800 rounded-2xl rotate-6 group-hover:rotate-12 transition-transform" />
                <div className="relative bg-gradient-to-br from-rose-600 to-rose-700 dark:from-rose-500 dark:to-rose-600 rounded-2xl w-20 h-20 flex items-center justify-center shadow-xl">
                  <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG3.svg" alt="html tag" className="w-10 h-10" />
                </div>
              </div>
              <h4 className="text-lg font-bold text-center text-gray-900 dark:text-white leading-snug">
                High Quality UI you can rely on us
              </h4>
            </div>
          </div>
        </section>
      </div>

      {/* Services Section */}
      <div className="container mx-auto px-4 sm:px-6 py-16 lg:py-24">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {/* Service 1 - Website Design */}
          <div className="group bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 overflow-hidden border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
            <div className="p-8">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-2xl shadow-lg mb-6 group-hover:scale-110 transition-transform">
                <svg className="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" />
                </svg>
              </div>
              <h3 className="text-2xl font-bold text-gray-900 dark:text-white mb-4 group-hover:text-indigo-600 dark:group-hover:text-indigo-400 transition-colors">
                Website Design
              </h3>
              <p className="text-gray-600 dark:text-gray-300 leading-relaxed">
                Encompassing today's website design technology to integrated and build solutions relevant to your business.
              </p>
            </div>
            <div className="h-2 bg-gradient-to-r from-indigo-500 to-purple-600 transform scale-x-0 group-hover:scale-x-100 transition-transform origin-left" />
          </div>

          {/* Service 2 - Branding */}
          <div className="group bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 overflow-hidden border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
            <div className="p-8">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-purple-500 to-pink-600 rounded-2xl shadow-lg mb-6 group-hover:scale-110 transition-transform">
                <svg className="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M20 6h-2.18c.11-.31.18-.65.18-1a2.996 2.996 0 00-5.5-1.65l-.5.67-.5-.68C10.96 2.54 10.05 2 9 2 7.34 2 6 3.34 6 5c0 .35.07.69.18 1H4c-1.11 0-1.99.89-1.99 2L2 19c0 1.11.89 2 2 2h16c1.11 0 2-.89 2-2V8c0-1.11-.89-2-2-2zm-5-2c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1zM9 4c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1z" />
                </svg>
              </div>
              <h3 className="text-2xl font-bold text-gray-900 dark:text-white mb-4 group-hover:text-purple-600 dark:group-hover:text-purple-400 transition-colors">
                Branding
              </h3>
              <p className="text-gray-600 dark:text-gray-300 leading-relaxed">
                Share relevant, engaging, and inspirational brand messages to create a connection with your audience.
              </p>
            </div>
            <div className="h-2 bg-gradient-to-r from-purple-500 to-pink-600 transform scale-x-0 group-hover:scale-x-100 transition-transform origin-left" />
          </div>

          {/* Service 3 - SEO Marketing */}
          <div className="group bg-white dark:bg-gray-800 rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 overflow-hidden border border-gray-100 dark:border-gray-700 hover:-translate-y-2">
            <div className="p-8">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-pink-500 to-rose-600 rounded-2xl shadow-lg mb-6 group-hover:scale-110 transition-transform">
                <svg className="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" />
                </svg>
              </div>
              <h3 className="text-2xl font-bold text-gray-900 dark:text-white mb-4 group-hover:text-pink-600 dark:group-hover:text-pink-400 transition-colors">
                SEO Marketing
              </h3>
              <p className="text-gray-600 dark:text-gray-300 leading-relaxed">
                Let us help you level up your search engine game, explore our solutions for digital marketing for your business.
              </p>
            </div>
            <div className="h-2 bg-gradient-to-r from-pink-500 to-rose-600 transform scale-x-0 group-hover:scale-x-100 transition-transform origin-left" />
          </div>
        </div>
      </div>
    </>
  )
}
