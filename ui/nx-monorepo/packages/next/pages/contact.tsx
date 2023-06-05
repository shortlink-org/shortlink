import { Layout } from 'components'
import { ArticleJsonLd, NextSeo, CorporateContactJsonLd } from 'next-seo'

export const ContactContent = () => (
  <Layout>
    <NextSeo
      title="A Better Way to Connect"
      description="We provide an easy way for you to contact us. You can either fill out the form on this page or use one of the other methods listed below. We will get back to you as soon as possible!"
      openGraph={{
        title: 'A Better Way to Connect',
        description:
          'We provide an easy way for you to contact us. You can either fill out the form on this page or use one of the other methods listed below. We will get back to you as soon as possible!',
        type: 'article',
        article: {
          publishedTime: '2021-08-01T05:00:00.000Z',
          modifiedTime: '2021-08-01T05:00:00.000Z',
          section: 'Contact',
          authors: ['https://batazor.ru'],
          tags: ['shortlink', 'contact'],
        },
      }}
    />
    <ArticleJsonLd
      url="https://shortlink.best/next/about"
      title="Contact"
      images={['https://shortlink.best/images/logo.png']}
      datePublished="2021-08-01T05:00:00.000Z"
      dateModified="2021-08-01T05:00:00.000Z"
      authorName={[
        {
          name: 'Login Viktor',
          url: 'https://batazor.ru',
        },
      ]}
      publisherName="Login Viktor"
      publisherLogo="https://shortlink.best/images/logo.png"
      description="We provide an easy way for you to contact us. You can either fill out the form on this page or use one of the other methods listed below. We will get back to you as soon as possible!"
    />
    <CorporateContactJsonLd
      url="https://shortlink.best/next/contact"
      logo="https://shortlink.best/images/logo.png"
      contactPoint={[
        {
          telephone: '+7 (999) 999-99-99',
          contactType: 'customer service',
          contactOption: 'TollFree',
          areaServed: 'RU',
          availableLanguage: ['Russian', 'English'],
        },
      ]}
    />

    <section className="bg-white dark:bg-gray-900 rounded-md">
      <div className="py-8 lg:py-16 px-4 mx-auto max-w-screen-md">
        <h1 className="mb-4 text-4xl tracking-tight font-extrabold text-center text-gray-900 dark:text-white">
          Contact Us
        </h1>
        <p className="mb-8 lg:mb-16 font-light text-center text-gray-500 dark:text-gray-400 sm:text-xl">
          Got a technical issue? Want to send feedback about a beta feature?
          Need details about our Business plan? Let us know.
        </p>
        <form action="#" className="space-y-8">
          <div>
            <label
              htmlFor="email"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
            >
              Your email
            </label>
            <input
              type="email"
              id="email"
              className="shadow-sm bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500 dark:shadow-sm-light"
              placeholder="name@flowbite.com"
              required
            />
          </div>
          <div>
            <label
              htmlFor="subject"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
            >
              Subject
            </label>
            <input
              type="text"
              id="subject"
              className="block p-3 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 shadow-sm focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500 dark:shadow-sm-light"
              placeholder="Let us know how we can help you"
              required
            />
          </div>
          <div className="sm:col-span-2">
            <label
              htmlFor="message"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
            >
              Your message
            </label>
            <textarea
              id="message"
              rows={6}
              className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg shadow-sm border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
              placeholder="Leave a comment..."
              defaultValue=""
            />
          </div>

          <button
            type="submit"
            className="py-3 px-5 text-sm font-medium text-center text-white rounded-lg bg-primary-700 sm:w-fit hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
          >
            Send message
          </button>
        </form>
      </div>
    </section>
  </Layout>
)

export default ContactContent
