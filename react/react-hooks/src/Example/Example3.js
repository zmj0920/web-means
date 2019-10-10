import React from 'react';
import ShowArea from './ShowArea';
import Buttons from './Buttons';
import { Color } from './Color';   //引入Color组件

function Example3(){
    return (
        <div>
            <Color>
                <ShowArea />
                <Buttons />
            </Color>
            
        </div>
    )
}

export default Example3