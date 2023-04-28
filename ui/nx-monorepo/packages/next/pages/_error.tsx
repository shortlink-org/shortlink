import React from 'react'
import { Layout } from 'components'

import UndrawBackInTheDay from '../public/assets/images/undraw_back_in_the_day_knsh.svg'
import { NextSeo } from 'next-seo'

export default function Error() {
  return (
    <Layout>
      <NextSeo
        title="Error"
        description="Error page for the shortlink service"
      />
      <div className="min-w-screen bg-blue-100 flex items-center p-5 lg:p-20 overflow-hidden relative rounded-md">
        <div className="flex-1 min-h-full min-w-full rounded-3xl bg-white dark:bg-gray-800 shadow-xl p-10 lg:p-20 text-gray-800 relative md:flex items-center text-center md:text-left">
          <div className="w-full md:w-1/2">
            <div className="mb-16 md:mb-16 text-gray-600 font-light">
              <h1 className="font-black uppercase text-3xl lg:text-5xl text-yellow-500 mb-10">
                You seem to be lost!
              </h1>
              <p>The page you're looking for isn't available.</p>
              <p>Try searching again or use the Go Back button below.</p>
            </div>
            <div className="mb-20 md:mb-0">
              <button className="text-lg font-light outline-none focus:outline-none transform transition-all hover:scale-110 text-yellow-500 hover:text-yellow-600">
                <i className="mdi mdi-arrow-left mr-2" />
                Go Back
              </button>
            </div>
          </div>
          <div className="w-full md:w-1/2 text-center p-5">
            <UndrawBackInTheDay />
          </div>
        </div>
        <div className="w-64 md:w-96 h-96 md:h-full bg-blue-200 bg-opacity-30 absolute -top-64 md:-top-96 right-20 md:right-32 rounded-full pointer-events-none -rotate-45 transform" />
        <div className="w-96 h-full bg-yellow-200 bg-opacity-20 absolute -bottom-96 right-64 rounded-full pointer-events-none -rotate-45 transform" />
      </div>
    </Layout>
  )
}
