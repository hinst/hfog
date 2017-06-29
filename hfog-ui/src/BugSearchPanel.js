import React from 'react';
import AppHeader from "./AppHeader.js";

class BugSearchPanel extends React.Component {

	render() {
		return (<div className="w3-container">
			{AppHeader()}
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={this.props.backClickReceiver}> ← Search</button>
				<div style={{height: "8px"}}></div>
				<label><b>Keywords:</b></label>
				<br/>
				<input className="w3-input w3-border" type="text"/>
				<div style={{height: "8px"}}></div>				
				<button className="w3-btn w3-black">Go</button>
			</div>
		</div>)
	}

}

export default BugSearchPanel;