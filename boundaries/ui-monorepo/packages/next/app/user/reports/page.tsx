'use client'

// @ts-ignore
import { Header } from '@shortlink-org/ui-kit'

import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'

function Page() {
  return (
    <Layout>
      {/*<NextSeo title="Reports" description="Reports page for your account." />*/}

      <Header title="Reports" />

      <div className="px-4 py-4 my-3 rounded mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-14 lg:px-8 lg:py-10 bg-white dark:bg-gray-800">
        <div className="flex flex-col">
          <p className="text-gray-800 dark:text-gray-300">
            Reporting is a critical part of our shortlink service. Depending on your settings, we can generate a comprehensive report on
            your vitals either daily, weekly, monthly, quarterly or yearly. This report can include things like how many people clicked on
            your links, where they came from, and what kind of device they were using. This information can be extremely valuable in
            understanding your audience and tailoring your content to them. Additionally, we can customize the reports to include only the
            information that you are interested in. Our reporting system is designed to be flexible and user-friendly, so you can get the
            most out of it.
          </p>

          <br />

          <p className="text-gray-800 dark:text-gray-300">
            Whether you're looking to track your progress over time or want to stay informed about the latest changes in your industry, our
            reporting feature will give you the information you need to make decisions that improve your bottom line.
          </p>
        </div>
      </div>

      <div className="w-full flex items-center justify-center">
        <div className="py-4 sm:py-6 md:py-8 bg-white dark:bg-gray-800 shadow rounded-lg">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 px-6 xl:px-10 gap-y-8 gap-x-12 2xl:gap-x-28">
            <div className="w-full">
              <p className="text-xs md:text-sm font-medium leading-none text-gray-500 dark:text-gray-400 uppercase">Sales</p>
              <p className="text-lg sm:text-xl md:text-2xl lg:text-3xl font-bold leading-3 text-gray-800 dark:text-gray-200 mt-3 md:mt-5">
                89.5%
              </p>
              <div className="flex flex-col md:w-64">
                <div className="w-full flex justify-end">
                  <div className="flex items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" width={16} height={16} viewBox="0 0 16 16" fill="none">
                      <path d="M8 3.33334V12.6667" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M12 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M4 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                    <p className="text-xs leading-none text-green-600">4.3%</p>
                  </div>
                </div>
                <div className="mt-2.5">
                  <div className="w-full h-1 bg-gray-200 rounded-full">
                    <div className="w-1/2 h-1 bg-blue-500 rounded-full" />
                  </div>
                </div>
              </div>
              <p className="mt-1.5 text-xs leading-3 text-gray-400 dark:text-gray-300">Yearly target</p>
            </div>
            <div className="w-full">
              <p className="text-xs md:text-sm font-medium leading-none text-gray-500 uppercase">revenue</p>
              <p className="text-lg sm:text-xl md:text-2xl lg:text-3xl font-bold leading-3 text-gray-800 dark:text-gray-200 mt-3 md:mt-5">
                $75,000
              </p>
              <div className="flex flex-col">
                <div className="h-4" />
                <div className="md:w-64 mt-2.5">
                  <div className="w-full h-1 bg-gray-200 rounded-full">
                    <div className="w-40 h-1 bg-lime-500 rounded-full" />
                  </div>
                </div>
              </div>
              <p className="mt-1.5 text-xs leading-3 text-gray-400 dark:text-gray-300">Yearly target</p>
            </div>
            <div className="w-full">
              <p className="text-xs md:text-sm font-medium leading-none text-gray-500 uppercase">customers</p>
              <p className="text-lg sm:text-xl md:text-2xl lg:text-3xl font-bold leading-3 text-gray-800 dark:text-gray-200 mt-3 md:mt-5">
                3922
              </p>
              <div className="flex flex-col md:w-64">
                <div className="w-full flex justify-end">
                  <div className="flex items-center">
                    <svg xmlns="http://www.w3.org/2000/svg" width={16} height={16} viewBox="0 0 16 16" fill="none">
                      <path d="M8 3.33334V12.6667" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M12 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                      <path d="M4 7.33334L8 3.33334" stroke="#16A34A" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                    <p className="text-xs leading-none text-green-600">9.1%</p>
                  </div>
                </div>
                <div className="mt-2.5">
                  <div className="w-full h-1 bg-gray-200 rounded-full">
                    <div className="w-44 h-1 bg-yellow-500 rounded-full" />
                  </div>
                </div>
              </div>
              <p className="mt-1.5 text-xs leading-3 text-gray-400 dark:text-gray-300">Yearly target</p>
            </div>
          </div>
        </div>
      </div>

      <Ready />
    </Layout>
  )
}

export default withAuthSync(Page)
