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
  console.log('remote url:', baseURL + '/repos');
  return fetch(baseURL + '/repos').then(res => {
    console.log('get the response');
    // console.log('res:', res.text());
    // debugger;
    return res.text();
    // console.log('res json :', res.json());
  });
}

export default {
  getReposStatus,
};
