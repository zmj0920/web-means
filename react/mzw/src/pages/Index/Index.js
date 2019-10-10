import React, { useEffect, useState } from 'react';
import $http from '../../api/api.js'
// function reducer(state, action) {
//     switch (action.type) {
//         case 'del':
//             return console.log(state.uid);
//         default:
//             return state
//     }
// }
function Index() {
     //const [state, dispatch] = useReducer(reducer,data)
    const [data, setData] = useState([]);
    useEffect(() => {
        $http.getData().then(res => {
            setData(res);
        })
    }, [])
    const del=(index)=>{
        console.log(index)
        setData(data.splice(1,index))  
    }
    return (
        <div>
            {
                <ul>
                    {data.map((item,index) => (
                        <li key={item.uid}>
                            <a href={item.uid}>{item.uname}</a>
                            <button onClick={()=>{del(index)}}>删除</button>
                        </li>
                    ))}
                </ul>
            }
        </div>
    )
}
export default Index