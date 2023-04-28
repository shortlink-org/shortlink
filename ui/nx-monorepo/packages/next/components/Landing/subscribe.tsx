// @ts-nocheck
import Image from 'next/image'
import UndrawDesigner from '../../public/assets/images/undraw_designer_re_5v95.svg'

export default function Subscribe() {
  return (
    <header className="bg-white dark:bg-gray-800 rounded border-t-4 border-indigo-500">
      <div className="container px-6 py-16 mx-auto rounded">
        <div className="items-center lg:flex">
          <div className="w-full lg:w-1/2">
            <div className="lg:max-w-lg">
              <h1 className="text-2xl font-semibold text-gray-800 dark:text-white lg:text-3xl">
                Subscribe To The{' '}
                <span className="text-indigo-500">Newsletter</span>
              </h1>

              <p className="mt-4 text-gray-600 dark:text-gray-400">
                be the first to knows when our{' '}
                <span className="font-medium text-indigo-500">Brand</span> is
                live
              </p>

              <div className="flex flex-col mt-8 space-y-3 lg:space-y-0 lg:flex-row">
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
                      <rect x={3} y={5} width={18} height={14} rx={2} />
                      <polyline points="3 7 12 13 21 7" />
                    </svg>
                  </div>
                  <input
                    id="email"
                    type="email"
                    className="px-4 py-2 text-gray-700 bg-white dark:bg-gray-800 border border-gray-300 w-64 h-10 flex items-center pl-12 rounded-lg dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring"
                    placeholder="Enter your email"
                    required
                  />
                </div>

                <button className="w-full px-4 py-2 text-sm font-medium tracking-wide text-white capitalize transition-colors duration-200 transform bg-indigo-500 rounded-lg lg:w-auto lg:mx-4 hover:bg-indigo-400 focus:outline-none focus:bg-indigo-400">
                  Subscribe
                </button>
              </div>
            </div>
          </div>

          <div className="flex items-center justify-center w-full mt-3 lg:mt-0 sm:w-1/2 p-3">
            <UndrawDesigner />
          </div>
        </div>
      </div>
    </header>
  )
}
