import Vue from 'vue'
import Router from 'vue-router'

import LinksPage from '~/components/pages/links'
import AboutPage from '~/components/pages/about'

Vue.use(Router)

export function createRouter() {
  return new Router({
    mode: 'history',
    routes: [
      {
        path: '/',
        component: LinksPage
      },
      {
        path: '/about',
        component: AboutPage
      }
    ]
  })
}
