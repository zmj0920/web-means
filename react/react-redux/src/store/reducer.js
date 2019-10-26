import {ADD, CHANGEINPUT} from './actionTypes'


const defalutState = {
    inputValue : '技术菜',
    list :[]
}



export default (state = defalutState,action) =>{
    if(action.type === CHANGEINPUT){
        let newState = JSON.parse(JSON.stringify(state))
        newState.inputValue = action.value
        return newState
    }
    if(action.type === ADD){
        let newState = JSON.parse(JSON.stringify(state))
        newState.list.push(newState.inputValue)
        newState.inputValue = ''
        return newState
    }
    return state
}