'use client'

import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'
import Header from 'components/Page/Header'

function Page() {
  return (
    <>
      {/*<NextSeo title="Audit" description="Audit your account" />*/}

      <Header title="Audit" />

      <div className="px-4 py-4 my-3 rounded mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-14 lg:px-8 lg:py-10 bg-white dark:bg-gray-800">
        <div className="flex flex-col">
          <p className="text-gray-800">
            Keep your business safe with shortlink. We add an audit trail so you can see what's happening in your account in real-time, and
            review past activity whenever you need to.
          </p>
        </div>
      </div>

      <Ready />

      <ul className="p-4 lg:p-8 bg-coolGray-100 text-coolGray-800">
        <li>
          <article>
            <a
              href="#"
              className="grid p-4 overflow-hidden md:grid-cols-5 rounded-xl lg:p-6 xl:grid-cols-12 hover:bg-white dark:bg-gray-800"
            >
              <h3 className="mb-1 ml-8 font-semibold md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                Earum at ipsa aliquid quis, exercitationem est.
              </h3>
              <time dateTime="" className="row-start-1 mb-1 md:col-start-1 xl:col-span-2 text-coolGray-600">
                Oct 13, 2020
              </time>
              <p className="ml-8 md:col-start-2 md:col-span-4 xl:col-start-3 xl:col-span-9 md:ml-0 text-coolGray-700">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Similique saepe exercitationem numquam, labore necessitatibus
                deleniti quasi. Illo porro nihil necessitatibus debitis delectus aperiam, fuga impedit assumenda odit, velit eveniet est.
              </p>
            </a>
          </article>
        </li>
        <li>
          <article>
            <a
              href="#"
              className="grid p-4 overflow-hidden md:grid-cols-5 rounded-xl lg:p-6 xl:grid-cols-12 hover:bg-white dark:bg-gray-800"
            >
              <h3 className="mb-1 ml-8 font-semibold md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                Earum at ipsa aliquid quis, exercitationem est.
              </h3>
              <time dateTime="" className="row-start-1 mb-1 md:col-start-1 xl:col-span-2 text-coolGray-600">
                Oct 13, 2020
              </time>
              <p className="ml-8 md:col-start-2 md:col-span-4 xl:col-start-3 xl:col-span-9 md:ml-0 text-coolGray-700">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Similique saepe exercitationem numquam, labore necessitatibus
                deleniti quasi. Illo porro nihil necessitatibus debitis delectus aperiam, fuga impedit assumenda odit, velit eveniet est.
              </p>
            </a>
          </article>
        </li>
        <li>
          <article>
            <a
              href="#"
              className="grid p-4 overflow-hidden md:grid-cols-5 rounded-xl lg:p-6 xl:grid-cols-12 hover:bg-white dark:bg-gray-800"
            >
              <h3 className="mb-1 ml-8 font-semibold md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                Earum at ipsa aliquid quis, exercitationem est.
              </h3>
              <time dateTime="" className="row-start-1 mb-1 md:col-start-1 xl:col-span-2 text-coolGray-600">
                Oct 13, 2020
              </time>
              <p className="ml-8 md:col-start-2 md:col-span-4 xl:col-start-3 xl:col-span-9 md:ml-0 text-coolGray-700">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Similique saepe exercitationem numquam, labore necessitatibus
                deleniti quasi. Illo porro nihil necessitatibus debitis delectus aperiam, fuga impedit assumenda odit, velit eveniet est.
              </p>
            </a>
          </article>
        </li>
        <li>
          <article>
            <a
              href="#"
              className="grid p-4 overflow-hidden md:grid-cols-5 rounded-xl lg:p-6 xl:grid-cols-12 hover:bg-white dark:bg-gray-800"
            >
              <h3 className="mb-1 ml-8 font-semibold md:col-start-2 md:col-span-4 md:ml-0 xl:col-start-3 xl:col-span-9">
                Earum at ipsa aliquid quis, exercitationem est.
              </h3>
              <time dateTime="" className="row-start-1 mb-1 md:col-start-1 xl:col-span-2 text-coolGray-600">
                Oct 13, 2020
              </time>
              <p className="ml-8 md:col-start-2 md:col-span-4 xl:col-start-3 xl:col-span-9 md:ml-0 text-coolGray-700">
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Similique saepe exercitationem numquam, labore necessitatibus
                deleniti quasi. Illo porro nihil necessitatibus debitis delectus aperiam, fuga impedit assumenda odit, velit eveniet est.
              </p>
            </a>
          </article>
        </li>
      </ul>
    </>
  )
}

export default withAuthSync(Page)
