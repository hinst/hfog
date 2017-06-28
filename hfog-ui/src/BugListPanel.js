import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
	}

	render() {
		const bugHeadItems = this.props.bugs.map(
			(bug) => {
				return (<BugHeadItem key={bug.Number} number={bug.Number} title={bug.Title}></BugHeadItem>);
			}
		);
		return (<div>
			{bugHeadItems}
		</div>);
	}

}

export default BugListPanel;