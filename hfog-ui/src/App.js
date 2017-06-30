import React, { Component } from 'react';
import './3pty/w3.css';
import BugListPanel from "./BugListPanel";
import "./App.css";
import AppHeader from "./AppHeader.js";
import { BrowserRouter as Router, Route, IndexRoute } from 'react-router-dom';
import BugView from './BugView';

const BugList = () => (
  <div>
    {AppHeader()}
    <BugListPanel></BugListPanel>
  </div>
);

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
    };
  }

  sampleBugs = [
    {number: 1, title: "Bug number one"},
    {number: 2, title: "Bug number two"},
  ];

  render() {
    return (
      <div className="w3-container">
        <Router>
          <div>
            <Route exact path="/" component={BugList}/>
            <Route path="/bug" component={BugView}/>
          </div>
        </Router>
      </div>
    );
  }

}

export default App;
