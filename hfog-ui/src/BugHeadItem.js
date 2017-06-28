import React from 'react';

class BugHeadItem extends React.Component {

	constructor(props) {
		super(props);
	}

	render() {
		return (<div className="w3-panel">
			<button className="w3-btn w3-black">{this.props.number}</button> {this.props.title}
		</div>)
	}

}

export default BugHeadItem;