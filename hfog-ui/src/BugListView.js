import React from 'react';
import BugHeadItem from "./BugHeadItem";

class BugListView extends React.PureComponent {

	render() {
		const bugs = this.props.bugs;
		let content = "";
		if (bugs != null) {
			if (bugs.length > 0) {
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
				content = bugHeadItems;
			} else {
				content = <div className="w3-panel">Empty list</div>;
			}
		}
		return (<div>{content}</div>);
	}

}

export default BugListView;