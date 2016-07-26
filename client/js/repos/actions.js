import {
  GET_REPOS_STATUS_SUCCEEDED,
} from './actionTypes';

function getReposStatusSucceeded(status) {
  return {
    type: GET_REPOS_STATUS_SUCCEEDED,
    payload: {
      status,
    },
  };
}

const actions = {
  getReposStatusSucceeded,
};

export default actions;
