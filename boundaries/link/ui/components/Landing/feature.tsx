import * as React from 'react'
import Balancer from 'react-wrap-balancer'

export default function Feature() {
  return (
    <>
      <section className="container mx-auto px-6 p-10 dark:bg-gray-900">
        <h2 className="text-4xl font-bold text-center text-gray-800 dark:text-white p-10">Features</h2>
        <div className="flex items-center flex-wrap mb-20">
          <div className="w-full md:w-1/2">
            <h4 className="text-3xl text-gray-800 font-bold mb-3">URL Shortener</h4>
            <p className="text-gray-600 mb-8">
              Free custom URL Shortener with many features that gives you better quality for links shortening. We do not display ads during
              direct redirecting to the original url.
            </p>
          </div>
          <div className="w-full md:w-1/2">
            <img src="https://www.dropbox.com/s/mimcvn6zxtoruis/health.svg?raw=1" alt="Monitoring" />
          </div>
        </div>
        <div className="flex items-center flex-wrap mb-20">
          <div className="w-full md:w-1/2">
            <img src="https://www.dropbox.com/s/hllo2ueo8zgi2tt/report.svg?raw=1" alt="Reporting" />
          </div>
          <div className="w-full md:w-1/2 pl-10">
            <h4 className="text-3xl text-gray-800 font-bold mb-3">Reporting</h4>
            <p className="text-gray-600 mb-8">
              Our shortlink service can generate a comprehensive report on your vitals depending on your settings either daily, weekly,
              monthly, quarterly or yearly.
            </p>
          </div>
        </div>
        <div className="flex items-center flex-wrap mb-20">
          <div className="w-full md:w-1/2">
            <h4 className="text-3xl text-gray-800 font-bold mb-3">Syncing</h4>
            <p className="text-gray-600 mb-8">
              Our shortlink service allows you to sync data across all your mobile devices whether iOS, Android or Windows OS and also to
              your laptop whether MacOS, GNU/LInux or Windows OS. This allows you to access your data from anywhere at any time.
            </p>
          </div>
          <div className="w-full md:w-1/2">
            <img src="https://www.dropbox.com/s/v0x0ywlvgmw04z6/sync.svg?raw=1" alt="Syncing" />
          </div>
        </div>
      </section>

      <div className="py-12 bg-white dark:bg-gray-800 rounded my-3">
        <section className="mx-auto container bg-white dark:bg-gray-800 pt-16">
          <div className="px-4 lg:px-0">
            <div role="contentinfo" className="flex items-center flex-col px-4">
              <p className="focus:outline-none uppercase text-sm text-center text-gray-500 dark:text-gray-300 leading-none">
                in few easy steps
              </p>
              <h1 className="focus:outline-none text-4xl lg:text-4xl pt-4 font-extrabold text-center leading-tight text-gray-800 dark:text-white lg:w-7/12 md:w-9/12 xl:w-5/12">
                <Balancer>Create Beautiful Short Links &amp; Use a Powerful Link Management</Balancer>
              </h1>
            </div>
          </div>
        </section>

        <div className="pt-16">
          <div className="max-w-8xl mx-auto container">
            <div aria-label="group of cards" className="focus:outline-none flex flex-wrap items-center justify-center sm:justify-between">
              <div aria-label="card 1" className="focus:outline-none flex flex-col items-center py-6 md:py-0 px-6 w-full sm:w-1/2 md:w-1/4">
                <div className="w-20 h-20 relative ml-6">
                  <div className="absolute top-0 right-0 bg-indigo-100 rounded w-16 h-16 mt-2 mr-1" />
                  <div className="text-white absolute bottom-0 left-0 bg-indigo-700 rounded w-16 h-16 flex items-center justify-center mt-2 mr-3">
                    <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG1.svg" alt="drawer" />
                  </div>
                </div>
                <h4 className="focus:outline-none text-lg font-medium leading-6 text-gray-800 text-center pt-5">
                  We're a friendly, open source project that respects your privacy
                </h4>
              </div>
              <div aria-label="card 2" className="focus:outline-none flex flex-col items-center py-6 md:py-0 px-6 w-full sm:w-1/2 md:w-1/4">
                <div className="w-20 h-20 relative ml-6">
                  <div className="absolute top-0 right-0 bg-indigo-100 rounded w-16 h-16 mt-2 mr-1" />
                  <div className="text-white absolute bottom-0 left-0 bg-indigo-700 rounded w-16 h-16 flex items-center justify-center mt-2 mr-3">
                    <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG2.svg" alt="check" />
                  </div>
                </div>
                <h4 className="focus:outline-none text-lg font-medium leading-6 text-gray-800 text-center pt-5">
                  Coded by Developers
                  <br />
                  for Developers
                </h4>
              </div>
              <div aria-label="card 4" className="focus:outline-none flex flex-col items-center py-6 md:py-0 px-6 w-full sm:w-1/2 md:w-1/4">
                <div className="w-20 h-20 relative ml-6">
                  <div className="absolute top-0 right-0 bg-indigo-100 rounded w-16 h-16 mt-2 mr-1" />
                  <div className="text-white absolute bottom-0 left-0 bg-indigo-700 rounded w-16 h-16 flex items-center justify-center mt-2 mr-3">
                    <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG4.svg" alt="monitor" />
                  </div>
                </div>
                <h4 className="focus:outline-none text-lg font-medium leading-6 text-gray-800 text-center pt-5">
                  We use cutting edge technologies to provide you with the best possible experience
                </h4>
              </div>
              <div aria-label="card 3" className="focus:outline-none flex flex-col items-center py-6 md:py-0 px-6 w-full sm:w-1/2 md:w-1/4">
                <div className="w-20 h-20 relative ml-6">
                  <div className="absolute top-0 right-0 bg-indigo-100 rounded w-16 h-16 mt-2 mr-1" />
                  <div className="text-white absolute bottom-0 left-0 bg-indigo-700 rounded w-16 h-16 flex items-center justify-center mt-2 mr-3">
                    <img src="https://tuk-cdn.s3.amazonaws.com/can-uploader/compact_heading_with_icon-SVG3.svg" alt="html tag" />
                  </div>
                </div>
                <h4 className="focus:outline-none text-lg font-medium leading-6 text-gray-800 text-center pt-5">
                  Hight Quality UI
                  <br />
                  you can rely on us
                </h4>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="sm:flex flex-wrap justify-center items-center text-center gap-8">
        <div className="w-full sm:w-1/2 md:w-1/2 lg:w-1/4 px-4 py-4 bg-white dark:bg-gray-800 mt-6  shadow-lg rounded-lg">
          <div className="flex-shrink-0">
            <div className="flex items-center mx-auto justify-center h-12 w-12 rounded-md bg-indigo-500 text-white">
              <svg
                width="20"
                height="20"
                fill="currentColor"
                className="h-6 w-6"
                viewBox="0 0 1792 1792"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path d="M491 1536l91-91-235-235-91 91v107h128v128h107zm523-928q0-22-22-22-10 0-17 7l-542 542q-7 7-7 17 0 22 22 22 10 0 17-7l542-542q7-7 7-17zm-54-192l416 416-832 832h-416v-416zm683 96q0 53-37 90l-166 166-416-416 166-165q36-38 90-38 53 0 91 38l235 234q37 39 37 91z" />
              </svg>
            </div>
          </div>
          <h3 className="text-2xl sm:text-xl text-gray-700 font-semibold dark:text-white py-4">Website Design</h3>
          <p className="text-md  text-gray-500 dark:text-gray-300 py-4">
            Encompassing today’s website design technology to integrated and build solutions relevant to your business.
          </p>
        </div>

        <div className="w-full sm:w-1/2 md:w-1/2 lg:w-1/4 px-4 py-4 mt-6 sm:mt-16 md:mt-20 lg:mt-24 bg-white dark:bg-gray-800 shadow-lg rounded-lg">
          <div className="flex-shrink-0">
            <div className="flex items-center mx-auto justify-center h-12 w-12 rounded-md bg-indigo-500 text-white">
              <svg
                width="20"
                height="20"
                fill="currentColor"
                className="h-6 w-6"
                viewBox="0 0 1792 1792"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path d="M491 1536l91-91-235-235-91 91v107h128v128h107zm523-928q0-22-22-22-10 0-17 7l-542 542q-7 7-7 17 0 22 22 22 10 0 17-7l542-542q7-7 7-17zm-54-192l416 416-832 832h-416v-416zm683 96q0 53-37 90l-166 166-416-416 166-165q36-38 90-38 53 0 91 38l235 234q37 39 37 91z" />
              </svg>
            </div>
          </div>
          <h3 className="text-2xl sm:text-xl text-gray-700 font-semibold dark:text-white py-4">Branding</h3>
          <p className="text-md text-gray-500 dark:text-gray-300 py-4">
            Share relevant, engaging, and inspirational brand messages to create a connection with your audience.
          </p>
        </div>

        <div className="w-full sm:w-1/2 md:w-1/2 lg:w-1/4 mt-6  px-4 py-4 bg-white dark:bg-gray-800 shadow-lg rounded-lg">
          <div className="flex-shrink-0">
            <div className="flex items-center mx-auto justify-center h-12 w-12 rounded-md bg-indigo-500 text-white">
              <svg
                width="20"
                height="20"
                fill="currentColor"
                className="h-6 w-6"
                viewBox="0 0 1792 1792"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path d="M491 1536l91-91-235-235-91 91v107h128v128h107zm523-928q0-22-22-22-10 0-17 7l-542 542q-7 7-7 17 0 22 22 22 10 0 17-7l542-542q7-7 7-17zm-54-192l416 416-832 832h-416v-416zm683 96q0 53-37 90l-166 166-416-416 166-165q36-38 90-38 53 0 91 38l235 234q37 39 37 91z" />
              </svg>
            </div>
          </div>
          <h3 className="text-2xl sm:text-xl text-gray-700 font-semibold dark:text-white py-4">SEO Marketing</h3>
          <p className="text-md  text-gray-500 dark:text-gray-300 py-4">
            Let us help you level up your search engine game, explore our solutions for digital marketing for your business.
          </p>
        </div>
      </div>
    </>
  )
}
