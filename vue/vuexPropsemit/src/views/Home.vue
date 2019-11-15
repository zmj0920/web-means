<template>
  <div class="home">
    <Parent message="Welcome to Your Vue.js App"/>
  </div>
</template>
<script>
import Parent from '@/components/Parent.vue'
export default {
  name: 'home',
  components: {
    Parent
  }
}
</script>
