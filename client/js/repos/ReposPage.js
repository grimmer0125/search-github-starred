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
    }
  }

  componentDidMount() {
    // const fetchAction = bindActionCreators(fetchDevices, this.props.dispatch);
    // fetchAction();
  }


  componentWillReceiveProps(nextProps) {
    console.log('componentWillReceiveProps');
    if (this.props.repos !== nextProps.repos) {
      console.log('different Props');


      // Optionally do something with data

      // if (!nextProps.isFetching) {
      //   this.startPoll();
      // }
    }
    clearTimeout(this.timeout);

    if (!this.hasData()) {
      console.log('start timer in receive props ');
      this.startPoll();
    }
  }

  // const type = 'DELETE_DEVICE';
  // this.props.dispatch({ type, payload: { sn } });
  // dispatch({type: 'USER_FETCH_REQUESTED', payload: {userId}})

  // ref: http://notjoshmiller.com/ajax-polling-in-react-with-redux/
//  this.timeout = setTimeout(() => this.props.dataActions.dataFetch(), 15000);
  startPoll() {
    const { dispatch } = this.props;
    // console.log('status timer runs !!!');

    // dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });

    this.timeout = setTimeout(() => {
      console.log('timer runs !!!');
      // debugger;
      dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });
    }, 2000);
  }

  hasData() {
    // const { starredRepos } = this.props;
    const starredRepos = this.props.repos;
    return (!starredRepos.error && starredRepos.fetchingStatus === FetchingStatus.INDEXED);
  }

  renderComponents() {
    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        Indexing is finished, start to query.
      </div>
    );
  }

  renderLoadingScreen() {
    // use a common loading screen
    // className="page"
    let statusStr = '';
    // console.log('status in render;', this.props);
    const { fetchingStatus, numOfStarred } = this.props.repos;

    switch (fetchingStatus) {
      case FetchingStatus.NOTSTART:
        statusStr = 'Fetching is not started yet';
        break;
      case FetchingStatus.FETCHING:
        statusStr = 'It is fetching, wait a moement...';
        break;
      case FetchingStatus.INDEXING:
        if (numOfStarred > 0) {
          statusStr = 'It is indexing ' + numOfStarred + ', wait a mement...';
        } else {
          statusStr = 'It is indexing, wait a mement...';
        }
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
  };
}

// export function mapDispatchToProps(dispatch) {
//   return bindActionCreators({
//     xxx: actions.xxx

//   }, dispatch);
// }
//
export default connect(mapStateToProps)(ReposPage);
