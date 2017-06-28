import React, { Component } from 'react';
import './3pty/w3.css';
import BugListPanel from "./BugListPanel";
import * as Api from "./Api";

class App extends Component {

  constructor(props) {
    super(props);
    this.state = {
      bugs: [],
    };
    this.requestBugs();
  }

  sampleBugs = [
    {number: 1, title: "Bug number one"},
    {number: 2, title: "Bug number two"},
  ];

  render() {
    return (
      <div className="w3-container">
        <div className="w3-container">
          <h1>FogBugz backup</h1>
        </div>
        <BugListPanel bugs={this.state.bugs}></BugListPanel>
      </div>
    );
  }

  requestBugs() {
    Api.LoadBugList(data => this.receiveBugs(data));
  }

  receiveBugs(data) {
    console.log(data.length);
    this.setState({bugs: data});
  }

}

export default App;
