export default function Subscribe() {
  return (
    <header className="bg-gray-50 dark:bg-gray-800 rounded border-t-4 border-indigo-500 cq-subscribe">
      <div className="container px-6 py-16 mx-auto rounded cq-subscribe__grid">
        <div className="cq-subscribe__content">
          <div className="max-w-lg">
            <h1 className="cq-subscribe__title font-semibold text-gray-800 dark:text-white">
              Subscribe To The <span className="text-indigo-500">Newsletter</span>
            </h1>

            <p className="mt-4 text-gray-600 dark:text-gray-400">
              be the first to knows when our <span className="font-medium text-indigo-500">Brand</span> is live
            </p>

            <div className="cq-subscribe__form mt-8">
              <div className="relative">
                <div className="absolute text-gray-600 dark:text-gray-400 flex items-center pl-4 h-full cursor-pointer">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="icon icon-tabler icon-tabler-mail"
                    width={18}
                    height={18}
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    fill="none"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path stroke="none" d="M0 0h24v24H0z" />
                    <rect x="3" y="5" width="18" height="14" rx="2" />
                    <polyline points="3 7 12 13 21 7" />
                  </svg>
                </div>
                <input
                  id="email"
                  type="email"
                  className="cq-subscribe__input px-4 py-2 text-gray-700 bg-white dark:bg-gray-900 border-2 border-gray-300 h-10 flex items-center pl-12 rounded-lg dark:text-gray-200 dark:border-gray-600 focus:border-indigo-500 dark:focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-200 dark:focus:ring-indigo-800 transition-all duration-200"
                  placeholder="Enter your email"
                  required
                />
              </div>

              <button
                type="button"
                className="cq-subscribe__button px-6 py-2.5 text-sm font-semibold tracking-wide text-white capitalize transition-all duration-200 transform bg-purple-700 dark:bg-indigo-500 rounded-lg hover:bg-purple-800 dark:hover:bg-indigo-600 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 shadow-md hover:shadow-lg active:scale-95"
              >
                Subscribe
              </button>
            </div>
          </div>
        </div>

        <div className="cq-subscribe__media w-full p-3">
          <img 
            src="/assets/images/undraw_designer_re_5v95.svg" 
            alt="Designer illustration" 
            className="w-full h-auto max-w-md" 
          />
        </div>
      </div>
    </header>
  )
}
