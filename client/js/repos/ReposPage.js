import React from 'react';
import FetchingStatus from './constants';
import { connect } from 'react-redux';
// import algoliasearch from 'algoliasearch';
import api from '../api';

import {
  FETCH_STARRRED_STATUS,
} from './actionTypes';

// testing
// api.queryToServer('react', 'grimmer0125', 0, (resp) => {
//   console.log('get the resp:', resp);
// });

class RepoList extends React.Component {

  render() {
    const createItem = function (item) {
      return (
        <li key={item.id}>
          <a href={item.url}>{item.repofullName}</a> {item.desc}
        </li>
      );
    };
    return <ul>{this.props.items.map(createItem)}</ul>;
  }
}

const QueryStatus = {
  NOTQUERY: 'NotQuery',
  QUERYING: 'Querying',
  QUERIED: 'Queried',
};

class ReposPage extends React.Component {

  resetQueryParameters() {
    this.state.currentPage = -1;
  }

  constructor(props) {
    super(props);

    this.state = {
      queryStats: QueryStatus.NOTQUERY,
      hits: [],
      total: 0,
      totalPage: 0,
      currentPage: -1,
      textOnQueryInput: '',
      queryCursor: '',
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onChange = this.onChange.bind(this);
    this.handleNext = this.handleNext.bind(this);
    this.handlePrev = this.handlePrev.bind(this);
    this.handleReIndex = this.handleReIndex.bind(this);
    this.handleQueryData = this.handleQueryData.bind(this);
    // props.handleSubmit = this.handleSubmit;
  }
  onChange(e) {
    this.setState({ textOnQueryInput: e.target.value });
  }

  // TODO: should add handling error case, e.g. currentPage
  // 1. error就回到第一頁或是query前一頁.
  // 2. handlePrev 直接改成用local資料, 這樣就還是可以把page++的logic放在得到資料時, 不然
  handleQueryData(resp) {
    this.state.queryStats = QueryStatus.QUERIED;

    // elasticsearch type
    // console.log('query result:', resp);
    const hitsList = resp.hits.hits;
    this.state.total = resp.hits.total;

    this.state.totalPage = Math.ceil(this.state.total / api.pageSize);
    // resp.hits.total /api.pageSize

    // algolia type
    // console.log('content:', content);
    // const hitsList = content.hits;
    // this.state.total = content.nbHits;
    // this.state.currentPage = content.page;
    // this.state.totalPage = content.nbPages;
    // this.state.queryCursor = content.query;

    this.state.queryStats = QueryStatus.QUERIED;

    const nextItems = [];
    const checkDict = {};
    for (const hitData of hitsList) {
      const hit = hitData._source; // elasticsearch type, algo: hit = hitData
      if (checkDict.hasOwnProperty(hit.repoURL) === false) {
        const item = { url: hit.repoURL,
          id: hit.repoURL, desc: hit.description, repofullName: hit.repofullName };
        checkDict[hit.repoURL] = 1;
        nextItems.push(item);
      }
    }

    this.setState({ hits: nextItems });
  }

  handleNext(e) {
    console.log('click next');

    api.queryToServer(this.state.queryCursor, this.props.repos.githubAccount,
    this.state.currentPage + 1, this.handleQueryData);

    this.state.currentPage++;
  }

  handlePrev(e) {
    console.log('click prev');

    api.queryToServer(this.state.queryCursor, this.props.repos.githubAccount,
    this.state.currentPage - 1, this.handleQueryData);

    this.state.currentPage--;
  }

  handleReIndex() {
    console.log('re-index, redirect to login page');
    const location = '/repos/_reindex';
    window.location = location;
  }

  handleSubmit(e) {
    e.preventDefault();

    if (this.state.textOnQueryInput !== '') {
      if (this.props.repos.githubAccount) {
        this.state.queryStats = QueryStatus.QUERYING;

        api.queryToServer(this.state.textOnQueryInput,
          this.props.repos.githubAccount, 0, this.handleQueryData);

        this.state.queryCursor = this.state.textOnQueryInput;
        this.resetQueryParameters();
        this.state.currentPage++;
      }
    }
  }

  renderReposComponents() {
    // const { repos } = this.props;
    const showPre = this.state.currentPage > 0 ?
    (<button onClick={this.handlePrev}>Prev</button>) : null;
    const showNext = this.state.currentPage < (this.state.totalPage - 1) ?
    (<button onClick={this.handleNext}>Next</button>) : null;

    const nextOperation = this.state.queryStats !== QueryStatus.QUERIED ? null : (
      <div className="flex-row">
        <span>
          Hits:{this.state.total}. Pages:{this.state.totalPage}. CurrentPage: {this.state.currentPage + 1}.
        </span>
        {showPre}
        {showNext}
      </div>
    );
    return (
          <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
            <form onSubmit={this.handleSubmit}>
              <input onChange={this.onChange} value={this.state.textOnQueryInput} />
              <button>Search</button>
            </form>
            {nextOperation}
            <RepoList items={this.state.hits} />
          </div>
    );
  }

  componentWillMount() {
    if (!this.hasData()) {
      this.startPoll();
    }
  }

  componentDidMount() {

  }

  componentWillReceiveProps(nextProps) {
    if (this.props.repos !== nextProps.repos) {

      // Optionally do something with data
      // if (!nextProps.isFetching) {
      //   this.startPoll();
      // }
    }
    clearTimeout(this.timeout);

    if (!this.hasData()) {
      // console.log('Polling indexing status');
      this.startPoll();
    }
    // else {
    //   this.queryToServer('react', this.props.repos.githubAccount);
    // }
  }

  // hitsPerPage
  // page:
  // filters: "facet1 AND facet2"
  // algoliasearch part
  // queryToServer(query, account, page = 0) {
  //   const appID = 'EQDRH6QSH7';
  //   const key = '6066c3e492d3a35cc0a425175afa89ff';
  //   const indexName = 'githubRepo';
  //   const attributesToSnippet = ['readmd:5', 'description:5', 'homepage:5', 'repoURL:5'];
  //   // const facet = 'starredBy:' + account;
  //   // const facetFilters = [facet];
  //   const filters = 'starredBy:' + account;
  //
  //   const client = algoliasearch(appID, key);
  //   const index = client.initIndex(indexName);
  //   const typoTolerance = false;
  //
  //   index.search(query, { attributesToSnippet, filters, page, typoTolerance }, (err, content) => {
  //     this.state.queryStats = QueryStatus.QUERIED;
  //
  //     if (err) {
  //       console.log('error:', err);
  //     }
  //     console.log('content:', content);
  //
  //     this.state.total = content.nbHits;
  //     this.state.currentPage = content.page;
  //     this.state.totalPage = content.nbPages;
  //     this.state.queryStats = QueryStatus.QUERIED;
  //
  //     const hitsList = content.hits;
  //     const nextItems = [];
  //     const checkDict = {};
  //     for (const hit of hitsList) {
  //       if (checkDict.hasOwnProperty(hit.repoURL) === false) {
  //         const item = { url: hit.repoURL,
  //           id: hit.repoURL, desc: hit.description, repofullName: hit.repofullName };
  //         checkDict[hit.repoURL] = 1;
  //         nextItems.push(item);
  //       }
  //     }
  //
  //     this.setState({ hits: nextItems });
  //   });
  // }

  startPoll() {
    const { dispatch } = this.props;

    // console.log('status timer runs !!!');
    // dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });

    this.timeout = setTimeout(() => {
      dispatch({ type: FETCH_STARRRED_STATUS }); // , payload: { text: 'Do something.' }
    }, 2000);
  }

  hasData() {
    const repos = this.props.repos;
    const afterIndexStatus = (repos.fetchingStatus === FetchingStatus.INDEXED);

    return (!repos.error && afterIndexStatus);
  }

  renderComponents() {
    const { numOfStarred } = this.props.repos;

    let statusStr = '';

    // if (numOfStarred > 0) {
    statusStr = `Total: ${numOfStarred}. Start to full-text search your starred repos.`;
    // } else {
    //   statusStr = 'Indexing is finished, start to query';
    // }

    const reposComponent = this.renderReposComponents();

    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        <div>
          <a href="https://grimmer.io/">About me</a>
        </div>
        <div className="flex-column">
          <span>{statusStr}</span>
          <button onClick={this.handleReIndex}>Re-Index</button>
        </div>
        <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
          {reposComponent}
        </div>
      </div>
    );
  }

