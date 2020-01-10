import { Component, Vue,  } from 'vue-property-decorator';
import Index from '../components/Index.tsx';
@Component
export default class Home extends Vue {
   render() {
    return (
    <div>
      <Index />
    </div>);
  }
}