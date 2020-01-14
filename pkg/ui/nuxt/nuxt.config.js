module.exports = {
  mode: 'spa',

  /*
   ** Build configuration
   */
  build: {
    babel: {
      plugins: ['transform-vue-jsx'],
    },
    vendor: ['axios', 'vue-material'],
  },

  /*
   ** Global CSS
   */
  css: [
    { src: 'assets/main.css', lang: 'css' },
    { src: 'vue-material/dist/vue-material.min.css', lang: 'css' },
    { src: '~/assets/theme.scss', lang: 'scss' }, // include vue-material theme engine
    { src: 'element-ui/lib/theme-chalk/reset.css', lang: 'css' },
    { src: 'element-ui/lib/theme-chalk/index.css', lang: 'css' },
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
      { rel: 'stylesheet', href: '//fonts.googleapis.com/css?family=Roboto' },
      { rel: 'stylesheet', href: '//fonts.googleapis.com/css?family=Roboto:400,500,700,400italic|Material+Icons' },
    ],
  },

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    { src: '~/plugins/vue-material' },
    { src: '@/plugins/element-ui' },
  ],
}
