// actions和之前讲的Mutations功能基本一样，不同点是，
// actions是异步的改变state状态，而Mutations是同步改变状态

export default {
    addAction(context) {
        context.commit('add', 10)
    },
    reduceAction({ commit }) {
        commit('reduce')
    }
}