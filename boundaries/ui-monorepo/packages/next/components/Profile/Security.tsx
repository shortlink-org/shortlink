import * as React from 'react'
import _ from 'lodash'

// @ts-ignore
const Secutiry = ({ session }) => (
  <div className="md:grid md:grid-cols-3 md:gap-6">
    <div className="md:col-span-1">
      <div className="px-4 sm:px-0">
        <h3 className="text-lg font-medium leading-6 text-gray-900 dark:text-gray-200">Password</h3>
        <p className="mt-1 text-sm text-gray-600 dark:text-gray-400">Change your password.</p>
      </div>
    </div>

    <div className="mt-5 md:mt-0 md:col-span-2">
      <form action="#" method="POST">
        <div className="shadow overflow-hidden sm:rounded-md">
          <div className="px-4 py-5 bg-white dark:bg-gray-800 sm:p-6">
            <div className="grid grid-cols-6 gap-6">
              <div className="col-span-6 sm:col-span-3">
                <label htmlFor="old_password" className="block text-sm font-medium text-gray-700">
                  Old Password
                </label>
                <input
                  type="text"
                  name="old_password"
                  id="old_password"
                  autoComplete="family-name"
                  value={_.get(session, 'kratos.identity.traits.name.last')}
                  className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                />
              </div>

              <div className="col-span-6 sm:col-span-3">
                <label htmlFor="new_password" className="block text-sm font-medium text-gray-700">
                  New Password
                </label>
                <input
                  type="text"
                  name="new_password"
                  id="new_password"
                  autoComplete="family-name"
                  value={_.get(session, 'kratos.identity.traits.name.last')}
                  className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                />
              </div>

              <div className="col-span-6 sm:col-span-3">
                <label htmlFor="confirm_new_password" className="block text-sm font-medium text-gray-700">
                  Confirm New Password
                </label>
                <input
                  type="text"
                  name="confirm_new_password"
                  id="confirm_new_password"
                  autoComplete="family-name"
                  value={_.get(session, 'kratos.identity.traits.name.last')}
                  className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"
                />
              </div>
            </div>
          </div>

          <div className="px-4 py-3 bg-gray-50 text-right sm:px-6">
            <button
              type="submit"
              className="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Save
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
)

export default Secutiry
