import { fork } from 'redux-saga/effects';

import repos from './repos';

export default function* rootSaga() {
  const forkedReposSagas = repos.sagas.map(saga => fork(saga));

  yield [
    ...forkedReposSagas,
  ];
}
