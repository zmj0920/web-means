import React, { createContext, useReducer } from 'react';
import { Counter } from './Counter'
export const CountContext = createContext()
function reducer(state, action) {
    switch (action.type) {
        case 'add':
            return { count: state.count + 1 };
        case 'sub':
            return { count: state.count - 1 };
        default:
            return state
    }
}
function Example2() {
    const initialState = { count: 0 };
    const [state, dispatch] = useReducer(reducer,initialState)
    return (
        <div>
            <p>You clicked {state.count} times</p>
            <button onClick={() => { dispatch({ type: 'add'})}}>add</button>
            <button onClick={() => { dispatch({ type: 'sub'})}}>sub</button>
            <CountContext.Provider value={state.count}>
                <Counter />
            </CountContext.Provider>
        </div>
    )
}

export default Example2;