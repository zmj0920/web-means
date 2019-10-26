import React from 'react';
import { connect } from 'react-redux'
import {ADD, CHANGEINPUT} from './store/actionTypes'
const TodoList = (props) => {
    let { inputValue, inputChange, clickButton, list } = props;
    return (
        <div>
            <div>
                <input value={inputValue} onChange={inputChange} />
                <button onClick={clickButton}>提交</button>
            </div>
            <ul>
                {
                    list.map((item, index) => {
                        return (<li key={index}>{item}</li>)
                    })
                }
            </ul>
        </div>
    );
}

const stateToProps = (state) => {
    return {
        inputValue: state.inputValue,
        list: state.list
    }
}

const dispatchToProps = (dispatch) => {
    return {
        inputChange: (e) => dispatch({
            type: CHANGEINPUT,
            value: e.target.value
        }),
        clickButton: () => dispatch({
            type: ADD
        })
    }
}
export default connect(stateToProps, dispatchToProps)(TodoList);