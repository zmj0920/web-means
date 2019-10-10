import {Observable} from 'rxjs'
import 'rxjs/add/observable/of'; //导入实例操作符 

// const source$=Observable.of(1,2,3)
// source$.subscribe(console.log)
//Subject (主体)
//Observer 观察者  subscribe (订阅)
//Observable（被观察着）=Publisher（发布者）+Iterator(迭代器)
//OnSubscribe正在订阅
// const OnSubscribe=observer=>{
//   observer.next(1);
//   observer.next(2);
//   observer.next(3);
// }
//source 来源
// const source$=new Observable(OnSubscribe)

// const theObserver={
//   next:item=>console.log(item)
// }

// source$.subscribe(theObserver)


const OnSubscribe=observer=>{
  let number=1;
  const handel=setInterval(()=>{
    observer.next(number++);
    if(number>3){
      clearInterval(handel);
      observer.complete()
    }
  },1000)
}
//complete 完成
const source$=new Observable(OnSubscribe)
const theObserver={
  next:item=>console.log(item),
  complete:()=>console.log("No More Data")
}

source$.subscribe(theObserver)

