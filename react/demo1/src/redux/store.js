import {createStore} from 'redux';

import Reducer from './reducer';

 const store=createStore(Reducer) //创建数据仓库

export default store //暴露仓库