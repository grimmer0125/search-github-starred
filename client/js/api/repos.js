import fetch from 'isomorphic-fetch';
import elasticsearch from 'elasticsearch';

let baseURL = '';

if (window.location.hostname === 'localhost') {
  baseURL = 'http://localhost:5000';
} else if (window.location.hostname === '0.0.0.0') {
  baseURL = 'http://localhost:5000';
} else {
  baseURL = window.location.protocol + '//' + window.location.hostname;
}

function getReposStatus() {
  const completeURL = baseURL + '/repos';
  // console.log('remote url:', completeURL);
  return fetch(completeURL, { credentials: 'include' }).then(res => {

    if (res.status === 401) { // statusText === 'Temporary Redirect') {
      const location = '/login';// res.headers.get('location');

      console.log('try to login');

      window.location = location;
    }

    return res.text();
    // console.log('res json :', res.json());
  });
}

// can not use .com/ , because the limiation of aws policy
const client = new elasticsearch.Client({
  host: 'https://search-searchgithub-7c4xubb6ne3t7keszcai7kqi3m.us-west-2.es.amazonaws.com/githubrepos',
});

const pageSize = 20;

function queryToServer(query, account, page, handler) {
  let finalQuery = query;
  let queryType = 'cross_fields';
  if (query.charAt(0) === '"' && query.charAt(query.length - 1) === '"') {
    queryType = 'phrase';
    finalQuery = query.substr(1, query.length - 2);

    // return a.substr(1, a.length - 2);
  }

  console.log('final query:' + finalQuery + ';type:' + queryType);

  client.search({
    // index: 'githubrepos',
    type: account,
    body: {
      query: {
        multi_match: {
          query: finalQuery, // 'components react interface', // 'react facebook',
          type: queryType, // 'cross_fields',
          fields: ['repofullName', 'description', 'homepage', 'readme'],
          operator: 'and',
        },
      },
      from: page * pageSize,
      size: pageSize,
    },
  }).then(function (resp) {
    handler(resp);
    // const hits = resp.hits.hits;
    // console.log('query result:', hits);
  //  debugger;
  //  const ttt = 0;
  }, function (err) {
    console.log('query to elasticsearch error!!!');
    console.trace(err.message);
  });
}

export default {
  getReposStatus,
  queryToServer,
  pageSize,
};
