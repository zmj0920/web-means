//过滤

export default {
    count: function (state) {
        return state.count += 100;
    },
    changePeople: (state) => {
        return state.people.filter(item => {
            if (item.age > 30) {
                return true
            }
        })
    }
}