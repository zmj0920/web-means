<template>
  <div>
    <header>子组件页面</header>
    <div>接收传值：{{messages}}</div>
    <button @click="emitParent">$emit向父组件传值</button>
  </div>
</template>
<script>
export default {
  props: ["messages"],
  methods: {
    emitParent() {
      this.$emit("toChildEvent", "子组件发送数据到父组件");
    }
  }
};
</script>