
import { Component, Vue, Watch } from 'vue-property-decorator';
import style from './test.module.scss';
import Child from '@components/Child.tsx';

@Component
export default class Parent extends Vue {

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
   render() {
    return (
      <div class={style.class1}>
        <input type='text' v-model={this.value} />
        {this.valueLength}
        <button class='class2' onClick={this.handleClick}>
          点我
        </button>
        <div>
          <Child msg='渲染组件成功' />
        </div>
      </div>
    );
  }
}
