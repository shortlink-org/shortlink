import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = () => new Vuex.Store({
  state: {
    links: []
  },
  mutations: {
    GET_LINKS(state, links) {
      state.links = links
    }
  },
  actions: {
    async GET_LINKS({ commit }) {
      const data = await this.$axios.$get('/api/links')
      commit('GET_LINKS', data)
    }
  }
})

export default store
