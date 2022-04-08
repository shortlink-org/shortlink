import { Layout } from 'components'

export const Faq = () => (
  <Layout>
    <div className="px-4 py-16 mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-24 lg:px-8 lg:py-20 bg-white rounded">
      <div className="max-w-xl sm:mx-auto lg:max-w-2xl">
        <div className="max-w-xl mb-10 md:mx-auto sm:text-center lg:max-w-2xl md:mb-12">
          <div>
            <p className="inline-block px-3 py-px mb-4 text-xs font-semibold tracking-wider text-teal-900 uppercase rounded-full bg-teal-accent-400">
              Brand new
            </p>
          </div>
          <h2 className="max-w-lg mb-6 font-sans text-3xl font-bold leading-none tracking-tight text-gray-900 sm:text-4xl md:mx-auto">
            <span className="relative inline-block">
              <svg
                viewBox="0 0 52 24"
                fill="currentColor"
                className="absolute top-0 left-0 z-0 hidden w-32 -mt-8 -ml-20 text-blue-gray-100 lg:w-32 lg:-ml-28 lg:-mt-10 sm:block"
              >
                <defs>
                  <pattern
                    id="70326c9b-4a0f-429b-9c76-792941e326d5"
                    x="0"
                    y="0"
                    width=".135"
                    height=".30"
                  >
                    <circle cx="1" cy="1" r=".7" />
                  </pattern>
                </defs>
                <rect
                  fill="url(#70326c9b-4a0f-429b-9c76-792941e326d5)"
                  width="52"
                  height="24"
                />
              </svg>
              <span className="relative">The</span>
            </span>{' '}
            quick, brown fox jumps over a lazy dog
          </h2>
          <p className="text-base text-gray-700 md:text-lg">
            Sed ut perspiciatis unde omnis iste natus error sit voluptatem
            accusantium doloremque rem aperiam, eaque ipsa quae.
          </p>
        </div>
      </div>
      <div className="max-w-screen-xl sm:mx-auto">
        <div className="grid grid-cols-1 gap-16 row-gap-8 lg:grid-cols-2">
          <div className="space-y-8">
            <div>
              <p className="mb-4 text-xl font-medium">
                The quick, brown fox jumps over a lazy dog?
              </p>
              <p className="text-gray-700">
                Space, the final frontier. These are the voyages of the Starship
                Enterprise. Its five-year mission: to explore strange new
                worlds.
                <br />
                <br />
                Many say exploration is part of our destiny, but it’s actually
                our duty to future generations.
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium">
                The first mate and his Skipper too will do?
              </p>
              <p className="text-gray-700">
                Well, the way they make shows is, they make one show. That
                show's called a pilot.
                <br />
                <br />
                Then they show that show to the people who make shows, and on
                the strength of that one show they decide if they're going to
                make more shows. Some pilots get picked and become television
                programs.Some don't, become nothing. She starred in one of the
                ones that became nothing.
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium">
                Is the Space Pope reptilian!?
              </p>
              <p className="text-gray-700">
                A flower in my garden, a mystery in my panties. Heart attack
                never stopped old Big Bear. I didn't even know we were calling
                him Big Bear.
              </p>
            </div>
          </div>
          <div className="space-y-8">
            <div>
              <p className="mb-4 text-xl font-medium">
                How much money you got on you?
              </p>
              <p className="text-gray-700">
                The first mate and his Skipper too will do their very best to
                make the others comfortable in their tropic island nest.
                <br />
                <br />
                Michael Knight a young loner on a crusade to champion the cause
                of the innocent. The helpless. The powerless in a world of
                criminals who operate above the law. Here he comes Here comes
                Speed Racer. He's a demon on wheels.
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium">
                Galaxies Orion's sword globular star cluster?
              </p>
              <p className="text-gray-700">
                A business big enough that it could be listed on the NASDAQ goes
                belly up. Disappears!
                <br />
                <br />
                It ceases to exist without me. No, you clearly don't know who
                you're talking to, so let me clue you in.
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium">
                When has justice ever been as simple as a rule book?
              </p>
              <p className="text-gray-700">
                This is not about revenge. This is about justice. A lot of
                things can change in twelve years, Admiral. Well, that's
                certainly good to know. About four years. I got tired of hearing
                how young I looked.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </Layout>
)

export default Faq
