import FetchingStatus from './constants';

import {
  GET_REPOS_STATUS_SUCCEEDED,
} from './actionTypes';


const initialState = {
  error: null,
  fetchingStatus: FetchingStatus.NOTSTART,
  numOfStarred: -1,
  githubAccount: '',
};

export default function repos(state = initialState, action) {
  console.log('into reducer:', action.type);
  switch (action.type) {
    case GET_REPOS_STATUS_SUCCEEDED:
      // console.log('papyload:', action.payload.status);
      const result = action.payload.result;
      return { ...state, fetchingStatus: result.status, numOfStarred: result.numOfStarred, githubAccount: result.githubAccount };
      // return Object.assign({}, state, { status });
    default:
      return state;
  }
}

// import * as ActionTypes from
//       return {
//         isFetching: false,
//         items:
