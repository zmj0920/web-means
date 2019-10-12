<template>
  <div>
    <Child msg="Hello World"   @update:name="handleUpdate"  />
  </div>
</template>

<script>
import Vue from "vue";
import { Component, Prop, Provide } from "vue-property-decorator";
import Child from "./Child";

@Component({
  components: {
    Child
  }
})
export default class Parent extends Vue {}
</script>