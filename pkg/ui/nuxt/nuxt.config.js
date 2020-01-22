module.exports = {
  mode: 'spa',

  /*
  ** Environment Variables
  */
  env: {},

  /*
   ** Build configuration
   */
  build: {
    babel: {
      plugins: ['transform-vue-jsx'],
    },
    vendor: ['axios'],
  },

  buildModules: [
    '@nuxtjs/date-fns',
    '@nuxtjs/router',
  ],

  dateFns: {
    format: 'yyyy-MM-dd',
  },

  generate: {
    routes: [
      '/'
    ]
  },

  /*
   ** Global CSS
   */
  css: [
    { src: 'assets/main.css', lang: 'css' },
  ],

  /*
   ** Headers of the page
   */
  head: {
    titleTemplate: '%s - Shortlink',
    meta: [
      { charset: 'utf-8' },
      { hid: 'description', name: 'description', content: 'Shortlink service' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    ],
    link: [
      { rel: 'stylesheet', href: '//fonts.googleapis.com/css?family=Roboto:100,300,400,500,700,900' },
      { rel: 'stylesheet', href: '//cdn.jsdelivr.net/npm/@mdi/font@4.x/css/materialdesignicons.min.css' },
    ],
    htmlAttrs: {
      lang: 'en',
      dir: 'auto' // The dir attribute specifies the text direction of the element's content.
    },
  },

  /*
   ** Nuxt.js modules
   */
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/vuetify',
  ],

  vuetify: {
    optionsPath: './vuetify.options.js'
  },

  /*
   ** Axios module configuration
   */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
    baseURL: 'http://localhost:7070',
    proxyHeaders: false,
    credentials: false
  },

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],
}
