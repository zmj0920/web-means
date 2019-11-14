<template>
  <div class="hello">
    <p>{{ message }}</p>
    <child v-bind:messages="childs" v-on:toChildEvent="mest"></child>
  </div>
</template>

<script>
import child from "./child";
export default {
  components: {
    child
  },
  data() {
    return {
      childs:'Hello World',
      message: " "
    };
  },
  methods: {
    mest(data) {
      this.message=data;
      console.log(data);
    }
  },
  created() {
   
  },
  mounted() {
   
  }
};
</script>
<style scoped lang="scss">
</style>
