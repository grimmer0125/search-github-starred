// import { takeEvery } from 'redux-saga';
// import { put, call, select } from 'redux-saga/effects';
//
// import { ActionTypes } from '../actions/deviceAction';
// import { AppManager } from '../utils/AppManager.js';
// import { getRidAliasFromState } from '../reducers/selectors';
// import { fetchDevices } from '../actions/deviceAction.js';
//
// import {
//   tryDropResource,
//   tryDeleteSN,
// } from '../actions/tryExofetch.js';
//
// // export const delay = ms => new Promise(resolve => setTimeout(resolve, ms))
//
// // Our worker Saga:
// export function* deleteDeviceAsync(action) {
//   const sn = action.payload.sn;
//   const ridAliased = yield select(getRidAliasFromState, sn);
//   const rid = ridAliased.rid;
//   const productID = AppManager.instance().getProductID();
//
//   if (rid) {
//     // step 1
//     yield call(tryDropResource, rid);
//     // console.log("drop response:" + dropResponse); // [{"id":1,"status":"ok"}]
//     // console.log('try to handle error case');
//
//     // step 2
//     yield call(tryDeleteSN, productID, sn);
//
//     // http response header :205, reset and empty body means sueccess
//     // and response shoulbe be undefined
//
//     yield put(fetchDevices());
//   }
// }
//
// export function* watchIncrementAsync() {
//   yield* takeEvery(ActionTypes.DELETE_DEVICE, deleteDeviceAsync);
// }
//
// // export function* helloSaga() {
// //   console.log('Hello Sagas!');
// // }
