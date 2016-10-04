import React from 'react';
import { Router, Route, browserHistory } from 'react-router';
import { Provider } from 'react-redux';
import { syncHistoryWithStore } from 'react-router-redux';

import configureStore from '../app/configureStore.js';
import Repos from '../repos';

const store = configureStore();
const history = syncHistoryWithStore(browserHistory, store);

export default function Root() {
  return (
    <Provider store={store}>
      <div className="layout-column-start-center" style={{ width: '100%' }}>
        <Router history={history}>
          <Route path="/" component={Repos.component} />
        </Router>
      </div>
    </Provider>
  );
}
