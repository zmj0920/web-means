import { Observable } from 'rxjs'
//OnSubscribe正在订阅
// const onSubscribe = observer => {
//     observer.next(1);
//     observer.error("渣渣新太渣");
//     observer.complete();
// }

// const source$ = new Observable(onSubscribe)

// const theObserver = {
//     next: item => console.log(item),
//     error: err => console.log(err),
//     complete: () => console.log("No More Data")

// }
// source$.subscribe(theObserver)

const onSubscribe=observer=>{
    let number=1;
    const handel=setInterval(()=>{
        observer.next(number++)
    },1000);
    return {
        unsubscribe:()=>{
            clearInterval(handel)
        }
    }
}

const source$=new Observable(onSubscribe)

const subscrption=source$.subscribe(item=>console.log(item));

setTimeout(()=>{
    subscrption.unsubscribe()
},3500)


