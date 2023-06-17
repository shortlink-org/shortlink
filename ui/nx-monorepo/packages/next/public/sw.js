if (!self.define) {
  let e,
    a = {}
  const s = (s, n) => (
    (s = new URL(s + '.js', n).href),
    a[s] ||
      new Promise((a) => {
        if ('document' in self) {
          const e = document.createElement('script')
          ;(e.src = s), (e.onload = a), document.head.appendChild(e)
        } else (e = s), importScripts(s), a()
      }).then(() => {
        let e = a[s]
        if (!e) throw new Error(`Module ${s} didn’t register its module`)
        return e
      })
  )
  self.define = (n, c) => {
    const t =
      e ||
      ('document' in self ? document.currentScript.src : '') ||
      location.href
    if (a[t]) return
    let i = {}
    const r = (e) => s(e, t),
      f = { module: { uri: t }, exports: i, require: r }
    a[t] = Promise.all(n.map((e) => f[e] || r(e))).then((e) => (c(...e), i))
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
          url: '/next/_next/static/1yeXGGu0kOV-JHCNDd3ho/_buildManifest.js',
          revision: '890244801378114bf097c64a0f899892',
        },
        {
          url: '/next/_next/static/1yeXGGu0kOV-JHCNDd3ho/_ssgManifest.js',
          revision: 'b6652df95db52feb4daf4eca35380933',
        },
        {
          url: '/next/_next/static/chunks/159-b4b97110e0b8d56e.js',
          revision: 'b4b97110e0b8d56e',
        },
        {
          url: '/next/_next/static/chunks/159-b4b97110e0b8d56e.js.map',
          revision: 'b7c6adf0579833379099c73e58c70a1f',
        },
        {
          url: '/next/_next/static/chunks/242-cf6360725c9a47c3.js',
          revision: 'cf6360725c9a47c3',
        },
        {
          url: '/next/_next/static/chunks/242-cf6360725c9a47c3.js.map',
          revision: 'ba7d7c0034cfc82e5468603d859f5388',
        },
        {
          url: '/next/_next/static/chunks/252-a162746d4ffd20ad.js',
          revision: 'a162746d4ffd20ad',
        },
        {
          url: '/next/_next/static/chunks/252-a162746d4ffd20ad.js.map',
          revision: '57ba915b277336998cba3850ea856f6a',
        },
        {
          url: '/next/_next/static/chunks/290-ce47ac552fbe66a5.js',
          revision: 'ce47ac552fbe66a5',
        },
        {
          url: '/next/_next/static/chunks/290-ce47ac552fbe66a5.js.map',
          revision: '1dbacedd7a360284df4122edefb58ea7',
        },
        {
          url: '/next/_next/static/chunks/2edb282b-99fa8a592c3e6564.js',
          revision: '99fa8a592c3e6564',
        },
        {
          url: '/next/_next/static/chunks/2edb282b-99fa8a592c3e6564.js.map',
          revision: 'e8bfcc3c2d9e8318e67cca20b7cf968e',
        },
        {
          url: '/next/_next/static/chunks/330-316387a756051cb7.js',
          revision: '316387a756051cb7',
        },
        {
          url: '/next/_next/static/chunks/330-316387a756051cb7.js.map',
          revision: 'a324ffa0e2c371801b7db792b40b9f53',
        },
        {
          url: '/next/_next/static/chunks/432-17c17ba52b93b09e.js',
          revision: '17c17ba52b93b09e',
        },
        {
          url: '/next/_next/static/chunks/432-17c17ba52b93b09e.js.map',
          revision: 'ab9c10afaa6693f00df8ac151909342c',
        },
        {
          url: '/next/_next/static/chunks/579-351c62aa2ca91c36.js',
          revision: '351c62aa2ca91c36',
        },
        {
          url: '/next/_next/static/chunks/579-351c62aa2ca91c36.js.map',
          revision: '55532229f29b656786ba1d50c42ca736',
        },
        {
          url: '/next/_next/static/chunks/690-8ba491c0f6203133.js',
          revision: '8ba491c0f6203133',
        },
        {
          url: '/next/_next/static/chunks/75-3c6a0e70c72c3b85.js',
          revision: '3c6a0e70c72c3b85',
        },
        {
          url: '/next/_next/static/chunks/75-3c6a0e70c72c3b85.js.map',
          revision: 'e0ec9487cb706f73f4064c0ee8ac2722',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-8be1448f11f60f41.js',
          revision: '8be1448f11f60f41',
        },
        {
          url: '/next/_next/static/chunks/97cc2b9f-8be1448f11f60f41.js.map',
          revision: 'a67b16099bb27958fbca12631a3d767a',
        },
        {
          url: '/next/_next/static/chunks/e1533f8b-94b7dda262964aeb.js',
          revision: '94b7dda262964aeb',
        },
        {
          url: '/next/_next/static/chunks/e1533f8b-94b7dda262964aeb.js.map',
          revision: '1e7c31e58ff298072c31cee7fd75c4c4',
        },
        {
          url: '/next/_next/static/chunks/framework-7e48c0ecae6bc03a.js',
          revision: '7e48c0ecae6bc03a',
        },
        {
          url: '/next/_next/static/chunks/framework-7e48c0ecae6bc03a.js.map',
          revision: '65d4bde6c8eb6d0218317089868f6735',
        },
        {
          url: '/next/_next/static/chunks/main-228488d2b5a45d9b.js',
          revision: '228488d2b5a45d9b',
        },
        {
          url: '/next/_next/static/chunks/main-228488d2b5a45d9b.js.map',
          revision: '7a14eabc10c65250e51eb55875f364a2',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-a282101c33432fdb.js',
          revision: 'a282101c33432fdb',
        },
        {
          url: '/next/_next/static/chunks/pages/_app-a282101c33432fdb.js.map',
          revision: '0292f6bf9d084d3ce4b07b325c4049b1',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-339a0915c9ebd998.js',
          revision: '339a0915c9ebd998',
        },
        {
          url: '/next/_next/static/chunks/pages/_error-339a0915c9ebd998.js.map',
          revision: 'c90820146bfdd82330194eae42ab8536',
        },
        {
          url: '/next/_next/static/chunks/pages/about-3b5563e25f759abb.js',
          revision: '3b5563e25f759abb',
        },
        {
          url: '/next/_next/static/chunks/pages/about-3b5563e25f759abb.js.map',
          revision: 'c5d7158f642c0a2351962f4069dd2ad0',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-83efe0ac651a99fd.js',
          revision: '83efe0ac651a99fd',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/groups-83efe0ac651a99fd.js.map',
          revision: 'a194129a15f2d6efba96696eb17fdd7e',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-41ea1a4cf930b43c.js',
          revision: '41ea1a4cf930b43c',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/links-41ea1a4cf930b43c.js.map',
          revision: '373ee0a28af652bf460209cfb99cb0df',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-d557f29a427bcd3a.js',
          revision: 'd557f29a427bcd3a',
        },
        {
          url: '/next/_next/static/chunks/pages/admin/users-d557f29a427bcd3a.js.map',
          revision: 'a72cf8af8f099387c1836ffb5768bfe7',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-3d3ea0a94d3660f9.js',
          revision: '3d3ea0a94d3660f9',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/forgot-3d3ea0a94d3660f9.js.map',
          revision: 'c2ff70c85e0931aaf424dbea74be2df2',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-d77dcf820a08e8bc.js',
          revision: 'd77dcf820a08e8bc',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/login-d77dcf820a08e8bc.js.map',
          revision: 'ceb0a2d97854e4f903feaa4f0fd81578',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-0d3d496061121c2f.js',
          revision: '0d3d496061121c2f',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/registration-0d3d496061121c2f.js.map',
          revision: '2c136a4fd4a21f3141cde79f70a9674b',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-3e6e600f9a6afe00.js',
          revision: '3e6e600f9a6afe00',
        },
        {
          url: '/next/_next/static/chunks/pages/auth/verification-3e6e600f9a6afe00.js.map',
          revision: 'd5cf43d2737d2b26b14ea222a4672c6f',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-7550941a24ac4e5e.js',
          revision: '7550941a24ac4e5e',
        },
        {
          url: '/next/_next/static/chunks/pages/contact-7550941a24ac4e5e.js.map',
          revision: '15a90480be0bb7b1e2e4d62abb6265a3',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-8040361c7443fefe.js',
          revision: '8040361c7443fefe',
        },
        {
          url: '/next/_next/static/chunks/pages/faq-8040361c7443fefe.js.map',
          revision: '39715b7b040923fbe2435497eb49c0ee',
        },
        {
          url: '/next/_next/static/chunks/pages/index-f5d98ccfef4ed597.js',
          revision: 'f5d98ccfef4ed597',
        },
        {
          url: '/next/_next/static/chunks/pages/index-f5d98ccfef4ed597.js.map',
          revision: 'f42fe9437664fcece994b2afa48b7cea',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-ff9d984ca0079531.js',
          revision: 'ff9d984ca0079531',
        },
        {
          url: '/next/_next/static/chunks/pages/pricing-ff9d984ca0079531.js.map',
          revision: 'f534631eb9e2b9e2012b58b45e767381',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-1aee16f0c6bf46d3.js',
          revision: '1aee16f0c6bf46d3',
        },
        {
          url: '/next/_next/static/chunks/pages/privacy-1aee16f0c6bf46d3.js.map',
          revision: '420e09b519c70f9810789aab696f859d',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-fa83ae1dc5354985.js',
          revision: 'fa83ae1dc5354985',
        },
        {
          url: '/next/_next/static/chunks/pages/user/addUrl-fa83ae1dc5354985.js.map',
          revision: '8267b945f4ca99761914c878d3e12e66',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-8ea478bb03febf48.js',
          revision: '8ea478bb03febf48',
        },
        {
          url: '/next/_next/static/chunks/pages/user/audit-8ea478bb03febf48.js.map',
          revision: 'a0218f0379601e7aec3c77c335d7d547',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-b71e6c64954103dc.js',
          revision: 'b71e6c64954103dc',
        },
        {
          url: '/next/_next/static/chunks/pages/user/billing-b71e6c64954103dc.js.map',
          revision: '258951e606c6ad8fda3b4b92cee24f5f',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-3d610c23c32639c0.js',
          revision: '3d610c23c32639c0',
        },
        {
          url: '/next/_next/static/chunks/pages/user/dashboard-3d610c23c32639c0.js.map',
          revision: 'cf77f6fb96370f3f001a89ef94ca8f06',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-55d7d5320a42c26d.js',
          revision: '55d7d5320a42c26d',
        },
        {
          url: '/next/_next/static/chunks/pages/user/integrations-55d7d5320a42c26d.js.map',
          revision: '8ec32e52854ba61a40a4d5ae6e77c805',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-37927e43b6c8962e.js',
          revision: '37927e43b6c8962e',
        },
        {
          url: '/next/_next/static/chunks/pages/user/links-37927e43b6c8962e.js.map',
          revision: 'c89ce63c91de33e0473ef20afb832ba3',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-9c2a3613e6aa2117.js',
          revision: '9c2a3613e6aa2117',
        },
        {
          url: '/next/_next/static/chunks/pages/user/profile-9c2a3613e6aa2117.js.map',
          revision: 'd076c85c59cbd85566727f825766fcff',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-93e048b6e1a33967.js',
          revision: '93e048b6e1a33967',
        },
        {
          url: '/next/_next/static/chunks/pages/user/reports-93e048b6e1a33967.js.map',
          revision: '210b6f0bacf3162fd5f64c933257b3f9',
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
          revision: '9860529928b55592610bd6a82e9e7631',
        },
        {
          url: '/next/_next/static/css/6fc13290ea89ec3c.css',
          revision: '6fc13290ea89ec3c',
        },
        {
          url: '/next/_next/static/css/6fc13290ea89ec3c.css.map',
          revision: 'd2ccf0a31d3841a0dfc9fca2e056a791',
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
      ({ sameOrigin: e, url: { pathname: a } }) =>
        !!e && !a.startsWith('/api/auth/') && !!a.startsWith('/api/'),
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
      ({ request: e, url: { pathname: a }, sameOrigin: s }) =>
        '1' === e.headers.get('RSC') &&
        '1' === e.headers.get('Next-Router-Prefetch') &&
        s &&
        !a.startsWith('/api/'),
      new e.NetworkFirst({
        cacheName: 'pages-rsc-prefetch',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ request: e, url: { pathname: a }, sameOrigin: s }) =>
        '1' === e.headers.get('RSC') && s && !a.startsWith('/api/'),
      new e.NetworkFirst({
        cacheName: 'pages-rsc',
        plugins: [
          new e.ExpirationPlugin({ maxEntries: 32, maxAgeSeconds: 86400 }),
        ],
      }),
      'GET',
    ),
    e.registerRoute(
      ({ url: { pathname: e }, sameOrigin: a }) => a && !e.startsWith('/api/'),
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
