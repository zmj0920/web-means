<template>
  <div>
    <p>{{count}}</p>
    <p>{{title}}</p>
    <p>{{list}}</p>
    <p v-for="item in changePeople">{{item.age}}</p>
    <div v-dir1>自定义指令</div>
    <button @click="add(5)">+</button>
    <!-- <button @click="$store.commit('add',5)">+</button> -->
    <button @click="reduce">-</button>
    <br />
    <button @click="addAction">+</button>
    <button @click="reduceAction">-</button>
  </div>
</template>

<script>
import { mapState, mapMutations, mapGetters, mapActions } from "vuex";
export default {
  name: "vuexstore",
  data() {
    return {
      name: "123456"
    };
  },
  watch: {
    //监听器监听name值发生变化
    name(newValue, oldValue) {
      console.log("监听name");
      console.log(newValue);
    }
  },
  methods: {
    //调用
    init() {
      this.name = "123";
    },
    // 模板获取Mutations方法
    //mapMutations和mapActions使用的时候只能在methods中调用否则报错
    //$store.commit('add',5)
    ...mapMutations(["add", "reduce", "addlist"]),
    ...mapActions(["addAction", "reduceAction", "actionslist"]),
    addList() {
      // this.$store.commit("addlist", { name: "zmj" });
      this.addlist({ name: "zmj1" });
    },
    actionsList() {
      this.actionslist({ name: "123" });
    }
  },
  created() {},
  computed: {
    // mapState和mapGetter的使用只能在computed计算属性中
    //1.通过mapState的对象来赋值
    //this.$store.state.count
    ...mapState(["count", "title", "list"]),

    //2.通过mapState的对象来赋值
    // ...mapState({
    //   count: state => state.count
    // }),

    //this.$store.getters.count
    ...mapGetters(["count", "changePeople"])
  },
  mounted() {
    this.init();
    this.addList();
    //异步执行
    this.actionsList()
  },
  directives: {
    // 指令名称
    dir1: {
      inserted(el) {
        // 指令中第一个参数是当前使用指令的DOM
        console.log(el);
        console.log(arguments);
        // 对DOM进行操作
        el.style.color = "red";
      }
    }
  }
};
</script>
