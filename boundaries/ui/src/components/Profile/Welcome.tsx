// @ts-nocheck

export default function Welcome({ nickname }: { nickname: string }) {
  return (
    <div className="md:w-auto flex items-center content-center my-6 flex-auto bg-indigo-600 text-white rounded shadow-xl px-5 w-full">
      <div className="flex flex-wrap content-center items-center">
        <div className="w-1/4 px-3 text-center hidden md:block">
          <div className="p-5 xl:px-8 md:py-5">
            <img 
              src="/assets/images/undraw_welcome_cats_thqn.svg" 
              alt="Welcome cats illustration" 
              className="w-full h-auto" 
            />
          </div>
        </div>
        <div className="w-full sm:w-1/2 md:w-2/4 px-3 text-left">
          <div className="p-5 xl:px-8 md:py-5">
            <h3 className="text-2xl">Welcome, {nickname}!</h3>
            <h5 className="text-xl mb-3">nice to meet you!</h5>
            <p className="text-sm text-indigo-200">
              Welcome to the Service Shortlink! We are excited to offer this new service to our customers. With the Service Shortlink, you
              will be able to easily access your favorite services with a short, easy-to-remember URL. Simply enter the URL into your
              browser and you will be taken directly to the service you requested. We hope you find this new service convenient and easy to
              use.
              <br />
              <br />
              <b>Thank you for choosing the Service Shortlink!</b>
            </p>
          </div>
        </div>
        <div className="w-full sm:w-1/2 md:w-1/4 px-3 text-center">
          <div className="p-5 xl:px-8 md:py-5">
            <a
              className="block w-full py-2 px-4 rounded text-indigo-600 bg-gray-200 hover:bg-white dark:bg-gray-800 hover:text-gray-900 focus:outline-none transition duration-150 ease-in-out mb-3"
              href="https://codepen.io/ScottWindon"
              target="_blank"
              rel="noreferrer"
            >
              Find out more?
            </a>
            <button
              type="button"
              className="w-full py-2 px-4 rounded text-white bg-indigo-900 hover:bg-gray-900 focus:outline-none transition duration-150 ease-in-out"
            >
              No thanks
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
