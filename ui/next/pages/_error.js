import React from 'react'
import { Layout } from 'components'
import { UndrawBackInTheDay } from 'react-undraw-illustrations'

export default function Error() {
  return (
    <Layout>
      <div className="min-w-screen bg-blue-100 flex items-center p-5 lg:p-20 overflow-hidden relative rounded-md">
        <div className="flex-1 min-h-full min-w-full rounded-3xl bg-white shadow-xl p-10 lg:p-20 text-gray-800 relative md:flex items-center text-center md:text-left">
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
          <div className="w-full md:w-1/2 text-center">
            <UndrawBackInTheDay primaryColor="#6c68fb" height="250px" />
          </div>
        </div>
        <div className="w-64 md:w-96 h-96 md:h-full bg-blue-200 bg-opacity-30 absolute -top-64 md:-top-96 right-20 md:right-32 rounded-full pointer-events-none -rotate-45 transform" />
        <div className="w-96 h-full bg-yellow-200 bg-opacity-20 absolute -bottom-96 right-64 rounded-full pointer-events-none -rotate-45 transform" />
      </div>
    </Layout>
  )
}

// @ts-ignore
Error.getInitialProps = ({ res, err }) => {
  const statusCode = (res && err && err.statusCode) || 404
  return { statusCode }
}

function Error404() {
    return (
        <div className="flex flex-col items-center justify-center py-24 lg:py-12 md:px-16 px-4">
            <h1 className="text-7xl font-bold text-indigo-700 pb-2">404</h1>
            <h2 className="lg:text-5xl md:text-4xl text-2xl font-bold text-gray-800 py-2">This Page Does Not Exist</h2>
            <p className="text-base text-gray-600 py-2 text-center">Sorry! We could not find you the page you are looking for. Please check URL in address bar and try again.</p>
            <div className="flex md:flex-row flex-col items-center justify-center md:gap-8 mt-4 mb-12 w-full">
                <button className="p-4 text-base text-center text-white md:w-auto md:mb-0 mb-4 w-full bg-indigo-700 border rounded-md hover:bg-indigo-800">Get back to Homepage</button>
                <button className="p-4 text-base font-semibold text-center md:w-auto w-full bg-gray-100 text-indigo-700 border rounded-md hover:bg-gray-200">Contact Support</button>
            </div>
            <div className="hidden md:grid place-content-center lg:w-1/3 w-1/2">
                <img src="https://i.ibb.co/JjmY1tm/tuk-component.png" alt="girl in an underconstruction site" />
            </div>
            <div className="md:hidden grid place-content-center">
                <img src="https://i.ibb.co/zxQ6hyF/undraw-warning-cyit-1-1.png" alt="girl in an underconstruction site" />
            </div>
        </div>

    );
}
