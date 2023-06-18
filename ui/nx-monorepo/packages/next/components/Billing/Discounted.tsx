export default function Discounted() {
  return (
    <div className="mt-4 my-4 relative max-w-screen-xl mx-auto px-4 sm:px-6 lg:px-8 lg:mt-5">
      <div className="max-w-md mx-auto lg:max-w-5xl ">
        <div className="rounded-lg px-6 py-8 sm:p-10 lg:flex lg:items-center bg-gray-800 hover:bg-gray-700">
          <div className="flex-1">
            <span className="inline-flex dark:text-gray-700 px-4 py-1 rounded-full text-sm leading-5 font-semibold tracking-wide uppercase bg-white dark:bg-darkmodebg">
              Discounted
            </span>
            <div className="mt-4 text-lg leading-7 text-gray-500 dark:text-gray-300">
              Get full access to all of standard license features for solo
              projects that make less than $20k gross revenue for &nbsp;
              <strong className="font-semibold text-white dark:text-gray-200">
                $29
              </strong>
              .
            </div>
          </div>
          <div className="mt-6 rounded-md shadow lg:mt-0 lg:ml-10 lg:flex-shrink-0">
            <a
              href="#"
              className="flex items-center justify-center px-5 py-3 border border-transparent text-base leading-6 font-medium rounded-md text-gray-900 bg-white hover:text-gray-700 focus:outline-none focus:shadow-outline transition duration-150 ease-in-out"
            >
              Buy Discounted License
            </a>
          </div>
        </div>
      </div>
    </div>
  )
}
