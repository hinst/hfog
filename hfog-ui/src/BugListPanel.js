import React from 'react';
import * as Api from "./Api";
import BugHeadItem from "./BugHeadItem";
import BugSearchPanel from "./BugSearchPanel";
import BugListView from "./BugListView";

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			bugs: [],
      		sortAscending: false,			
			searchPanelVisible: false,
		}
	    this.requestBugs();
	}

	render() {
		return (<div>
			<div className="overlay" style={{ zIndex: 1, display: this.getSearchPanelDisplay() }}>
				<div className="overlay-content">
					<BugSearchPanel backClickReceiver={ ()=>this.receiveSearchPanelClickBack() }>
					</BugSearchPanel>
				</div>
			</div>
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>{this.state.sortAscending ? "▲" : "▼"}</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black">Refresh</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={ () => this.receiveSearchClick() }>Search</button>
			</div>
			<BugListView bugs={this.state.bugs} sortAscending={this.state.sortAscending}></BugListView>
		</div>);
	}

	getSearchPanelDisplay() {
		return this.state.searchPanelVisible ? null : "none";
	}

	changeSortDirection() {
		this.setState({sortAscending: ! this.state.sortAscending});
	}

	receiveSearchClick() {
		this.setState({searchPanelVisible: ! this.state.searchPanelVisible});
	}

	receiveSearchPanelClickBack() {
		this.setState({searchPanelVisible: false});
	}

  requestBugs() {
    Api.LoadBugList(data => this.receiveBugs(data));
  }

  receiveBugs(data) {
    console.log(data.length);
    this.setState({bugs: data});
  }

}

export default BugListPanel;