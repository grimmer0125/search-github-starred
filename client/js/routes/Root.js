// import getMuiTheme from 'material-ui/styles/getMuiTheme';
// import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
// import { baseExoTheme } from '../src/js/theme';
// const muiTheme = getMuiTheme(baseExoTheme);

// IndexRoute
import React from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { Provider } from 'react-redux';
import { syncHistoryWithStore } from 'react-router-redux';

import configureStore from '../app/configureStore.js';
import MainPage from '../main/MainPage';

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

export default function Root() {
  return (
    <Provider store={store}>
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        <Router history={history}>
          <Route path="/" component={MainPage} />
        </Router>
      </div>
    </Provider>
  );
}
