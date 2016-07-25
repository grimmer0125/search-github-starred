import React from 'react';

export default class MainPage extends React.Component {

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
    // className="page"
    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>

        <div className="activatedColor">
          Loading
        </div>
      </div>
    );
  }

  render() {
    if (this.hasData()) {
      console.log("has data");
      return this.renderComponents();
    }

    console.log("has no data");

    return this.renderLoadingScreen();
  }
}
