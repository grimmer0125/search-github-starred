import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';

// import * as devicesReducers from './devicesSN';


const rootReducer = combineReducers({
  routing, //...devicesReducers,
});

export default rootReducer;
