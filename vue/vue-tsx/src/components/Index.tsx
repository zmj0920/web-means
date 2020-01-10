
import { Component, Vue, Watch } from 'vue-property-decorator';

import style from './test.module.scss';

import Child from './Child.tsx';

@Component

export default class Index extends Vue {
  public value = '';
  public msg = '';

  public get valueLength() {
    return this.value.length;
  }

  public created() {
    console.log('我在组件创建时被调用');
  }

  public handleClick() {
    console.log('我被点了');
  }

  @Watch('value')
  protected valueWatch(newV: any, oldV: any) {
    this.msg = `监听到属性value发生变化，新的值为：${newV}`;
  }

  handleUpdate(){
    
  }
  protected render() {
    return (
      <div class={style.class1}>
          {this.$route.path}
        <input type='text' v-model={this.value} />
        {this.valueLength}
        <button class='class2' onClick={this.handleClick}>
          点我
        </button>
        <div>
     
          <Child  props={{msg:'传值'}}  on={{['update:name']:this.handleUpdate}}   />
        </div>
      </div>
    );
  }
}
