import React from 'react';
import FetchingStatus from './constants';
import { connect } from 'react-redux';
// import { bindActionCreators } from 'redux';

class ReposPage extends React.Component {

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
    // const { starredRepos } = this.props;
    const starredRepos = this.props.repos;
    return (!starredRepos.error && starredRepos.fetchingStatus === FetchingStatus.FETCHED);
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

        <div className="loading-text">
          Indexing, wait a moement...
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

ReposPage.propTypes = {
  starredRepos: React.PropTypes.object,
};

function test(state){
  console.log("into test");
  return state.repos;
}

export function mapStateToProps(state) {
  return {
    repos: test(state),
    // businessId: state.business.id,
  };
}

export default connect(mapStateToProps)(ReposPage);

//
// export function mapDispatchToProps(dispatch) {
//   return bindActionCreators({
//     getBillingData: actions.getBillingData,
//     openChangeCreditCardModal: changeCreditCard.actions.openChangeCreditCardModal,
//     openMakePaymentModal: makePayment.actions.openMakePaymentModal,
//   }, dispatch);
// }
//
// export default connect(mapStateToProps, mapDispatchToProps)(BillingPage);
