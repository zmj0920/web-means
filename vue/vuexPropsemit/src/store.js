import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)
export default new Vuex.Store({
  state: { //定义状态  在模板中可以用$store.state.count输出
    count: 0
  },
  mutations: { //Mutations修改状态 改变state的计算方法
    //在组件中调用方式
    // <button @click="$store.commit('add')">+</button>
    // <button @click="$store.commit('reduce')">-</button>

    add(state, n) {
      state.count += n;
    },
    reduce(state) {
      state.count--;
    }
  },
  // actions和之前讲的Mutations功能基本一样，不同点是，
  // actions是异步的改变state状态，而Mutations是同步改变状态
  actions: { //异步计算
    addAction(context) {
      context.commit('add', 10)
    },
    reduceAction({ commit }) {
      commit('reduce')
    }
  },
  getters: { //过滤
    count: function (state) {
      return state.count += 100;
    }
  },
  modules:{
    // const moduleA = {
    //   state: { ... },
    //   mutations: { ... },
    //   actions: { ... },
    //   getters: { ... }
    //  }
     //当状态过多时可以把每个模块用module来管理
    // moduleA
  }
})
