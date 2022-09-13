// @ts-nocheck
import React from 'react'
import { Layout } from 'components'
import Ready from 'components/Landing/Ready'
import withAuthSync from 'components/Private'
import { NextSeo } from "next-seo";

export function Reports() {
  return (
    <Layout>
      <NextSeo
        title="Reports"
        description="Reports page for your account."
      />

      <div className="px-4 py-4 my-3 rounded mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-14 lg:px-8 lg:py-10 bg-white dark:bg-gray-800">
        <div className="flex flex-col">
          <p className="text-gray-800">
            Reporting is a critical part of our shortlink service.
            Depending on your settings, we can generate a comprehensive report on your vitals either daily, weekly, monthly, quarterly or yearly.
            This report can include things like how many people clicked on your links, where they came from, and what kind of device they were using.
            This information can be extremely valuable in understanding your audience and tailoring your content to them.
            Additionally, we can customize the reports to include only the information that you are interested in.
            Our reporting system is designed to be flexible and user-friendly, so you can get the most out of it.
          </p>

          <br />

          <p className="text-gray-800">
            Whether you're looking to track your progress over time or want to stay informed about the latest changes in your industry,
            our reporting feature will give you the information you need to make decisions that improve your bottom line.
          </p>
        </div>
      </div>

      <Ready />
    </Layout>
  )
}

export default withAuthSync(Reports)
