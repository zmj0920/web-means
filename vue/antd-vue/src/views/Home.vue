<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
     <a-button type="primary">hello world</a-button>
     <a href="">小黑</a>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import { Button } from 'ant-design-vue';
export default {
  name: 'home',
  components: {
    HelloWorld,
    AButton:Button
  }
}
</script>
<style lang="less">
@link-color:red;
a,
.link {
  color: @link-color;
}
</style>
