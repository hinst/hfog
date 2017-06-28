import React, { Component } from 'react';
import './3pty/w3.css';
import BugListPanel from "./BugListPanel";

class App extends Component {

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
        <BugListPanel bugs={this.sampleBugs}></BugListPanel>
      </div>
    );
  }

}

export default App;
