// import api from './apiService';
import fetch from 'isomorphic-fetch';

let baseURL = '';

if (window.location.hostname === 'localhost') {
  baseURL = 'http://localhost:5000';
  // this.socket = io('http://localhost:8000');
}
else if (window.location.hostname === '0.0.0.0') {
  baseURL = 'http://localhost:5000';

//  this.socket = io('http://0.0.0.0:8000');
} else {
  baseURL = 'https://' + window.location.hostname;
}

// return fetch(tokenURL, { credentials: 'include' })

function getReposStatus() {
  const completeURL = baseURL + '/repos';
  console.log('remote url:', completeURL);
  return fetch(completeURL, { credentials: 'include' }).then(res => {
    console.log('get the response');
    // console.log('res:', res.text());
    // debugger;
    // console.log('status:', res.statusText);
    // console.log('res:', res);

    if (res.status === 401) { // statusText === 'Temporary Redirect') {
      const location = '/login';// res.headers.get('location');

      console.log('try to login');

      // if (location) {
      window.location = location;
      // }
    }

    return res.text();
    // console.log('res json :', res.json());
  });
}

export default {
  getReposStatus,
};

// function verifyValidationKey(key) {
//   return api.get(`/key/${key}`);
// }
//
// function submitValidationForm(email, key, password) {
//   return api.post(`/key/${key}`, { email, password });
// }
//
