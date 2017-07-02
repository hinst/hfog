import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListView extends React.PureComponent {

	render() {
		const bugs = this.getSortedBugs();
		const bugHeadItems = bugs.map(
			(bug) => {
				return (
					<BugHeadItem 
						key={bug.Number} 
						number={bug.Number} 
						title={bug.Title}
						clickReceiver={this.props.itemClickReceiver}
					/>
				);
			}
		);
		return (<div>{bugHeadItems}</div>);
	}

	getSortedBugs() {
		const sortAscending = this.props.sortAscending;
		const bugs = this.props.bugs.slice();
		bugs.sort((a, b) => {
			var rankDiff = a.Rank - b.Rank;
			if (rankDiff === 0) {
				var numberDiff = a.Number - b.Number;
				if ( !sortAscending ) {
					numberDiff = -numberDiff;
				}
				return numberDiff;
			} else {
				return -rankDiff;
			}
		});
		return bugs;
	}

}

export default BugListView;