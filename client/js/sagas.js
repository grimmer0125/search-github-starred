// import { fork } from 'redux-saga/effects';
//
// import auth from './auth';
// import billing from './billing';
// import business from './business';
// import products from './products';
// import profile from './profile';
// import solutions from './solutions';
// import user from './user';
// import verification from './verification';
//
// export default function* rootSaga() {
//   const forkedAddProductSagas = products.addProduct.sagas.map(saga => fork(saga));
//   const forkedAddSolutionSagas = solutions.addSolution.sagas.map(saga => fork(saga));
//   const forkedBillingSagas = billing.sagas.map(saga => fork(saga));
//   const forkedBusinessSagas = business.sagas.map(saga => fork(saga));
//   const forkedLoginSagas = auth.login.sagas.map(saga => fork(saga));
//   const forkedMakePaymentSagas = billing.makePayment.sagas.map(saga => fork(saga));
//   const forkedProductSagas = products.sagas.map(saga => fork(saga));
//   const forkedProfileSagas = profile.sagas.map(saga => fork(saga));
//   const forkedSolutionSagas = solutions.sagas.map(saga => fork(saga));
//   const forkedUserSagas = user.sagas.map(saga => fork(saga));
//   const forkedVerificationSagas = verification.sagas.map(saga => fork(saga));
//
//   yield [
//     ...forkedAddProductSagas,
//     ...forkedAddSolutionSagas,
//     ...forkedBillingSagas,
//     ...forkedBusinessSagas,
//     ...forkedLoginSagas,
//     ...forkedMakePaymentSagas,
//     ...forkedProductSagas,
//     ...forkedProfileSagas,
//     ...forkedSolutionSagas,
//     ...forkedUserSagas,
//     ...forkedVerificationSagas,
//   ];
// }
