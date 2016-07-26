import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';

import repos from './repos';

// import * as devicesReducers from './devicesSN';

// ...devicesReducers,
const rootReducer = combineReducers({
  routing,
  repos: repos.reducer,
});

export default rootReducer;
