import React, { useState, useMemo } from 'react';

function Example4() {
    const [zhazhaxin, setzhazhaxin] = useState('渣渣新待客状态')
    const [zhiling, setZhiling] = useState('志玲待客状态')
    return (
        <div>
            <button onClick={() => { setzhazhaxin(new Date().getTime()) }}>渣渣新</button>
            <button onClick={() => { setZhiling(new Date().getTime() + ',志玲向我们走来了') }}>志玲</button>
            <ChildCompoent name={zhazhaxin}>{zhiling}</ChildCompoent>
        </div>
        )
    }
    
function ChildCompoent({ name, children }) {
    function zhazhaxin(name) {
        console.log("她来了，她来了。渣渣新向我们走来了")
        return name + ',渣渣新向我们走来了'
    }
    //const action = zhazhaxin(name)
      const action = useMemo(() => zhazhaxin(name),[name])
    return (
        <div>
            <div>{action}</div>
            <div>{children}</div>
        </div>
     )
}


export default Example4