import React from 'react';
import FetchingStatus from './constants';
import { connect } from 'react-redux';
import algoliasearch from 'algoliasearch';
import { reduxForm } from 'redux-form';

// import { bindActionCreators } from 'redux';

import {
  FETCH_STARRRED_STATUS,
} from './actionTypes';

class RepoList extends React.Component {

  render() {
    const createItem = function (item) {
      return (
        <li key={item.id}>
          <a href={item.url}>{item.text}</a>
        </li>
      );
    };
    return <ul>{this.props.items.map(createItem)}</ul>;
  }
}

class ReposPage extends React.Component {

  constructor(props) {
    super(props);

    this.state = { items: [], text: '' };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.onChange = this.onChange.bind(this);
    // props.handleSubmit = this.handleSubmit;
  }
  onChange(e) {
    this.setState({ text: e.target.value });
  }

  handleSubmit(e) {
    console.log('handleSubmit !!! ');
    e.preventDefault();

    if (this.state.text !== '') {
      if (this.props.repos.githubAccount) {
        console.log('Start to query !!!!!!');
        this.queryToServer(this.state.text, this.props.repos.githubAccount);
      }
    }
  }

  /* <h3>Starred Repo</h3>*/

  renderReposComponents() {
    return (
          <div>
            <form onSubmit={this.handleSubmit}>
              <input onChange={this.onChange} value={this.state.text} />
              <button>Search</button>
            </form>
            <RepoList items={this.state.items} />
          </div>
        );
  }


//  expect(handleSubmit(submit, values, props, asyncValidate)).toBe(undefined);

  // handleSubmit() {
  //   console.log('handleSubmit');
  // }

  componentWillMount() {
    // const { businessId, getBillingData } = this.props;
    if (!this.hasData()) {
      this.startPoll();
    }
  }

  componentDidMount() {
    // const fetchAction = bindActionCreators(fetchDevices, this.props.dispatch);
    // fetchAction();
  }


  componentWillReceiveProps(nextProps) {
    console.log('componentWillReceiveProps');
    if (this.props.repos !== nextProps.repos) {
      console.log('different Props');


      // Optionally do something with data

      // if (!nextProps.isFetching) {
      //   this.startPoll();
      // }
    }
    clearTimeout(this.timeout);

    if (!this.hasData()) {
      console.log('start timer in receive props ');
      this.startPoll();
    }
    // else {
    //   this.queryToServer('react', this.props.repos.githubAccount);
    // }
  }

  // filters: "facet1 AND facet2"
  queryToServer(query, account) {
    const appID = 'EQDRH6QSH7';
    const key = '6066c3e492d3a35cc0a425175afa89ff';
    const indexName = 'githubRepo';
    const attributesToSnippet = ['readmd:5', 'description:5', 'homepage:5', 'repoURL:5'];
    const facet = 'starredBy:' + account;
    const facetFilters = [facet];
    const client = algoliasearch(appID, key);
    const index = client.initIndex(indexName);
    index.search(query, { attributesToSnippet, facetFilters }, (err, content) => {
      console.log('error:', err);
      console.log('content:', content);

      const hits = content.nbHits;

      const currentPage = content.page;
      const totalPage = content.nbPages;

      // ownerURL
      // :
      // "https://github.com/reactjs"

      const hitsList = content.hits;
      const nextItems = [];
      const checkDict = {};
      for (const hit of hitsList) {
        if (checkDict.hasOwnProperty(hit.repoURL) === false) {
          const item = { text: hit.repoURL, url: hit.repoURL, id: hit.repoURL };
          checkDict[hit.repoURL] = 1;
          nextItems.push(item);
        }
      }

      // const nextItems = this.state.items.concat([{ text: this.state.text, id: Date.now() }]);
      const nextText = '';
      this.setState({ items: nextItems, text: nextText });
    });
  }

  // const type = 'DELETE_DEVICE';
  // this.props.dispatch({ type, payload: { sn } });
  // dispatch({type: 'USER_FETCH_REQUESTED', payload: {userId}})

  // ref: http://notjoshmiller.com/ajax-polling-in-react-with-redux/
//  this.timeout = setTimeout(() => this.props.dataActions.dataFetch(), 15000);
  startPoll() {
    const { dispatch } = this.props;
    // console.log('status timer runs !!!');

    // dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });

    this.timeout = setTimeout(() => {
      console.log('timer runs !!!');
      // debugger;
      dispatch({ type: FETCH_STARRRED_STATUS, payload: { text: 'Do something.' } });
    }, 2000);
  }

  hasData() {
    // const { starredRepos } = this.props;
    const starredRepos = this.props.repos;
    return (!starredRepos.error && starredRepos.fetchingStatus === FetchingStatus.INDEXED);
  }

  renderComponents() {
    // redux-form part
    // let { fields: { firstName, lastName, email }, handleSubmit } = this.props;
    // let { fields: { firstName, lastName, email }, handleSubmit } = this.props;
    // handleSubmit = this.handleSubmit;
    // const reduxPart = (
    //   <form onSubmit={handleSubmit}>
    //     <div>
    //       <label>First Name</label>
    //       <input type="text" placeholder="First Name" {...firstName} />
    //     </div>
    //     <div>
    //       <label>Last Name</label>
    //       <input type="text" placeholder="Last Name" {...lastName} />
    //     </div>
    //     <div>
    //       <label>Email</label>
    //       <input type="email" placeholder="Email" {...email} />
    //     </div>
    //     <button type="submit">Submit</button>
    //   </form>
    // );


    const { numOfStarred } = this.props.repos;

    let statusStr = '';

    if (numOfStarred > 0) {
      statusStr = 'Indexing is finished. Total repos: ' + numOfStarred + '. Start to query.';
    } else {
      statusStr = 'Indexing is finished, start to query';
    }

    const reposComponent = this.renderReposComponents();

    return (
      <div className="flex-column layout-column-start-center" style={{ width: '100%' }}>
        {statusStr}
        <div>
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
        statusStr = 'Fetching is not started yet';
        break;
      case FetchingStatus.FETCHING:
        statusStr = 'It is fetching, wait a moement...';
        break;
      case FetchingStatus.INDEXING:
        if (numOfStarred > 0) {
          statusStr = 'It is indexing ' + numOfStarred + ', wait a mement...';
        } else {
          statusStr = 'It is indexing, wait a mement...';
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
      console.log('has data');
      return this.renderComponents();
    }

    console.log('has no data');

    return this.renderLoadingScreen();
  }
}

ReposPage.propTypes = {
  repos: React.PropTypes.object,
  dispatch: React.PropTypes.func,
};

function test(state) {
  console.log('test test');
  return state.repos;
}

export function mapStateToProps(state) {
  return {
    repos: test(state),
  };
}

// export function mapDispatchToProps(dispatch) {
//   return bindActionCreators({
//     xxx: actions.xxx

//   }, dispatch);
// }


//

// ReposPage = reduxForm({ // <----- THIS IS THE IMPORTANT PART!
//   form: 'contact',                           // a unique name for this form
//   fields: ['firstName', 'lastName', 'email'], // all the fields in your form
// })(ReposPage);

export default connect(mapStateToProps)(ReposPage);


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
