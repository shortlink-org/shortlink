'use client'

import UndrawCoworkers from '../../public/assets/images/undraw_back_in_the_day_knsh.svg'

// <NextSeo
// title="Page"
// description="Page page for shortlink."
// openGraph={{
//   title: 'About',
//     description: 'About page for shortlink.',
//     type: 'article',
//     article: {
//     publishedTime: '2021-08-01T05:00:00.000Z',
//       modifiedTime: '2021-08-01T05:00:00.000Z',
//       section: 'About',
//       authors: ['https://batazor.ru'],
//       tags: ['shortlink', 'about'],
//   },
// }}
// />
// <ArticleJsonLd
//   url="https://shortlink.best/next/about"
//   title="Page"
//   images={['https://shortlink.best/images/logo.png']}
//   datePublished="2021-08-01T05:00:00.000Z"
//   dateModified="2021-08-01T05:00:00.000Z"
//   authorName={[
//     {
//       name: 'Login Viktor',
//       url: 'https://batazor.ru',
//     },
//   ]}
//   publisherName="Login Viktor"
//   publisherLogo="https://shortlink.best/images/logo.png"
//   description="Page page for shortlink."
// />

function Page() {
  return (
    <>
      <div className="sm:flex items-center max-w-screen-xl bg-white dark:bg-gray-800 rounded my-8">
        <div className="sm:w-1/2 p-5">
          <div className="image object-center text-center">
            <UndrawCoworkers />
          </div>
        </div>
        <div className="sm:w-1/2 p-5">
          <div className="text">
            <span className="text-gray-500 border-b-2 border-indigo-600 uppercase">Page us</span>
            <h2 className="my-4 font-bold text-3xl  sm:text-4xl ">
              Page <span className="text-indigo-600">Shortlink</span>
            </h2>

            <article className="prose dark:prose-invert">
              <p className="text-gray-700">
                At shortlink, we're all about making it easy for people to connect and share. Our product is designed to work effectively
                with links, making it simple and convenient for people to share content and collaborate online. We believe that
                communication should be easy and accessible for everyone, and our mission is to make sure that's always the case. With
                shortlink, sharing content and collaborating with others is more convenient than ever before.
              </p>

              <p className="text-gray-700">
                We're on a mission to provide the best possible solution for software developers. We want to make it easy for developers to
                find the best practices and follow them. That's why we created shortlink.
              </p>

              <p className="text-gray-700">
                Shortlink is an open source project that provides a pretty user interface and respects GDPR. We use edge technologies and
                have many years of experience. We're constantly researching the best solutions on the market so that we can benefit our
                community and solve a problem for millions of people.
              </p>

              <p className="text-gray-700">
                We know that there are not enough advanced and flexible solutions out there. That's why we're offering our product for free.
                We want to help as many people as possible and become market leaders in the process.
              </p>

              <p className="text-gray-700">
                If you're looking for the best possible solution, look no further than shortlink. Try it today and see for yourself how easy
                it is to find the best practices and follow them.
              </p>
            </article>
          </div>
        </div>
      </div>
    </>
  )
}

export default Page
