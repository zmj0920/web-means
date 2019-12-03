<template>
  <div>
    <h2>Child</h2>
    <p>data: {{title}}</p>
    <div>
      <input v-model="title" />
    </div>
    <div>
      Watch:
      <input type="text" v-model="inputValue" />
    </div>
    <div>
      Model:
      <input type="checkbox" :checked="checked" @change="change" />
    </div>
    <p>这是从父组件中传过来的值Prop: {{msg}}</p>
    <p>name{{name}}</p>
    <div>来自provide中的值--1--{{foo}}--2--{{foo}}</div>
    <p>computed msg: {{computedMsg}}</p>
    <button @click="greet">Greet</button>
    <button @click="triggerEmit('qqq')">触发emit</button>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import {
  Component,
  Prop,
  PropSync,
  Emit,
  Watch,
  Model,
  Inject,
  Provide
} from "vue-property-decorator";

@Component
export default class Child extends Vue {
  //default默认值
  //感叹号是非null和非undefined的类型断言
  @Prop({ default: "msg默认值" }) private msg!: string;

  @PropSync("name", { type: String })
  syncedName!: string;

  public setSyncedName() {
    this.syncedName = "delete";
  }

  private inputValue: string = "";

  private title: string = "君吟";

  @Watch("inputValue")
  valueChange(newValue: string, oldValue: string) {
    console.log(`新值${newValue} 旧值  ${oldValue}`);
  }
  // 对watch的配置为第二个参数，以对象形式传入
  // @Watch('inputValue',{ deep: true })
  // valueChange(newValue: string, oldValue: string) {
  //   console.log(newValue, oldValue)
  // }

  @Model("change", { type: Boolean }) readonly checked!: boolean;

  @Emit("change")
  change() {
    // alert(e.isTrusted);
   //  console.log(e.target.checked);
  }

  @Emit("demo-log")
  triggerEmit(n: string) {
    alert(n);
  }

  @Provide()
  foo = "fooss";

  // @Inject()
  // foo!: string;

  mounted() {
    //this.greet()
    this.$on("demo-log", (data: string): void => {
      console.log("hhhh");
    });
  }

  // computed
  get computedMsg() {
    return "computed " + this.msg;
  }

  // 函数
  greet() {
    alert("greeting: " + this.msg);
  }
}
</script>