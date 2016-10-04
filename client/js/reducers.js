import { routerReducer as routing } from 'react-router-redux';
import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';

import repos from './repos';

const rootReducer = combineReducers({
  routing,
  repos: repos.reducer,
  form: formReducer,
});

export default rootReducer;
