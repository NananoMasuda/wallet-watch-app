// src/store/index.js
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: ""
  },
  mutations: {
    saveToken (state,token) {
      state.token = token
    },
    removeToken(state){
      state.token = ""
    }
  },
  actions: {
    saveToken({commit},token){
      commit("saveToken",token)
    },
    removeToken({commit}){
      commit("removeToken")
    }
  },
  getters: {
    getToken(state){
      return state.token
    }
  }
})