import React from 'react'

function Faq() {
  // <NextSeo
  //   title="FAQ"
  //   description="Do you have questions about shortlinks? This page provides answers to the most common questions we receive. If you need more help, please contact us!"
  //   openGraph={{
  //     title: 'FAQ',
  //     description:
  //       'Do you have questions about shortlinks? This page provides answers to the most common questions we receive. If you need more help, please contact us!',
  //     type: 'article',
  //     article: {
  //       publishedTime: '2021-08-01T05:00:00.000Z',
  //       modifiedTime: '2021-08-01T05:00:00.000Z',
  //       section: 'FAQ',
  //       authors: ['https://batazor.ru'],
  //       tags: ['shortlink', 'faq'],
  //     },
  //   }}
  // />
  // <ArticleJsonLd
  //   url="https://shortlink.best/next/about"
  //   title="FAQ"
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
  //   description="Do you have questions about shortlinks? This page provides answers to the most common questions we receive. If you need more help, please contact us!"
  // />
  // <FAQPageJsonLd
  //   mainEntity={[
  //     {
  //       questionName: 'What is a shortlink?',
  //       acceptedAnswerText:
  //         'A shortlink is a short URL that redirects to a longer URL.',
  //     },
  //     {
  //       questionName: 'How do I create a shortlink?',
  //       acceptedAnswerText:
  //         "You can create a shortlink by clicking on the 'Create' button on the main page.",
  //     },
  //     {
  //       questionName: 'Can I use a shortlink for free?',
  //       acceptedAnswerText: 'Yes, you can use shortlinks for free.',
  //     },
  //   ]}
  // />

  return (
    <div className="px-4 py-16 mx-auto bg-white dark:bg-gray-800 rounded max-w-screen-xl md:px-24 lg:px-8 lg:py-20">
      <div className="max-w-xl mx-auto lg:max-w-2xl">
        <div className="mx-auto mb-10 text-center max-w-xl lg:max-w-2xl md:mb-12">
          <div>
            <p className="inline-block px-3 py-1 mb-4 text-xs font-semibold tracking-wider uppercase rounded-full text-blue-500">
              Thanks for asking!
            </p>
          </div>
          <h2 className="max-w-lg mb-6 font-sans text-3xl font-bold leading-none tracking-tight text-blue-500 sm:text-4xl md:mx-auto">
            The frequently asked questions
          </h2>
          <p className="text-base text-gray-700 md:text-lg">
            Sed ut perspiciatis unde omnis iste natus error sit voluptatem
            accusantium doloremque rem aperiam, eaque ipsa quae.
          </p>
        </div>
      </div>
      <div className="max-w-screen-xl mx-auto">
        <div className="grid gap-16 row-gap-8 lg:grid-cols-2">
          <div className="space-y-8">
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                What kind of product is this?
              </p>
              <p className="text-gray-700">
                This service is for creating shortlinks. It's a great way to
                make sure that your links are easy to remember and share, and
                it's perfect for when you want to post a link on social media or
                in an email. Plus, it's free to use!
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                Is it free?
              </p>
              <p className="text-gray-700">
                Yes, it's free! This is something that's always been free and
                always will be. There are no hidden fees or catches - you can
                use it as much as you want, for as long as you want. We'll never
                ask for your payment information or try to upsell you on any
                features.
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                Is it safe to use this product?
              </p>
              <p className="text-gray-700">
                Yes, it is safe to use this product. We respect GDPR and the
                security of our users' data very seriously. We have implemented
                multiple layers of security to protect users' data, and we will
                never share user data with any third party without explicit
                consent from the user.
              </p>
            </div>
          </div>

          <div className="space-y-8">
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                How is this product better than others?
              </p>
              <p className="text-gray-700">
                This product is absolutely free, and it's made for easy use by
                everyone. There are no catches or gimmicks. We want everyone to
                be able to use this product and get the most out of it.
                <br />
                <br />
                The other products out there might charge you a subscription
                fee, or they might have a lot of complicated features that you
                don't need. With this product, you get exactly what you need
                without any extras.
                <br />
                <br />
                We also have outstanding customer support. If you ever have any
                questions or problems, our team is always happy to help. We're
                here for you 24/7!
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                Is there anything I can do to help this product?
              </p>
              <p className="text-gray-700">
                Yes! You can help us out by forking our project on GitHub and
                making commits. We would really appreciate the help!
                <br />
                <br />
                Thanks for your interest in our project!
              </p>
            </div>
            <div>
              <p className="mb-4 text-xl font-medium text-blue-500">
                Can I deploy this product on my machine?
              </p>
              <p className="text-gray-700">
                Yes, you can deploy this product on your machine. For
                instructions on how to do so, please consult the readme or read
                the docs.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Faq
