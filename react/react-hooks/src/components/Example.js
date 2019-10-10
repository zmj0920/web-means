import React, { useState, useEffect  } from 'react';
function Example() {
    const [count, setCount] = useState(0) //数组结构
    //useState的用法，分别是声明、读取、使用（修改）
    const [age, setAge] = useState(18)
    const [sex] = useState('男')
    const [work] = useState('前端程序员')
    //useEffect函数来代替生命周期函数
    //React首次渲染和之后的每次渲染都会调用一遍useEffect函数
    //useEffect中定义的函数的执行不会阻碍浏览器更新视图异步执行的
    //useEffect可以解决比如我们的定时器要清空，避免发生内存泄漏;比如登录状态要取消掉，避免下次进入信息出错
    useEffect(() => {
        console.log(`useEffect=>You clicked ${count} times`)
        return ()=>{
            console.log('====================')
        }
    },[count])
    return (
        <div>
            <p>渣渣新 今年:{age}岁</p>
            <p>性别:{sex}</p>
            <p>工作是:{work}</p>
            <p>You clicked {count} times</p>
            <button onClick={() => { setCount(count + 1) }}>count</button>
            <button onClick={() => { setAge(age + 1) }}>age</button>
        </div>
    );
}

export default Example;