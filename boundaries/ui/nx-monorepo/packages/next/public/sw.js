if (!self.define) {
  let e,
    s = {}
  const n = (n, a) => (
    (n = new URL(n + '.js', a).href),
    s[n] ||
      new Promise((s) => {
        if ('document' in self) {
          const e = document.createElement('script')
          ;(e.src = n), (e.onload = s), document.head.appendChild(e)
        } else (e = n), importScripts(n), s()
      }).then(() => {
        let e = s[n]
        if (!e) throw new Error(`Module ${n} didn’t register its module`)
        return e
      })
  )
  self.define = (a, t) => {
    const i =
      e ||
      ('document' in self ? document.currentScript.src : '') ||
      location.href
    if (s[i]) return
    let c = {}
    const r = (e) => n(e, i),
      d = { module: { uri: i }, exports: c, require: r }
    s[i] = Promise.all(a.map((e) => d[e] || r(e))).then((e) => (t(...e), c))
  }
}
define(['./workbox-93afbf15'], function (e) {
  'use strict'
  importScripts(),
    self.skipWaiting(),
    e.clientsClaim(),
    e.precacheAndRoute(
      [
        {
          url: '/next/_next/static/AeiEoDitUioPk1k33rx2X/_buildManifest.js',
          revision: 'c3529d73917748edeea15c10cab45d7e',
        },
        {
          url: '/next/_next/static/AeiEoDitUioPk1k33rx2X/_ssgManifest.js',
          revision: 'b6652df95db52feb4daf4eca35380933',
        },
        {
          url: '/next/_next/static/chunks/203-db9408f91da32d18.js',
          revision: 'db9408f91da32d18',
        },
        {
          url: '/next/_next/static/chunks/2edb282b-12c2213d9d30139d.js',
          revision: '12c2213d9d30139d',
        },
        {
          url: '/next/_next/static/chunks/432-9a6e987b65c1d220.js',
          revision: '9a6e987b65c1d220',
        },
        {
          url: '/next/_next/static/chunks/517-adc56f9186d029ed.js',
          revision: 'adc56f9186d029ed',
        },
        {
          url: '/next/_next/static/chunks/579-aa6e598147f2915f.js',
          revision: 'aa6e598147f2915f',
        },
        {
          url: '/next/_next/static/chunks/666-4e8911b6ce5a3797.js',
          revision: '4e8911b6ce5a3797',
        },
        {
          url: '/next/_next/static/chunks/75-4da1d1d4402a8509.js',
          revision: '4da1d1d4402a8509',
        },
        {
          url: '/next/_next/static/chunks/838-19e1c85934a605ac.js',
          revision: '19e1c85934a605ac',
        },
        {
          url: '/next/_next/static/chunks/892-deff1e8c5c179007.js',
          revision: 'deff1e8c5c179007',
        },
        {
          url: '/next/_next/static/chunks/ff417084-2d3c8f724bb03645.js',
          revision: '2d3c8f724bb03645',
        },
        {
          url: '/next/_next/static/chunks/framework-eabb24dd63114e61.js',
          revision: 'eabb24dd63114e61',
        },
        {
          url: '/next/_next/static/chunks/main-6158f457dcc7dbc3.js',
          revision: '6158f457dcc7dbc3',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-4ac2823e2df7fca8.js',
          revision: '4ac2823e2df7fca8',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-3f60081508868eb6.js',
          revision: '3f60081508868eb6',
        },
        {
          url: '/next/_next/static/chunks/pages/about-81a6abff485b2536.js',
          revision: '81a6abff485b2536',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-a6e25a2908529a1e.js',
          revision: 'a6e25a2908529a1e',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-c0c99c6e5deb5eb6.js',
          revision: 'c0c99c6e5deb5eb6',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-b9ad1171b03d1117.js',
          revision: 'b9ad1171b03d1117',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-f436e53c3f16e683.js',
          revision: 'f436e53c3f16e683',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-51f4710ba374e9ef.js',
          revision: '51f4710ba374e9ef',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-1c2a19186991f660.js',
          revision: '1c2a19186991f660',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-a06b9480e988005f.js',
          revision: 'a06b9480e988005f',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-dc805ae791834872.js',
          revision: 'dc805ae791834872',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-8249971256757c99.js',
          revision: '8249971256757c99',
        },
        {
          url: '/next/_next/static/chunks/pages/index-3a3a843353219587.js',
          revision: '3a3a843353219587',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-ba2c132f891c5954.js',
          revision: 'ba2c132f891c5954',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-8189a873b9e708bd.js',
          revision: '8189a873b9e708bd',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-abb717e2a98fd08e.js',
          revision: 'abb717e2a98fd08e',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-249c68b530713bd2.js',
          revision: '249c68b530713bd2',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-09d8ee941393b7ab.js',
          revision: '09d8ee941393b7ab',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-b0d75751f320b9bf.js',
          revision: 'b0d75751f320b9bf',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-884ce8b7bb76ae9f.js',
          revision: '884ce8b7bb76ae9f',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-d7744d0cd34d8627.js',
          revision: 'd7744d0cd34d8627',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-57343857310d56f4.js',
          revision: '57343857310d56f4',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-1ef7dc753c60b103.js',
          revision: '1ef7dc753c60b103',
        },
        {
          url: '/next/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js',
          revision: '837c0df77fd5009c9e46d446188ecfd0',
        },
        {
          url: '/next/_next/static/chunks/webpack-8c891280b3e5141a.js',
          revision: '8c891280b3e5141a',
        },
        {
          url: '/next/_next/static/css/f20e139843013d5d.css',
          revision: 'f20e139843013d5d',
        },
        {
          url: '/next/assets/images/undraw_back_in_the_day_knsh.svg',
          revision: 'aebc6c499a138c3e107e65a208aec647',
        },
        {
          url: '/next/assets/images/undraw_co_workers_re_1i6i.svg',
          revision: 'cb908c2f6d43c3d5bced6e0804dac2e9',
        },
        {
          url: '/next/assets/images/undraw_designer_re_5v95.svg',
          revision: '435c0b4cb909d0ceb63048a4e7ebc9f5',
        },
        {
          url: '/next/assets/images/undraw_welcome_cats_thqn.svg',
          revision: 'ed0c3358facded075949f5e0ab20a080',
        },
        {
          url: '/next/assets/styles.css',
          revision: '4195c317713612b842845b38863389f6',
        },
        {
          url: '/next/favicon.ico',
          revision: 'c30c7d42707a47a3f4591831641e50dc',
        },
        {
          url: '/next/firebase-messaging-sw.js',
          revision: '47db0543b0c9d21608ee0cda826ce944',
        },
        {
          url: '/next/manifest.json',
          revision: '44354e9e77eae0d74431df55175e5566',
        },
        {
          url: '/next/sitemap-0.xml',
          revision: '2ef8f99fa4fba2551f9effced9738793',
        },
      ],
      { ignoreURLParametersMatching: [/^utm_/, /^fbclid$/] },
    ),
    e.cleanupOutdatedCaches(),
    e.registerRoute(
      '/next',
      new e.NetworkFirst({
        cacheName: 'start-url',
        plugins: [
          {
            cacheWillUpdate: async ({ response: e }) =>
              e && 'opaqueredirect' === e.type
                ? new Response(e.body, {
                    status: 200,
                    statusText: 'OK',
                    headers: e.headers,
                  })
                : e,
          },
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,
      new e.CacheFirst({
        cacheName: 'google-fonts-webfonts',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 4, maxAgeSeconds: 31536e3 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,
      new e.StaleWhileRevalidate({
        cacheName: 'google-fonts-stylesheets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 4, maxAgeSeconds: 604800 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'static-font-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 4, maxAgeSeconds: 604800 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'static-image-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 64, maxAgeSeconds: 2592e3 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\/_next\/static.+\.js$/i,
      new e.CacheFirst({
        cacheName: 'next-static-js-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 64, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\/_next\/image\?url=.+$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'next-image',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 64, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:mp3|wav|ogg)$/i,
      new e.CacheFirst({
        cacheName: 'static-audio-assets',
        plugins: [
          new e.RangeRequestsPlugin(),
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:mp4|webm)$/i,
      new e.CacheFirst({
        cacheName: 'static-video-assets',
        plugins: [
          new e.RangeRequestsPlugin(),
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:js)$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'static-js-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 48, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:css|less)$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'static-style-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\/_next\/data\/.+\/.+\.json$/i,
      new e.StaleWhileRevalidate({
        cacheName: 'next-data',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      /\.(?:json|xml|csv)$/i,
      new e.NetworkFirst({
        cacheName: 'static-data-assets',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ sameOrigin: e, url: { pathname: s } }) =>
        !(!e || s.startsWith('/api/auth/callback') || !s.startsWith('/api/')),
      new e.NetworkFirst({
        cacheName: 'apis',
        networkTimeoutSeconds: 10,
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 16, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ request: e, url: { pathname: s }, sameOrigin: n }) =>
        '1' === e.headers.get('RSC') &&
        '1' === e.headers.get('Next-Router-Prefetch') &&
        n &&
        !s.startsWith('/api/'),
      new e.NetworkFirst({
        cacheName: 'pages-rsc-prefetch',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ request: e, url: { pathname: s }, sameOrigin: n }) =>
        '1' === e.headers.get('RSC') && n && !s.startsWith('/api/'),
      new e.NetworkFirst({
        cacheName: 'pages-rsc',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ url: { pathname: e }, sameOrigin: s }) => s && !e.startsWith('/api/'),
      new e.NetworkFirst({
        cacheName: 'pages',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ sameOrigin: e }) => !e,
      new e.NetworkFirst({
        cacheName: 'cross-origin',
        networkTimeoutSeconds: 10,
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 3600 }),
        ],
      }),
      'GET',
    )
})
//# sourceMappingURL=sw.js.map
