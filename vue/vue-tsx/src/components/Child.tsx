import { Component, Vue, Prop, PropSync } from 'vue-property-decorator';

@Component
export default class Child extends Vue {

  @Prop({ required: false })


  @PropSync("name", { type: String }) syncedName!: string;

  public msg?: string;

  protected render() {
    return (
      <div>
        {this.msg}
        <p>{ this.syncedName }</p>
      </div>
    );
  }
}