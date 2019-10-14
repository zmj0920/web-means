<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <icon name="flag"></icon>
    <faBell  class="svg-icon"/>
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import faBell from "../assets/svg/fa-bell.svg";
export default {
  name: "home",
  components: {
    HelloWorld,
    faBell
  }
};
</script>
<style lang="scss" scoped>
.svg-icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
