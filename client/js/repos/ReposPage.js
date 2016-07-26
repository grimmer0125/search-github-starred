import React from 'react';
import FetchingStatus from './constants';
import { connect } from 'react-redux';
// import { bindActionCreators } from 'redux';

import {
  FETCH_STARRRED_STATUS,
} from './actionTypes';


class ReposPage extends React.Component {

  componentWillMount() {
    // const { businessId, getBillingData } = this.props;
    if (!this.hasData()) {
      this.startPoll();
    //   getBillingData(businessId);
    }
  }

  componentDidMount() {
    // const fetchAction = bindActionCreators(fetchDevices, this.props.dispatch);
    // fetchAction();
  }


  componentWillReceiveProps(nextProps) {
    console.log('componentWillReceiveProps');
    if (this.props.repos !== nextProps.repos) {
      // clearTimeout(this.timeout);
      console.log('different Props');

      // Optionally do something with data

      // if (!nextProps.isFetching) {
      //   this.startPoll();
      // }
    }
  }

  // const type = 'DELETE_DEVICE';
  // this.props.dispatch({ type, payload: { sn } });
  // dispatch({type: 'USER_FETCH_REQUESTED', payload: {userId}})

//  this.timeout = setTimeout(() => this.props.dataActions.dataFetch(), 15000);
  startPoll() {
    const { dispatch } = this.props;
    console.log('0 timer runs !!!');

    dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });

    // this.timeout = setTimeout(() => {
    //   console.log('timer runs !!!');
    //   debugger;
    //   dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });
    // }, 2000);
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
    let statusStr = '';
    console.log('status in render;', this.props);
    const { fetchingStatus } = this.props.repos;

    switch (fetchingStatus) {
      case FetchingStatus.NOTSTART:
        statusStr = 'Indexing is not started yet';
        break;
      case FetchingStatus.FETCHING:
        statusStr = 'Indexing, wait a moement...';
        break;
      case FetchingStatus.FETCHED:
        statusStr = 'Indexing is ok...';
        break;
      default:
        statusStr = 'unknown status';
    }

    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>

        <div className="loading-text">
          {statusStr}
        </div>
      </div>
    );
  }

  render() {
    if (this.hasData()) {
      console.log('has data');
      return this.renderComponents();
    }

    console.log('has no data');

    return this.renderLoadingScreen();
  }
}

ReposPage.propTypes = {
  repos: React.PropTypes.object,
  dispatch: React.PropTypes.func,
};

function test(state) {
  console.log('test test');
  return state.repos;
}

export function mapStateToProps(state) {
  return {
    repos: test(state),
    // businessId: state.business.id,
  };
}


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

export default connect(mapStateToProps)(ReposPage);
