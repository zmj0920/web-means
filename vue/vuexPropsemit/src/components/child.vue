<template>
  <div>
    <header>子组件页面</header>
    <div>{{messages}}</div>
    <button @click="emitParent">向父组件传值</button>
  </div>
</template>
<script>
export default {
  props: ["messages"],
  methods: {
    emitParent() {
      this.$emit("toChildEvent", "我是子组件");
    }
  }
};
</script>