import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
      		sortAscending: false,			
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
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>{this.state.sortAscending ? "▲" : "▼"}</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black">Refresh</button>
			</div>
			{bugHeadItems}
		</div>);
	}

	changeSortDirection() {
		this.setState({sortAscending: !this.state.sortAscending});
	}

}

export default BugListPanel;