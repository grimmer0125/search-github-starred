import {
  GET_REPOS_STATUS_SUCCEEDED,
} from './actionTypes';

function getReposStatusSucceeded(result) {
  return {
    type: GET_REPOS_STATUS_SUCCEEDED,
    payload: {
      result,
    },
  };
}

const actions = {
  getReposStatusSucceeded,
};

export default actions;
