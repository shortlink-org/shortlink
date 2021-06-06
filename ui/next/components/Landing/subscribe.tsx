import { UndrawDesigner } from 'react-undraw-illustrations';

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
                <input
                  id="email"
                  type="text"
                  className="px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-500 dark:focus:border-blue-500 focus:outline-none focus:ring"
                  placeholder="Email Address"
                />

                <button className="w-full px-4 py-2 text-sm font-medium tracking-wide text-white capitalize transition-colors duration-200 transform bg-indigo-500 rounded-lg lg:w-auto lg:mx-4 hover:bg-indigo-400 focus:outline-none focus:bg-indigo-400">
                  Subscribe
                </button>
              </div>
            </div>
          </div>

          <div className="flex items-center justify-center w-full mt-6 lg:mt-0 lg:w-1/2">
            <UndrawDesigner
              skinColor="brown"
              className="w-full h-full max-w-md"
              hairColor="#222"
            />
          </div>
        </div>
      </div>
    </header>
  );
}
