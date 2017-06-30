import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListView extends React.PureComponent {

	render() {
		const bugs = this.getSortedBugs();
		const bugHeadItems = bugs.map(
			(bug) => {
				return (<BugHeadItem key={bug.Number} number={bug.Number} title={bug.Title}></BugHeadItem>);
			}
		);
		return (<div>{bugHeadItems}</div>);
	}

	getSortedBugs() {
		const bugs = this.props.bugs.slice();
		bugs.sort((a, b) => {
			var diff = a.Number - b.Number;
			if ( ! this.props.sortAscending) {
				diff = -diff;
			}
			return diff;
		});
		return bugs;
	}

}

export default BugListView;