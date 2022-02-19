import React from 'react'

function Index() {
  return (
    <div className="w-full bg-gray-100 dark:bg-gray-900 py-10">
      <div className="container mx-auto px-6 flex items-start justify-center">
        <div className="w-full">
          {/* Card is full width. Use in 12 col grid for best view. */}
          {/* Card code block start */}
          <div className="flex flex-col lg:flex-row mx-auto bg-white dark:bg-gray-800 shadow rounded">
            <div className="w-full lg:w-1/3 px-12 flex flex-col items-center py-10">
              <div className="w-24 h-24 mb-3 p-2 rounded-full bg-gray-200 dark:bg-gray-700 flex items-center justify-center">
                <img
                  className="w-full h-full overflow-hidden object-cover rounded-full"
                  src="https://tuk-cdn.s3.amazonaws.com/assets/components/grid_cards/gc_2.png"
                  alt="avatar"
                />
              </div>
              <h2 className="text-gray-800 dark:text-gray-100 text-xl tracking-normal font-medium mb-1">
                Timothy Jon Oliphant
              </h2>
              <p className="flex text-gray-600 dark:text-gray-100 text-sm tracking-normal font-normal mb-3 text-center">
                <span className="cursor-pointer mr-1 text-gray-600 dark:text-gray-100">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="icon icon-tabler icon-tabler-map-pin"
                    width={20}
                    height={20}
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    fill="none"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path stroke="none" d="M0 0h24v24H0z" />
                    <circle cx={12} cy={11} r={3} />
                    <path d="M17.657 16.657L13.414 20.9a1.998 1.998 0 0 1 -2.827 0l-4.244-4.243a8 8 0 1 1 11.314 0z" />
                  </svg>
                </span>
                Las Vegas, Nevada
              </p>
              <p className="text-gray-600 dark:text-gray-100 text-sm tracking-normal font-normal mb-8 text-center w-10/12">
                The more I deal with the work as something that is my own, as
                something that is personal, the more successful it is.
              </p>
              <div className="flex items-start">
                <div>
                  <h2 className="text-gray-600 dark:text-gray-100 text-2xl leading-6 mb-2 text-center">
                    82
                  </h2>
                  <p className="text-gray-800 dark:text-gray-100 text-sm leading-5">
                    Reviews
                  </p>
                </div>
                <div className="mx-6 lg:mx-3 xl:mx-6 px-8 lg:px-4 xl:px-8 border-l border-r">
                  <h2 className="text-gray-600 dark:text-gray-100 text-2xl leading-6 mb-2 text-center">
                    28
                  </h2>
                  <p className="text-gray-800 dark:text-gray-100 text-sm leading-5">
                    Projects
                  </p>
                </div>
                <div>
                  <h2 className="text-gray-600 dark:text-gray-100 text-2xl leading-6 mb-2 text-center">
                    42
                  </h2>
                  <p className="text-gray-800 dark:text-gray-100 text-sm leading-5">
                    Approved
                  </p>
                </div>
              </div>
            </div>
            <div className="w-full lg:w-1/3 px-12 border-t border-b lg:border-t-0 lg:border-b-0 lg:border-l lg:border-r border-gray-300 flex flex-col items-center py-10">
              <div className="mb-3 w-24 h-24 rounded-full bg-gray-100 flex items-center justify-center cursor-pointer text-indigo-700">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="icon icon-tabler icon-tabler-stack"
                  width={48}
                  height={48}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  fill="none"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" d="M0 0h24v24H0z" />
                  <polyline points="12 4 4 8 12 12 20 8 12 4" />
                  <polyline points="4 12 12 16 20 12" />
                  <polyline points="4 16 12 20 20 16" />
                </svg>
              </div>
              <h2 className="text-gray-800 dark:text-gray-100 text-xl tracking-normal text-center font-medium mb-1">
                Senior Product Designer
              </h2>
              <p className="text-gray-600 dark:text-gray-100 text-sm tracking-normal font-normal mb-3 text-center">
                Freelance
              </p>
              <p className="text-gray-600 dark:text-gray-100 text-sm tracking-normal font-normal mb-8 text-center w-10/12">
                John is a true asset to us, providing advanced designing skills
                from years of experience as UX designer.
              </p>
              <div className="flex items-start">
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded text-xs leading-3 py-2 px-3">
                  Interface
                </div>
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded mx-2 text-xs leading-3 py-2 px-3">
                  Interface
                </div>
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded text-xs leading-3 py-2 px-3">
                  Interface
                </div>
              </div>
              <div className="mt-2 flex items-start">
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded text-xs leading-3 py-2 px-3">
                  Interface
                </div>
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded mx-2 text-xs leading-3 py-2 px-3">
                  Interface
                </div>
                <div className="bg-gray-200 text-gray-600 dark:text-gray-100 dark:bg-gray-700 rounded text-xs leading-3 py-2 px-3">
                  Interface
                </div>
              </div>
            </div>
            <div className="w-full lg:w-1/3 flex-col flex justify-center items-center px-12 py-8">
              <h2 className="text-center text-2xl text-gray-800 dark:text-gray-100 font-medium tracking-normal">
                $90
              </h2>
              <h2 className="text-center text-sm text-gray-600 dark:text-gray-100 font-normal mt-2 mb-4 tracking-normal">
                Hourly Rate
              </h2>
              <h2 className="text-center text-2xl text-gray-800 dark:text-gray-100 font-medium tracking-normal">
                $32,000
              </h2>
              <h2 className="text-center text-sm text-gray-600 dark:text-gray-100 font-normal mt-2 mb-4 tracking-normal">
                Total Earned
              </h2>
              <h2 className="text-center text-2xl text-gray-800 dark:text-gray-100 font-medium tracking-normal">
                2000
              </h2>
              <h2 className="text-center text-sm text-gray-600 dark:text-gray-100 font-normal mt-2 mb-4 tracking-normal">
                Hours Worked
              </h2>
              <h2 className="text-center text-2xl text-gray-800 dark:text-gray-100 font-medium tracking-normal">
                95%
              </h2>
              <h2 className="text-center text-sm text-gray-600 dark:text-gray-100 font-normal mt-2 mb-4 tracking-normal">
                Success Rate
              </h2>
              <div className="flex items-center">
                <svg
                  className="cursor-pointer w-4 mr-1 text-yellow-400 icon icon-tabler icon-tabler-star"
                  xmlns="http://www.w3.org/2000/svg"
                  width={20}
                  height={20}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" fill="none" d="M0 0h24v24H0z" />
                  <path
                    fill="currentColor"
                    d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z"
                  />
                </svg>
                <svg
                  className="cursor-pointer w-4 mr-1 text-yellow-400 icon icon-tabler icon-tabler-star"
                  xmlns="http://www.w3.org/2000/svg"
                  width={20}
                  height={20}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" fill="none" d="M0 0h24v24H0z" />
                  <path
                    fill="currentColor"
                    d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z"
                  />
                </svg>
                <svg
                  className="cursor-pointer w-4 mr-1 text-yellow-400 icon icon-tabler icon-tabler-star"
                  xmlns="http://www.w3.org/2000/svg"
                  width={20}
                  height={20}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" fill="none" d="M0 0h24v24H0z" />
                  <path
                    fill="currentColor"
                    d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z"
                  />
                </svg>
                <svg
                  className="cursor-pointer w-4 mr-1 text-yellow-400 icon icon-tabler icon-tabler-star"
                  xmlns="http://www.w3.org/2000/svg"
                  width={20}
                  height={20}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" fill="none" d="M0 0h24v24H0z" />
                  <path
                    fill="currentColor"
                    d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z"
                  />
                </svg>
                <svg
                  className="cursor-pointer w-4 text-gray-200 dark:text-gray-400 icon icon-tabler icon-tabler-star"
                  xmlns="http://www.w3.org/2000/svg"
                  width={20}
                  height={20}
                  viewBox="0 0 24 24"
                  strokeWidth="1.5"
                  stroke="currentColor"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                >
                  <path stroke="none" fill="none" d="M0 0h24v24H0z" />
                  <path
                    fill="currentColor"
                    d="M12 17.75l-6.172 3.245 1.179-6.873-4.993-4.867 6.9-1.002L12 2l3.086 6.253 6.9 1.002-4.993 4.867 1.179 6.873z"
                  />
                </svg>
              </div>
            </div>
          </div>
          {/* Card code block end */}
        </div>
      </div>
    </div>
  )
}
export default Index
