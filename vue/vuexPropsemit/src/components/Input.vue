<template>
  <div>
    <p>
      文本
      <input v-model="text" placeholder="edit me" />
      {{text}}
    </p>
    <p>
      多行文本
      <textarea v-model="textarea" placeholder="add multiple lines"></textarea>
      {{textarea}}
    </p>

    <p>
      复选框:
      <input type="checkbox" id="checkbox" v-model="checkbox" />
      <label for="checkbox">{{ checkbox }}</label>
    </p>

    <p>
      多个复选框:
      <input type="checkbox" id="jack" value="Jack" v-model="checkedNames" />
      <label for="jack">Jack</label>
      <input type="checkbox" id="john" value="John" v-model="checkedNames" />
      <label for="john">John</label>
      <input type="checkbox" id="mike" value="Mike" v-model="checkedNames" />
      <label for="mike">Mike</label>
      <br />
      <span>Checked names: {{ checkedNames }}</span>
    </p>
    <p>
      单选按钮:
      <input type="radio" id="one" value="One" v-model="picked" />
      <label for="one">One</label>
      <input type="radio" id="two" value="Two" v-model="picked" />
      <label for="two">Two</label>
      <br />
      <span>Picked: {{ picked }}</span>
    </p>

    <p>
      下拉菜单：
      <select v-model="selected">
        <option disabled value>请选择</option>
        <option>A</option>
        <option>B</option>
        <option>C</option>
      </select>
      <span>Selected: {{ selected }}</span>
    </p>

    <p>
      ref使用 ：
      <input type="text" ref="msgText" v-model="msg" />
      <button @click="getElement">获取元素值</button>
    </p>
    <p>
     ref获取子组件的input值  <RefInput ref="childComponent"></RefInput>
    </p>
  </div>
</template>
<script>
import RefInput from "./RefInput";
export default {
  data() {
    return {
      text: "",
      textarea: "",
      checkbox: "",
      checkedNames: [],
      picked: "",
      selected: "",
      msg: "Hello ref"
    };
  },
  components:{
    RefInput
  },
  mounted() {
    console.log("mounted: " + this.$refs.msgText.value);
  },
  methods: {
    getElement() {
      console.log(this.$refs.msgText.value);
      console.log(
        "子组件 input 输入框的值为：" + this.$refs.childComponent.local
      );
    }
  }
};
</script>