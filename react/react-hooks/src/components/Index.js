import React, { useEffect } from 'react';

function Index() {
    //但是当传空数组[]时，就是当组件将被销毁时才进行解绑
    useEffect(() => {
        console.log('useEffect=>老弟你来了！Index页面')
        return () => {
            console.log('老弟，你走了!Index页面')
        }
    },[])
    return <h2>JSPang.com</h2>;
}

export default Index;