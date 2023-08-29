if (!self.define) {
  let e,
    s = {}
  const a = (a, c) => (
    (a = new URL(a + '.js', c).href),
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
  self.define = (c, n) => {
    const t =
      e ||
      ('document' in self ? document.currentScript.src : '') ||
      location.href
    if (s[t]) return
    let i = {}
    const r = (e) => a(e, t),
      u = { module: { uri: t }, exports: i, require: r }
    s[t] = Promise.all(c.map((e) => u[e] || r(e))).then((e) => (n(...e), i))
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
          url: '/next/_next/static/9Uc_dcjJGWu7fSviroh_L/_buildManifest.js',
          revision: 'f889d524d279e2799642e0bc226f7a72',
        },
        {
          url: '/next/_next/static/9Uc_dcjJGWu7fSviroh_L/_ssgManifest.js',
          revision: 'b6652df95db52feb4daf4eca35380933',
        },
        {
          url: '/next/_next/static/chunks/159-ae91797b357abe87.js',
          revision: 'ae91797b357abe87',
        },
        {
          url: '/next/_next/static/chunks/159-ae91797b357abe87.js.map',
          revision: 'cac0b22697b1a683ef362a87021e6e55',
        },
        {
          url: '/next/_next/static/chunks/242-506849d60e73ec29.js',
          revision: '506849d60e73ec29',
        },
        {
          url: '/next/_next/static/chunks/242-506849d60e73ec29.js.map',
          revision: '7ddba3df319b287f3f4146c2339a328a',
        },
        {
          url: '/next/_next/static/chunks/252-fa32efc866a5a488.js',
          revision: 'fa32efc866a5a488',
        },
        {
          url: '/next/_next/static/chunks/252-fa32efc866a5a488.js.map',
          revision: '5889de31bb938c8b8616a3139be1771d',
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
          url: '/next/_next/static/chunks/432-c1905c35b54b58b1.js',
          revision: 'c1905c35b54b58b1',
        },
        {
          url: '/next/_next/static/chunks/432-c1905c35b54b58b1.js.map',
          revision: 'a082d918256e840ced3c0f4ab45b4f1e',
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
          url: '/next/_next/static/chunks/690-e8f572c543ddee74.js',
          revision: 'e8f572c543ddee74',
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
          url: '/next/_next/static/chunks/838-19290cbb557dc997.js',
          revision: '19290cbb557dc997',
        },
        {
          url: '/next/_next/static/chunks/838-19290cbb557dc997.js.map',
          revision: 'a3b7f7bccef1bc2c9742e4338a71b791',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-65c6af72df013a7c.js',
          revision: '65c6af72df013a7c',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-65c6af72df013a7c.js.map',
          revision: '20ee26d68ed125ca5e45aba03af5b878',
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
          url: '/next/_next/static/chunks/pages/_app-36bfa1fa65225599.js',
          revision: '36bfa1fa65225599',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-36bfa1fa65225599.js.map',
          revision: 'f19d47609a88ce17baa01ed175db5dc9',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-59569270d0b4ac69.js',
          revision: '59569270d0b4ac69',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-59569270d0b4ac69.js.map',
          revision: '275d4e56edc51a0b861590672a5a99a6',
        },
        {
          url: '/next/_next/static/chunks/pages/about-32ff57453b3b84fd.js',
          revision: '32ff57453b3b84fd',
        },
        {
          url: '/next/_next/static/chunks/pages/about-32ff57453b3b84fd.js.map',
          revision: 'e8a6b311693b5c35a0a1fc980747b254',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-c8646081b42045f2.js',
          revision: 'c8646081b42045f2',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-c8646081b42045f2.js.map',
          revision: '7c43d3db834fb94b7688d281fd5135a1',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-1b0c3024c23293aa.js',
          revision: '1b0c3024c23293aa',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-1b0c3024c23293aa.js.map',
          revision: '7a3b4db2c45a1c7826faf2b60bc00fbe',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-a7f2537951c3c975.js',
          revision: 'a7f2537951c3c975',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-a7f2537951c3c975.js.map',
          revision: 'c72930dcd9e772bfbe1817bbd35e7ada',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-d43496c0ccaac4c6.js',
          revision: 'd43496c0ccaac4c6',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-d43496c0ccaac4c6.js.map',
          revision: '1e62015afb83ecee2c922dba7f4a74d0',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-13d18728cef72540.js',
          revision: '13d18728cef72540',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-13d18728cef72540.js.map',
          revision: '183deb63eaee43cc29b17e50f5b5ef74',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-47acf0d8d38b7766.js',
          revision: '47acf0d8d38b7766',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-47acf0d8d38b7766.js.map',
          revision: 'eb5ec5a1da4aade08c69fa6a25b12770',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-fdf842d71cf95d07.js',
          revision: 'fdf842d71cf95d07',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-fdf842d71cf95d07.js.map',
          revision: '5a694e6fef0c3d8c3deb9bcf4fe4b1b6',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-90d9526c35ce156c.js',
          revision: '90d9526c35ce156c',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-90d9526c35ce156c.js.map',
          revision: 'b7712f2405db501cc7641c4c5d84a233',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-edf73bba45144faf.js',
          revision: 'edf73bba45144faf',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-edf73bba45144faf.js.map',
          revision: 'fdaa46c53443c754d539dc804af9f6d9',
        },
        {
          url: '/next/_next/static/chunks/pages/index-5183b9ce751cccfc.js',
          revision: '5183b9ce751cccfc',
        },
        {
          url: '/next/_next/static/chunks/pages/index-5183b9ce751cccfc.js.map',
          revision: '1b222be449529e8130e42dc76209250a',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-043115c5c3d2ac2f.js',
          revision: '043115c5c3d2ac2f',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-043115c5c3d2ac2f.js.map',
          revision: 'f697774f02c306ac2968d0ff2ff4a343',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-a6ee54425e5472da.js',
          revision: 'a6ee54425e5472da',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-a6ee54425e5472da.js.map',
          revision: '2ba3c8219704069f81755f0ca5bf2bfe',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-09e68282dc81324b.js',
          revision: '09e68282dc81324b',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-09e68282dc81324b.js.map',
          revision: '5122d2ef7adfc6aca12440c0014f0510',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-6e228cc71220f610.js',
          revision: '6e228cc71220f610',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-6e228cc71220f610.js.map',
          revision: 'a160df354fcd783be5b44812636faee0',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-b9908374e4b48a34.js',
          revision: 'b9908374e4b48a34',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-b9908374e4b48a34.js.map',
          revision: '82be933056fd12a8cf2a2bad4dbe62e0',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-40fde1790d0b0691.js',
          revision: '40fde1790d0b0691',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-40fde1790d0b0691.js.map',
          revision: '340cf167c42d681309a1b53ee1aea335',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-bf3ce82a1ad81d0d.js',
          revision: 'bf3ce82a1ad81d0d',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-bf3ce82a1ad81d0d.js.map',
          revision: '5be548631d77ac39fe39896d5256b782',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-f2e43a3bbe6924b6.js',
          revision: 'f2e43a3bbe6924b6',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-f2e43a3bbe6924b6.js.map',
          revision: 'db4faf7cb46af6a7849a09ba31d4d446',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-854b8c9f607f3f1f.js',
          revision: '854b8c9f607f3f1f',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-854b8c9f607f3f1f.js.map',
          revision: '22ccb227ecf2c60c81d32b768b2ccb7a',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-e9e3bcd49efcd2e9.js',
          revision: 'e9e3bcd49efcd2e9',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-e9e3bcd49efcd2e9.js.map',
          revision: '1f91280427096303c083355efdc9b144',
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
          url: '/next/_next/static/css/5e8a6beca3a80d57.css',
          revision: '5e8a6beca3a80d57',
        },
        {
          url: '/next/_next/static/css/5e8a6beca3a80d57.css.map',
          revision: '023c2c662d782371de7d24c64c5435f8',
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
      /\.(?:mp4)$/i,
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
