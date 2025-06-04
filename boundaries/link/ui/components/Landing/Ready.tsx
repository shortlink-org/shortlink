import * as React from 'react'

export default function Ready() {
  return (
    <div className="bg-gray-50 dark:bg-gray-800 rounded my-5">
      <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:py-16 lg:px-8 lg:flex lg:items-center lg:justify-between">
        <h2 className="text-3xl font-extrabold tracking-tight text-gray-900 dark:text-white sm:text-4xl">
          <span className="block">Ready to dive in?</span>
          <span className="block text-indigo-600 dark:text-indigo-400">Start your free trial today.</span>
        </h2>
        <div className="mt-8 flex lg:mt-0 lg:flex-shrink-0">
          <div className="inline-flex rounded-md shadow">
            <a
              href="#"
              className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
            >
              Get started
            </a>
          </div>
          <div className="ml-3 inline-flex rounded-md shadow">
            <a
              href="#"
              className="inline-flex items-center justify-center px-5 py-3 border border-transparent text-base font-medium rounded-md text-indigo-600 bg-white dark:bg-gray-800 dark:text-gray-300 hover:bg-indigo-50 dark:hover:bg-gray-700"
            >
              Learn more
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}
