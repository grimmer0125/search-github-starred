//
// import * as ActionTypes from '../actions/deviceAction';
//
// const unknownSN = '';
// const unknownRID = '';
// const unknownName = '';
//
// // will change later, e.g. track.js?
//
// export function okami(state = { productId: '', productName: '', endpoint: '' }, action) {
//   switch (action.type) {
//     case ActionTypes.GETPRODUCTINFO:
//       if (action.error) {
//         return state;
//       }
//       return {
//         ...state,
//         productName: action.payload.info.label,
//         endpoint: action.payload.info.endpoint,
//       };
//     case ActionTypes.getProductID:
//       return { ...state, productId: action.payload.productId };
//     default:
//       return state;
//   }
// }
//
//
// export function selectedDeviceToken(state = '', action) {
//   switch (action.type) {
//     case ActionTypes.GETAUTHTOKEN:
//       return action.payload.token;
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return '';
//     default:
//       return state;
//   }
// }
//
// export function addModelResouceName(state = '', action) {
//   switch (action.type) {
//     case ActionTypes.changeModelResourceName:
//       return action.payload.name;
//     default:
//       return state;
//   }
// }
//
// export function addModelResourceType(state = 0, action) {
//   switch (action.type) {
//     case ActionTypes.selecteModelResourceType:
//       return action.payload.resourceType;
//     default:
//       return state;
//   }
// }
//
// export function addModelResourceFormat(state = 0, action) {
//   switch (action.type) {
//     case ActionTypes.selecteModelResourceFormat:
//       return action.payload.resourceFormat;
//     default:
//       return state;
//   }
// }
//
// export function addModelResourceDialogOpen(state = false, action) {
//   switch (action.type) {
//     case ActionTypes.addModelDialogResourceSwitch:
//
//       return action.payload.status;
//     default:
//       return state;
//   }
// }
//
// const defaultLimitArg = {
//   limitSms: 0,
//   limitEmail: 100,
//   limitHttp: 5000,
// };
//
// export function definitionModel(state = { ...defaultLimitArg }, action) {
//   switch (action.type) {
//     case ActionTypes.GETMODELLIMITS:
//       return { limitSms: action.payload.limits.sms,
//             limitEmail: action.payload.limits.email,
//             limitHttp: action.payload.limits.http };
//     default:
//       return state;
//   }
// }
//
// // ~ selectedDevice
// export function modelRID(state = '', action) {
//   switch (action.type) {
//     case ActionTypes.GETMODEL_RID:
//       return action.payload.rid;
//     default:
//       return state;
//   }
// }
//
// export function modelResources(state = { items: [], isFetching: false }, action) {
//   switch (action.type) {
//     case ActionTypes.getModelResources:
//       return {
//         isFetching: false,
//         items: action.payload.resources,
//       };
//     default:
//       return state;
//   }
// }
//
// export const selectedDevice = (state = { sn: unknownSN, rid: unknownRID,
//   name: unknownName }, action) => {
//   switch (action.type) {
//     case ActionTypes.NAVI_DETAILPAGE:
//       return {
//         sn: action.payload.sn,
//         rid: action.payload.rid,
//         name: state.name,
//         status: state.status,
//       };
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return { sn: unknownSN, rid: unknownRID, name: unknownName };
//     default:
//       return state;
//   }
// };
//
// // { sn: '123', status: 'activated' },
// // { sn: '456', status: 'nonactivated' },
// const defaultDevices = [];
//
// const defaultResources = [
// ];
//
//
// export function resourceData(state = {}, action) {
//   switch (action.type) {
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return {};
//     case ActionTypes.getModelResourceData:
//     case ActionTypes.getResourceData: {
//       if (action.error) {
//         return state;
//       }
//       const resourceIDData = {};
//       resourceIDData.lastValue = action.payload.value;
//       const copySource = {};
//       copySource[action.payload.rid] = resourceIDData;
//       return {
//         ...state, ...copySource,
//       };
//     }
//     default:
//       return state;
//   }
// }
//
// export function resourceInfos(state = {}, action) {
//   switch (action.type) {
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return {};
//     case ActionTypes.getModelResourceInfo:
//     case ActionTypes.getResourceInfo: {
//       if (action.error) {
//         return state;
//       }
//       const resourceIDInfo = {};
//       resourceIDInfo.info = action.payload.info;
//       resourceIDInfo.format = action.payload.format;
//       resourceIDInfo.subscription = action.payload.subscription;
//       const copySource = {};
//       copySource[action.payload.rid] = resourceIDInfo;
//
//
//       return {
//         ...state, ...copySource,
//       };
//     }
//     default:
//       return state;
//   }
// }
//
// export function resources(
//   state = { isFetching: false, items: defaultResources },
//   action
// ) {
//   switch (action.type) {
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return {
//         isFetching: false,
//         items: defaultResources,
//       };
//     case ActionTypes.getDeviceResources:
//       return {
//         isFetching: false,
//         items: action.payload.resources,
//       };
//     default:
//       return state;
//   }
// }
//
//
// const defaultResourcesArg = {
//   /* Dataport Page*/
//   dataportAlias: '',
//   dataportFormat: '',
//   dataportFormatIndex: 0,
//   dataportLimitsType: '',
//   dataportMeta: '',
//   dataportName: '',
//   dataportPreprocess: [],
//   dataportPublicState: false,
//   dataportShares: [],
//   dataportSubscribe: '',
//   dataportValueHistory: [],
//
//   /* Datarule Page*/
//   dataRuleAlias: '',
//   dataRuleFormat: '',
//   dataRuleMeta: '',
//   dataRulePreprocesses: [],
//   dataRulePublicState: false,
//   dataRuleRules: '',
//   dataRuleSubscribed: '',
//
//   /* Dispatch Page*/
//   dispatchFormatIndex: 0,
//   dispatchLockedState: false,
//   dispatchMessage: '',
//   dispatchName: '',
//   dispatchPreprocess: [],
//   dispatchPublicState: false,
//   dispatchShares: [],
//   dispatchSubject: '',
//   dispatchSubscribe: '',
//   dispatchValueHistory: [],
//
//   /* Script Page */
//   luaScript: '',
//   scriptLog: [],
//   scriptName: '',
//   scriptPublicState: false,
// };
//
// export function selectedResource(
//     state = { isFetching: false, type: 'client', rid: '', ...defaultResourcesArg },
//     action
// ) {
//   switch (action.type) {
//     case ActionTypes.SHOW_RESOURCEDETAIL:
//       return {
//         isFetching: false,
//         rid: action.payload.rid,
//         type: action.payload.type,
//
//       };
//     case ActionTypes.LEAVEDEVICEDETAIL:
//       return { isFetching: false, type: 'client', rid: '', ...defaultResourcesArg };
//     default:
//       return state;
//   }
// }
//
//
// function updateSomeDeviceStatus(state, rid, status) {
//   const newItems = [];
//   const foundOut = false;
//   for (const item of state.items) {
//     const newItem = { ...item };
//     if (!foundOut && item.rid === rid) {
//       newItem.status = status;
//     }
//     newItems.push(newItem);
//   }
//
//   return newItems;
// }
//
// function updateSomeDeviceNameAliases(state, rid, name, aliases) {
//   const newItems = [];
//   const foundOut = false;
//   for (const item of state.items) {
//     const newItem = { ...item };
//     if (!foundOut && item.rid === rid) {
//       newItem.name = name;
//       newItem.aliases = aliases;
//     }
//     newItems.push(newItem);
//   }
//
//   return newItems;
// }
//
// export function devices(
//   state = { isFetching: false, items: defaultDevices },
//   action
// ) {
//   switch (action.type) {
//     case ActionTypes.getDevice:
//       return {
//         isFetching: false,
//         items: action.payload.devices,
//       };
//     case ActionTypes.getDeviceStatus:
//       return {
//         isFetching: false,
//         items: updateSomeDeviceStatus(state, action.payload.rid, action.payload.status),
//       };
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
