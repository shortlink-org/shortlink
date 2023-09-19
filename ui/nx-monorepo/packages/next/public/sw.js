if (!self.define) {
  let e,
    s = {}
  const a = (a, n) => (
    (a = new URL(a + '.js', n).href),
    s[a] ||
      new Promise((s) => {
        if ('document' in self) {
          const e = document.createElement('script')
          ;(e.src = a), (e.onload = s), document.head.appendChild(e)
        } else (e = a), importScripts(a), s()
      }).then(() => {
        let e = s[a]
        if (!e) throw new Error(`Module ${a} didn’t register its module`)
        return e
      })
  )
  self.define = (n, t) => {
    const c =
      e ||
      ('document' in self ? document.currentScript.src : '') ||
      location.href
    if (s[c]) return
    let i = {}
    const r = (e) => a(e, c),
      f = { module: { uri: c }, exports: i, require: r }
    s[c] = Promise.all(n.map((e) => f[e] || r(e))).then((e) => (t(...e), i))
  }
}
define(['./workbox-8e51679a'], function (e) {
  'use strict'
  importScripts(),
    self.skipWaiting(),
    e.clientsClaim(),
    e.precacheAndRoute(
      [
        {
          url: '/next/_next/static/chunks/103-7a3d92f2c397f1e8.js',
          revision: '7a3d92f2c397f1e8',
        },
        {
          url: '/next/_next/static/chunks/103-7a3d92f2c397f1e8.js.map',
          revision: '97f23bf28496f63f6ad52bebd33d20c1',
        },
        {
          url: '/next/_next/static/chunks/159-6babbe0bff343fc1.js',
          revision: '6babbe0bff343fc1',
        },
        {
          url: '/next/_next/static/chunks/159-6babbe0bff343fc1.js.map',
          revision: '617fed7beb7a0d793da6a19e2b396ea8',
        },
        {
          url: '/next/_next/static/chunks/242-a0c86a69953eade7.js',
          revision: 'a0c86a69953eade7',
        },
        {
          url: '/next/_next/static/chunks/242-a0c86a69953eade7.js.map',
          revision: '9d146ae97f9fe4d9f5c23566c0442d1d',
        },
        {
          url: '/next/_next/static/chunks/252-9d9e965332346746.js',
          revision: '9d9e965332346746',
        },
        {
          url: '/next/_next/static/chunks/252-9d9e965332346746.js.map',
          revision: '695ebf5584b06ebf9cae862e7ad1a22d',
        },
        {
          url: '/next/_next/static/chunks/2edb282b-8ea3523ca57960c6.js',
          revision: '8ea3523ca57960c6',
        },
        {
          url: '/next/_next/static/chunks/2edb282b-8ea3523ca57960c6.js.map',
          revision: '0411ef9b3930e75100a4fea98f75963b',
        },
        {
          url: '/next/_next/static/chunks/330-e04bd837b9bc3ef9.js',
          revision: 'e04bd837b9bc3ef9',
        },
        {
          url: '/next/_next/static/chunks/330-e04bd837b9bc3ef9.js.map',
          revision: '406f4fed68662b930640f7aac1861770',
        },
        {
          url: '/next/_next/static/chunks/432-deba2960cb4eb44f.js',
          revision: 'deba2960cb4eb44f',
        },
        {
          url: '/next/_next/static/chunks/432-deba2960cb4eb44f.js.map',
          revision: '1f743e2a79cbf35e950975e02d7a171e',
        },
        {
          url: '/next/_next/static/chunks/579-351c62aa2ca91c36.js',
          revision: '351c62aa2ca91c36',
        },
        {
          url: '/next/_next/static/chunks/579-351c62aa2ca91c36.js.map',
          revision: 'ba3487bc523298ade8f5f2804a324e01',
        },
        {
          url: '/next/_next/static/chunks/690-7c399d580f31b9ec.js',
          revision: '7c399d580f31b9ec',
        },
        {
          url: '/next/_next/static/chunks/75-5657538f1d9046a9.js',
          revision: '5657538f1d9046a9',
        },
        {
          url: '/next/_next/static/chunks/75-5657538f1d9046a9.js.map',
          revision: 'ed92c595a0023b56d540abc0e5955a51',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-00fc895afcce0d01.js',
          revision: '00fc895afcce0d01',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-00fc895afcce0d01.js.map',
          revision: '4d49c1b789a4f31f3cf0cf543b37a381',
        },
        {
          url: '/next/_next/static/chunks/e1533f8b-8291906d765f7c61.js',
          revision: '8291906d765f7c61',
        },
        {
          url: '/next/_next/static/chunks/e1533f8b-8291906d765f7c61.js.map',
          revision: '9fb63f8aaa7310ce21b10fb36d0219e7',
        },
        {
          url: '/next/_next/static/chunks/framework-10711a76a3aa9ab5.js',
          revision: '10711a76a3aa9ab5',
        },
        {
          url: '/next/_next/static/chunks/framework-10711a76a3aa9ab5.js.map',
          revision: '027c1525b0f9ef4554ff9ddffdb01bac',
        },
        {
          url: '/next/_next/static/chunks/main-41a81036960ec0fd.js',
          revision: '41a81036960ec0fd',
        },
        {
          url: '/next/_next/static/chunks/main-41a81036960ec0fd.js.map',
          revision: 'e79a5e4f4e852fe90c275929cbc52500',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-813168fc1fe7134a.js',
          revision: '813168fc1fe7134a',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-813168fc1fe7134a.js.map',
          revision: '93dc52b198ad39446f8561fba7a27ca0',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-b37f45c890707023.js',
          revision: 'b37f45c890707023',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-b37f45c890707023.js.map',
          revision: '468ca005dd1b3f51394f17eb14efdc02',
        },
        {
          url: '/next/_next/static/chunks/pages/about-04a5f637540a6c2f.js',
          revision: '04a5f637540a6c2f',
        },
        {
          url: '/next/_next/static/chunks/pages/about-04a5f637540a6c2f.js.map',
          revision: 'c73b7750b33662ba018931ae6405a16b',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-df31305715d7de0e.js',
          revision: 'df31305715d7de0e',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-df31305715d7de0e.js.map',
          revision: '059176005b43bb45a0e7b6f3d9158527',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-c0b6a3878bde5832.js',
          revision: 'c0b6a3878bde5832',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-c0b6a3878bde5832.js.map',
          revision: '48812bc7cb87d2430863d54c4ca87724',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-d2cfdd094383ec5b.js',
          revision: 'd2cfdd094383ec5b',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-d2cfdd094383ec5b.js.map',
          revision: '3cf7466bf06c3074f58f6c742a0a1dc5',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-fe0b7b056333e921.js',
          revision: 'fe0b7b056333e921',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-fe0b7b056333e921.js.map',
          revision: '354701ecde655e81416227faca10a5b2',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-6ed64509aeb72511.js',
          revision: '6ed64509aeb72511',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-6ed64509aeb72511.js.map',
          revision: '55c5471f6cafd23e5bdf21b08901168c',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-8da03ceaea8e7cae.js',
          revision: '8da03ceaea8e7cae',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-8da03ceaea8e7cae.js.map',
          revision: 'ab8a7e110896c362061b515944d07f95',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-6a1a05b2bd1ef644.js',
          revision: '6a1a05b2bd1ef644',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-6a1a05b2bd1ef644.js.map',
          revision: 'bea3207db65097c5679159627f5443a7',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-1167eefb6fa3c24c.js',
          revision: '1167eefb6fa3c24c',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-1167eefb6fa3c24c.js.map',
          revision: 'bdb09a932fcec2780b7db990c200f977',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-6e2618ecc6210d5e.js',
          revision: '6e2618ecc6210d5e',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-6e2618ecc6210d5e.js.map',
          revision: '74896a71a94a26eee0393387a4b1bd31',
        },
        {
          url: '/next/_next/static/chunks/pages/index-2c4febb416db7a83.js',
          revision: '2c4febb416db7a83',
        },
        {
          url: '/next/_next/static/chunks/pages/index-2c4febb416db7a83.js.map',
          revision: '2f4cad7f0f9be68aa685cd555fc55ee2',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-2f49f9a58bbae4f9.js',
          revision: '2f49f9a58bbae4f9',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-2f49f9a58bbae4f9.js.map',
          revision: 'd6ba11f4735fc1f05917df8a7f5cdd45',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-4c09500405bc3547.js',
          revision: '4c09500405bc3547',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-4c09500405bc3547.js.map',
          revision: '6541888bb52ea5b16ad373f02ea7e59b',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-d3b212c720f0582c.js',
          revision: 'd3b212c720f0582c',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-d3b212c720f0582c.js.map',
          revision: 'ae04715beca2519aa9183a87772f599b',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-d529c41c495f5445.js',
          revision: 'd529c41c495f5445',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-d529c41c495f5445.js.map',
          revision: 'ca7862d7cfc78f774d811cf6de009fed',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-422189d3ad144c4e.js',
          revision: '422189d3ad144c4e',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-422189d3ad144c4e.js.map',
          revision: '9b111cb3330fd341ba3404c9f51bc623',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-5636a2a354007c51.js',
          revision: '5636a2a354007c51',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-5636a2a354007c51.js.map',
          revision: '4664be7b4cc403a2761094765047556e',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-931f380ed7b10a62.js',
          revision: '931f380ed7b10a62',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-931f380ed7b10a62.js.map',
          revision: 'fddd2c2040b3a0da63a2bf41b1d1fa71',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-328003d5ca4f03dc.js',
          revision: '328003d5ca4f03dc',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-328003d5ca4f03dc.js.map',
          revision: '13ef2ca40e633549fcd03772bea27b49',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-7b63b919f17d82f9.js',
          revision: '7b63b919f17d82f9',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-7b63b919f17d82f9.js.map',
          revision: '5a4a37eb139a351087908710d4805efb',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-db753e69ceed74ff.js',
          revision: 'db753e69ceed74ff',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-db753e69ceed74ff.js.map',
          revision: '3135dc1f65df3d3f9d8f3a49e7ed0f18',
        },
        {
          url: '/next/_next/static/chunks/polyfills-78c92fac7aa8fdd8.js',
          revision: '79330112775102f91e1010318bae2bd3',
        },
        {
          url: '/next/_next/static/chunks/webpack-8c891280b3e5141a.js',
          revision: '8c891280b3e5141a',
        },
        {
          url: '/next/_next/static/chunks/webpack-8c891280b3e5141a.js.map',
          revision: '9acab3b76bfcdb9af09408b701e03c4e',
        },
        {
          url: '/next/_next/static/css/07df0638c2c5e066.css',
          revision: '07df0638c2c5e066',
        },
        {
          url: '/next/_next/static/css/07df0638c2c5e066.css.map',
          revision: '2e9dc6485c3e85f382a37307ddcd2ad3',
        },
        {
          url: '/next/_next/static/rFdvhK3FB7jPihBMt6qQb/_buildManifest.js',
          revision: '74372978a62fddd1e96f94a04916054e',
        },
        {
          url: '/next/_next/static/rFdvhK3FB7jPihBMt6qQb/_ssgManifest.js',
          revision: 'b6652df95db52feb4daf4eca35380933',
        },
      ],
      { ignoreURLParametersMatching: [] },
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
        !(!e || s.startsWith('/api/auth/') || !s.startsWith('/api/')),
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
      ({ request: e, url: { pathname: s }, sameOrigin: a }) =>
        '1' === e.headers.get('RSC') &&
        '1' === e.headers.get('Next-Router-Prefetch') &&
        a &&
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
      ({ request: e, url: { pathname: s }, sameOrigin: a }) =>
        '1' === e.headers.get('RSC') && a && !s.startsWith('/api/'),
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
