<template>
    <div class="father">
        <h3>这里是父组件</h3>
        <child>
           <template v-slot:header="user">
               <li v-for="item in user.data">{{item}}</li>
           </template>
        </child>
    </div>
</template>

<script>
  import Child from './Child.vue'
    export default {
        data: function () {
            return {
                msg: ''
            }
        },
        components:{
          'child': Child
        }
    }
</script>

<style scoped>
  .father{
    width:100%;
    background-color: #ccc;
    height: 650px;
  }
  .tmpl{
    display: flex;
    justify-content: space-around;
    flex-direction: row;
    width: 30%;
    margin: 0 auto;
  }
  .tmpl span{
    border:1px solid red;
   
  }
</style>
