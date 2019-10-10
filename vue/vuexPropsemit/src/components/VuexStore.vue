<template>
  <div>
    <p>{{this.$store.state.count}}---{{count}}</p>
    <div v-dir1>自定义指令</div>
    <button @click="$store.commit('add',5)">+</button>
    <button @click="reduce">-</button>
    <br />
    <button @click="addAction">+</button>
    <button @click="reduceAction">-</button>
  </div>
</template>

<script>
import { mapState, mapMutations, mapGetters, mapActions } from "vuex";
export default {
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
    ...mapMutations(["add", "reduce"]),
    ...mapActions(["addAction", "reduceAction"])
  },
  computed: {
    //1.通过mapState的对象来赋值
    ...mapState(["count"]),
     //2.通过mapState的对象来赋值
    // ...mapState({
    //   count: state => state.count
    // }),
    ...mapGetters(["count"])
  },
  mounted() {
    this.init();
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
