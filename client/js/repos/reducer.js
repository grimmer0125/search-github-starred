import FetchingStatus from './constants';

import {
  GET_REPOS_STATUS_SUCCEEDED,
} from './actionTypes';


const initialState = {
  error: null,
  fetchingStatus: FetchingStatus.NOTSTART,
  numStarred: -1,
};

export default function repos(state = initialState, action) {
  console.log('into reducer:', action.type);
  switch (action.type) {
    case GET_REPOS_STATUS_SUCCEEDED:
      console.log('papyload:', action.payload.status);
      const status = action.payload.status;
      return { ...state, fetchingStatus: status };
      // return Object.assign({}, state, { status });
    default:
      return state;
  }
}

// import * as ActionTypes from '../actions/deviceAction';

//     case ActionTypes.getDeviceInfo:
//       return {
//         isFetching: false,
//         items: updateSomeDeviceNameAliases(
//           state,
//           action.payload.rid,
//           action.payload.name,
//           action.payload.aliases
//         ),
//       };
//     default:
//       return state;
//   }
// }
