module.exports = {
  build: {
    babel: {
      plugins: ['transform-vue-jsx'],
    },
    vendor: ['axios'],
  },

  css: [
    'assets/main.css',
  ],

  head: {
    titleTemplate: '%s - Shortlink',
    meta: [
      { charset: 'utf-8' },
      { hid: 'description', name: 'description', content: 'Shortlink service' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    ],
    link: [
      { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css?family=Roboto' },
    ],
  },
}
