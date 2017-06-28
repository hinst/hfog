import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
	}

	render() {
		const bugHeadItems = this.props.bugs.map(
			(bug) => {
				return (<BugHeadItem key={bug.number} number={bug.number} title={bug.title}></BugHeadItem>)
			}
		);
		return (<div>
			{bugHeadItems}
		</div>);
	}

}

export default BugListPanel;