  renderLoadingScreen() {
    // use a common loading screen
    // className="page"
    let statusStr = '';
    // console.log('status in render;', this.props);
    const { fetchingStatus, numOfStarred } = this.props.repos;

    switch (fetchingStatus) {
      case FetchingStatus.NOTSTART:
        statusStr = 'Loading...';// 'Fetching is not started yet';
        break;
      case FetchingStatus.FETCHING:
        statusStr = 'It is fetching, wait a moment...';
        break;
      case FetchingStatus.INDEXING:
        if (numOfStarred > 0) {
          statusStr = `It is indexing ${numOfStarred} repos, wait a moment...`;
        } else {
          statusStr = 'It is indexing, wait a moment...';
        }
        break;
      default:
        statusStr = 'unknown status';
    }

    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        <div className="loading-text">
          {statusStr}
        </div>
      </div>
    );
  }

  render() {
    if (this.hasData()) {
      return this.renderComponents();
    }

    return this.renderLoadingScreen();
  }
}

ReposPage.propTypes = {
  repos: React.PropTypes.object,
  dispatch: React.PropTypes.func,
};

export function mapStateToProps(state) {
  return {
    repos: state.repos,
  };
}

export default connect(mapStateToProps)(ReposPage);

// export function mapDispatchToProps(dispatch) {
//   return bindActionCreators({
//     xxx: actions.xxx
//   }, dispatch);
// }

// ReposPage = reduxForm({ // <----- THIS IS THE IMPORTANT PART!
//   form: 'contact',                           // a unique name for this form
//   fields: ['firstName', 'lastName', 'email'], // all the fields in your form
// })(ReposPage);

//
// // index.search("react", {
// //   "getRankingInfo": 1,
// //   "facets": "*",
// //   "attributesToRetrieve": "*",
// //   "highlightPreTag": "<em>",
// //   "highlightPostTag": "</em>",
// //   "hitsPerPage": 10,
// //   "facetFilters": [
// //     "starredBy:grimmer0125"
// //   ],
// //   "maxValuesPerFacet": 100
// // });
//
// index.search('query string', {
//   attributesToRetrieve: ['firstname', 'lastname'],
//   hitsPerPage: 50,
// }, function searchDone(err, content) {
//   if (err) {
//     console.error(err);
//     return;
//   }
//
//   for (const h in content.hits) {
//     console.log('Hit(' + content.hits[h].objectID + '): ' + content.hits[h].toString());
//   }
// });
