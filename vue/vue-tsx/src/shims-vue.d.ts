declare module '*.vue' {
  import Vue from 'vue'
  export default Vue
}

declare module "*.scss" {
  const scss: any;
  export default scss;
}

declare module "*.tsx" {
  const tsx: any;
  export default tsx;
}


// declare module "vue/types/options" {
//   interface ComponentOptions<V extends Vue> {
//     [propName: string]: any;
//     ref?: string;
//   }
// }
