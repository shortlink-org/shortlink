import React from 'react'
import { PaperClipIcon } from '@heroicons/react/24/solid'
import { Layout } from 'components'
// @ts-ignore
import UndrawCoworkers from '../public/assets/images/undraw_back_in_the_day_knsh.svg'
import { ArticleJsonLd, NextSeo } from "next-seo";

function NewComponent() {
  return (
    <div className="sm:flex items-center max-w-screen-xl bg-white dark:bg-gray-800 rounded my-8">
      <NextSeo
        title="About"
        description="About page for shortlink."
        openGraph={{
          title: "About",
          description: "About page for shortlink.",
          type: "article",
          article: {
            publishedTime: "2021-08-01T05:00:00.000Z",
            modifiedTime: "2021-08-01T05:00:00.000Z",
            section: "About",
            authors: [
              "https://batazor.ru",
            ],
            tags: [ "shortlink", "about" ],
          }
        }}
      />
      <ArticleJsonLd
        url="https://architecture.ddns.net/next/about"
        title="About"
        images={[
          'https://architecture.ddns.net/images/logo.png',
        ]}
        datePublished="2021-08-01T05:00:00.000Z"
        dateModified="2021-08-01T05:00:00.000Z"
        authorName={[
          {
            name: 'Login Viktor',
            url: 'https://batazor.ru',
          },
        ]}
        publisherName="Login Viktor"
        publisherLogo="https://architecture.ddns.net/images/logo.png"
        description="About page for shortlink."
      />
      <div className="sm:w-1/2 p-5">
        <div className="image object-center text-center">
          <UndrawCoworkers />
        </div>
      </div>
      <div className="sm:w-1/2 p-5">
        <div className="text">
          <span className="text-gray-500 border-b-2 border-indigo-600 uppercase">
            About us
          </span>
          <h2 className="my-4 font-bold text-3xl  sm:text-4xl ">
            About <span className="text-indigo-600">Shortlink</span>
          </h2>

          <article className="prose dark:prose-invert">
            <p className="text-gray-700">
              At shortlink, we're all about making it easy for people to connect and share.
              Our product is designed to work effectively with links, making it simple and convenient for people to share content and collaborate online.
              We believe that communication should be easy and accessible for everyone, and our mission is to make sure that's always the case.
              With shortlink, sharing content and collaborating with others is more convenient than ever before.
            </p>

            <p className="text-gray-700">
              We're on a mission to provide the best possible solution for software developers.
              We want to make it easy for developers to find the best practices and follow them.
              That's why we created shortlink.
            </p>

            <p className="text-gray-700">
              Shortlink is an open source project that provides a pretty user interface and respects GDPR.
              We use edge technologies and have many years of experience.
              We're constantly researching the best solutions on the market so that we can benefit our community and
              solve a problem for millions of people.
            </p>

            <p className="text-gray-700">
              We know that there are not enough advanced and flexible solutions out there.
              That's why we're offering our product for free.
              We want to help as many people as possible and become market leaders in the process.
            </p>

            <p className="text-gray-700">
              If you're looking for the best possible solution, look no further than shortlink.
              Try it today and see for yourself how easy it is to find the best practices and follow them.
            </p>
          </article>
        </div>
      </div>
    </div>
  )
}

const About = () => (
  <Layout>
    <NewComponent />

    <div className="bg-white dark:bg-gray-800 shadow overflow-hidden sm:rounded-lg">
      <div className="px-4 py-5 sm:px-6">
        <h3 className="text-lg leading-6 font-medium text-gray-900">
          Applicant Information
        </h3>
        <p className="mt-1 max-w-2xl text-sm text-gray-500">
          Personal details and application.
        </p>
      </div>
      <div className="border-t border-gray-200">
        <dl>
          <div className="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">Full name</dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              Margot Foster
            </dd>
          </div>
          <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">
              Application for
            </dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              Backend Developer
            </dd>
          </div>
          <div className="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">Email address</dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              margotfoster@example.com
            </dd>
          </div>
          <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">
              Salary expectation
            </dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              $120,000
            </dd>
          </div>
          <div className="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">About</dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              Fugiat ipsum ipsum deserunt culpa aute sint do nostrud anim
              incididunt cillum culpa consequat. Excepteur qui ipsum aliquip
              consequat sint. Sit id mollit nulla mollit nostrud in ea officia
              proident. Irure nostrud pariatur mollit ad adipisicing
              reprehenderit deserunt qui eu.
            </dd>
          </div>
          <div className="bg-white dark:bg-gray-800 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
            <dt className="text-sm font-medium text-gray-500">Attachments</dt>
            <dd className="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2">
              <ul className="border border-gray-200 rounded-md divide-y divide-gray-200">
                <li className="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                  <div className="w-0 flex-1 flex items-center">
                    <PaperClipIcon
                      className="flex-shrink-0 h-5 w-5 text-gray-400"
                      aria-hidden="true"
                    />
                    <span className="ml-2 flex-1 w-0 truncate">
                      resume_back_end_developer.pdf
                    </span>
                  </div>
                  <div className="ml-4 flex-shrink-0">
                    <a
                      href="#"
                      className="font-medium text-indigo-600 hover:text-indigo-500"
                    >
                      Download
                    </a>
                  </div>
                </li>
                <li className="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                  <div className="w-0 flex-1 flex items-center">
                    <PaperClipIcon
                      className="flex-shrink-0 h-5 w-5 text-gray-400"
                      aria-hidden="true"
                    />
                    <span className="ml-2 flex-1 w-0 truncate">
                      coverletter_back_end_developer.pdf
                    </span>
                  </div>
                  <div className="ml-4 flex-shrink-0">
                    <a
                      href="#"
                      className="font-medium text-indigo-600 hover:text-indigo-500"
                    >
                      Download
                    </a>
                  </div>
                </li>
              </ul>
            </dd>
          </div>
        </dl>
      </div>
    </div>
  </Layout>
)

export default About
