/** @type {import('next-sitemap').IConfig} */
const config = {
  siteUrl: process.env.SITE_URL || 'https://shortlink.best',
  generateRobotsTxt: true,
  // optional
  robotsTxtOptions: {
    additionalSitemaps: [],
  },
}

module.exports = config
