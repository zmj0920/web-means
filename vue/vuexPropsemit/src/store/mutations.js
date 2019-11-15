//Mutations修改状态 改变state的计算方法 都是同步事务
//在组件中调用方式
// <button @click="$store.commit('add')">+</button>
// <button @click="$store.commit('reduce')">-</button>


export default {
    add(state, n) {
        state.count += n;
    },
    reduce(state) {
        state.count--;
    },
    addlist(state,data) {
         state.list.push(data)
    }
}