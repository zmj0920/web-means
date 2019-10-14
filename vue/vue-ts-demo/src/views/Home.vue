<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <!-- <HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/> -->
  </div>
</template>

<script lang="ts">
import { Component, Vue ,Model,Emit} from 'vue-property-decorator';


@Component({
  components: {

  },
})
export default class Home extends Vue {}
</script>
