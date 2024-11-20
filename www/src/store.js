import { createStore } from "vuex";

export default createStore({
  state: {
    session: null, // X-CSRF-TOKEN
  },
  mutations: {
    setSession(state, data) {
      state.session = data;
    },
  },
  actions: {},
  modules: {},
});
