<template>
  <div>
    <input type="datetime" name="datetime" v-model="local" />
    <button @click="getLocalData">获取当前时间</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      local: ""
    };
  },
  methods: {
    getLocalData() {
      var date = new Date();
      this.local = date.toLocaleString();
    }
  }
};
</script>