import React, { useContext} from 'react';
import {CountContext} from './Example2'
export function Counter() {
    const count = useContext(CountContext)  //一句话就可以得到count
    return (<h2>{count}</h2>)
}