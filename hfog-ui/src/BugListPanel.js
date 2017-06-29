import React from 'react';
import BugHeadItem from "./BugHeadItem";
import BugSearchPanel from "./BugSearchPanel";

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
      		sortAscending: false,			
			searchPanelVisible: false,
		}
	}

	getSortedBugs() {
		const bugs = this.props.bugs.slice();
		bugs.sort((a, b) => {
			var diff = a.Number - b.Number;
			if (false == this.state.sortAscending) {
				diff = -diff;
			}
			return diff;
		});
		return bugs;
	}

	render() {
		const bugs = this.getSortedBugs();
		const bugHeadItems = bugs.map(
			(bug) => {
				return (<BugHeadItem key={bug.Number} number={bug.Number} title={bug.Title}></BugHeadItem>);
			}
		);
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
			{bugHeadItems}
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

}

export default BugListPanel;