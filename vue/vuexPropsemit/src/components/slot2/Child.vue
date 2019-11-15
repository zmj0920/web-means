<template>
  <div class="child">
    <p>up</p>
    <slot name="up"></slot>
    <h3>这里是子组件</h3>
    <p>down</p>
    <slot name="down"></slot>
    <p>未指定名字</p>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    data: function(){
      return {
        msg: ''
      }
    },
    computed: {

    },
    methods:{

    },
    components: {

    }
  }
</script>

<style scoped>
  .child{
    background-color: #00bbff;
    width: 100%;
    padding: 10px;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
</style>
