import Balancer from 'react-wrap-balancer'

export default function Header() {
  return (
    <div className="relative bg-gradient-to-br from-slate-50 to-gray-100 dark:bg-gradient-to-br dark:from-gray-900 dark:to-gray-800 rounded-lg cq-hero">
      <div className="cq-hero__grid mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="cq-hero__content relative z-10">
          <svg
            className="cq-hero__divider absolute right-0 inset-y-0 h-full w-48 text-gray-50 dark:text-gray-900 transform translate-x-1/2"
            fill="currentColor"
            viewBox="0 0 100 100"
            preserveAspectRatio="none"
            aria-hidden="true"
          >
            <polygon points="50,0 100,0 50,100 0,100" />
          </svg>

          <main className="cq-hero__body">
            <h1 className="cq-hero__title tracking-tight font-extrabold text-gray-900 dark:text-gray-100">
              <Balancer>
                <span className="block xl:inline">Shorten your links</span>{' '}
                <span className="block text-indigo-600 xl:inline dark:text-indigo-300">in one click</span>
              </Balancer>
            </h1>
            <p className="cq-hero__lead mt-3 text-gray-700 dark:text-gray-300">
              Get full control over your links in one place. <br />
              Easily create, manage, and track your links. <br />
              Get started today!
            </p>
            <div className="mt-6 cq-hero__actions">
              <a
                href="#"
                className="group inline-flex items-center justify-center px-6 py-3.5 font-semibold text-white bg-purple-700 dark:bg-indigo-500 rounded-xl shadow-lg hover:shadow-xl hover:bg-purple-800 dark:hover:bg-indigo-600 transform hover:-translate-y-0.5 transition-all duration-200 sm:px-8 md:py-4 md:text-lg lg:px-10"
              >
                <span className="flex items-center gap-2">
                  Get started
                  <svg className="w-5 h-5 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 7l5 5m0 0l-5 5m5-5H6" />
                  </svg>
                </span>
              </a>
              <a
                href="#"
                className="group inline-flex items-center justify-center px-6 py-3.5 font-semibold text-indigo-700 dark:text-indigo-300 bg-white dark:bg-gray-800 border-2 border-indigo-600 dark:border-indigo-400 rounded-xl shadow-md hover:shadow-lg hover:bg-indigo-50 dark:hover:bg-gray-700 transform hover:-translate-y-0.5 transition-all duration-200 sm:px-8 md:py-4 md:text-lg lg:px-10"
              >
                <span className="flex items-center gap-2">
                  Learn more
                  <svg className="w-5 h-5 group-hover:scale-110 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
                  </svg>
                </span>
              </a>
            </div>
          </main>
        </div>
        <div className="cq-hero__media">
          <img
            className="cq-hero__image"
            src="https://images.unsplash.com/photo-1551434678-e076c223a692?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2850&q=80"
            alt=""
          />
        </div>
      </div>
    </div>
  )
}
