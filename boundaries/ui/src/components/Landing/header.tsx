import Balancer from 'react-wrap-balancer'

export default function Header() {
  return (
    <div className="relative bg-gradient-to-br from-slate-50 to-gray-100 dark:bg-gradient-to-br dark:from-gray-900 dark:to-gray-800 rounded-lg">
      <div className="max-w-7xl mx-auto rounded-lg">
        <div className="relative pt-1 px-3 z-10 pb-8 bg-gradient-to-br from-slate-50 to-gray-100 dark:bg-gradient-to-br dark:from-gray-900 dark:to-gray-800 sm:pb-16 md:pb-20 lg:max-w-2xl lg:w-full lg:pb-28 xl:pb-32">
          <svg
            className="hidden lg:block absolute right-0 inset-y-0 h-full w-48 text-gray-50 dark:text-gray-900 transform translate-x-1/2"
            fill="currentColor"
            viewBox="0 0 100 100"
            preserveAspectRatio="none"
            aria-hidden="true"
          >
            <polygon points="50,0 100,0 50,100 0,100" />
          </svg>

          <main className="mx-auto max-w-7xl px-4 sm:mt-6 sm:px-6 md:mt-8 lg:mt-10 lg:px-8 xl:mt-14">
            <div className="sm:text-center lg:text-left">
              <h1 className="text-3xl tracking-tight font-extrabold text-gray-900 dark:text-gray-100 sm:text-4xl md:text-5xl lg:text-6xl">
                <Balancer>
                  <span className="block xl:inline">Shorten your links</span>{' '}
                  <span className="block text-indigo-600 xl:inline dark:text-indigo-300">in one click</span>
                </Balancer>
              </h1>
              <p className="mt-3 text-sm text-gray-700 dark:text-gray-300 sm:text-base sm:mt-5 sm:text-lg sm:max-w-xl sm:mx-auto md:mt-5 md:text-xl lg:mx-0">
                Get full control over your links in one place. <br />
                Easily create, manage, and track your links. <br />
                Get started today!
              </p>
              <div className="mt-5 sm:mt-8 flex flex-col sm:flex-row gap-3 sm:gap-4 sm:justify-center lg:justify-start">
                <a
                  href="#"
                  className="group relative w-full sm:w-auto inline-flex items-center justify-center px-6 py-3.5 overflow-hidden font-semibold text-white bg-gradient-to-r from-indigo-600 to-purple-600 dark:from-indigo-500 dark:to-purple-500 rounded-xl shadow-lg hover:shadow-2xl transform hover:-translate-y-0.5 transition-all duration-200 sm:px-8 md:py-4 md:text-lg lg:px-10"
                >
                  <span className="relative z-10 flex items-center gap-2">
                    Get started
                    <svg className="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
                    </svg>
                  </span>
                  <div className="absolute inset-0 bg-gradient-to-r from-indigo-700 to-purple-700 dark:from-indigo-600 dark:to-purple-600 opacity-0 group-hover:opacity-100 transition-opacity rounded-xl" />
                </a>
                <a
                  href="#"
                  className="group w-full sm:w-auto inline-flex items-center justify-center px-6 py-3.5 font-semibold text-indigo-700 dark:text-indigo-300 bg-white dark:bg-gray-800 border-2 border-indigo-600 dark:border-indigo-400 rounded-xl shadow-md hover:shadow-lg hover:bg-indigo-50 dark:hover:bg-gray-700 transform hover:-translate-y-0.5 transition-all duration-200 sm:px-8 md:py-4 md:text-lg lg:px-10"
                >
                  <span className="flex items-center gap-2">
                    Learn more
                    <svg className="w-5 h-5 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                    </svg>
                  </span>
                </a>
              </div>
            </div>
          </main>
        </div>
      </div>
      <div className="lg:absolute lg:inset-y-0 lg:right-0 lg:w-1/2">
        <img
          className="h-56 w-full object-cover sm:h-72 md:h-96 lg:w-full lg:h-full"
          src="https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80"
          alt=""
        />
      </div>
    </div>
  )
}
