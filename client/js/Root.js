// import getMuiTheme from 'material-ui/styles/getMuiTheme';
// import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

import React from 'react';
// import ReactDOM from 'react-dom';

// import { Router, Route, hashHistory, IndexRoute } from 'react-router';
// import { Provider } from 'react-redux';
// import { syncHistoryWithStore } from 'react-router-redux';
//
// import DeviceDetailPage from './DeviceDetailPage';
// import MainPage from './MainPage';
// import { baseExoTheme } from '../src/js/theme';
// import { initElfApp } from '../actions/deviceAction.js';

// import InfoContainer from '../containers/InfoContainer.js';
// import DefinitionContainer from '../containers/DefinitionContainer.jsx';
// import DevicesList from '../containers/DevicesList.js';
// import TokenContainer from '../containers/TokenContainer.jsx';

// import Resources from '../containers/Resources.jsx';
// import configureStore from '../store/configureStore.js';
//
// const store = configureStore();
// const history = syncHistoryWithStore(hashHistory, store);

// const breadcrumbs = {
//   product: [
//   { name: 'Product', specificURL: '' },
//   { name: 'ProductName', alias: '' }],
//
//   detail: [
//   { name: 'Product', specificURL: '' },
//   { name: 'ProductName', alias: '', specificURL: '/devices' },
//   { name: 'DeviceSN', alias: '' }],
// };

export class Root extends React.Component {

  componentWillMount() {
    // const { businessId, getBillingData } = this.props;
    // if (!this.hasData()) {
    //   getBillingData(businessId);
    // }
  }

  componentDidMount() {
    // const fetchAction = bindActionCreators(fetchDevices, this.props.dispatch);
    // fetchAction();
  }

  hasData() {
    return false;
    // const { billing } = this.props;
    // return !billing.error &&
    //   !billing.fetchingData &&
    //   billing.accountManager &&
    //   billing.balance &&
    //   billing.creditCard &&
    //   billing.tier;
  }

  renderComponents() {
    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        hello world
      </div>
    );
  }

  renderLoadingScreen() {
    // use a common loading screen
    return <div className="page">Loading</div>;
  }

  render() {
    if (this.hasData()) return this.renderComponents();

    return this.renderLoadingScreen();
  }
}


// const muiTheme = getMuiTheme(baseExoTheme);

// export default function Root() {
//   return (
//     <div><a href="/logout">sign out2</a></div>
//   );
// }

/* <MuiThemeProvider muiTheme={muiTheme}>
  <Provider store={store}>
    <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
      <Router history={history}>
        <Route path="/" component={MainPage} breadcrumbs={breadcrumbs.product}>
          <IndexRoute component={InfoContainer} />
          <Route path="info" component={InfoContainer} />
          <Route path="definition" component={DefinitionContainer} />
          <Route path="devices" component={DevicesList} />
        </Route>
        <Route path="/detail" component={DeviceDetailPage} breadcrumbs={breadcrumbs.detail}>
          <IndexRoute component={Resources} />
          <Route path="resources" component={Resources} />
          <Route path="tokens" component={TokenContainer} />
        </Route>
      </Router>
    </div>
  </Provider>
</MuiThemeProvider>

store.dispatch(initElfApp());*/

// References:
// http://zhuanlan.zhihu.com/purerender/20381597
// https://github.com/reactjs/react-router/blob/master/upgrade-guides/v2.0.0.md#using-history-with-router
// https://medium.com/@slashtu/react-router-redux-62872860e8a#.i1u4g2d9f

// Notice:
// browserHistory.push('/preview/info'); still need add /preview
// <Link to="/info">About</Link>
// <Route path="/preview/editor/module/:id">  still need add "/preview", //https://github.com/reactjs/react-router/issues/2261 ->check later.
// ownProps.params.id,  provides route information via a route component's props.
