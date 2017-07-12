import React from 'react';
import * as AppHeader from "./AppHeader.js";

class BugSearchPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			keywords: this.props.keywords,
			searchBodyEnabled: this.props.searchBodyEnabled,
		};
	}

	render() {
		return (<div className="w3-container">
			{AppHeader.AppHeaderPanel()}
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={this.props.backClickReceiver}> ← Search</button>
				<div style={{height: "8px"}}></div>
				<label><b>Keywords:</b></label>
				<br/>
				<input
					type="text"
					autoFocus
					ref={(thing) => { this.keywordsField = thing; }} 
					className="w3-input w3-border" 
					onChange={
						(event) => {
							this.setState(
								{keywords: event.target.value}
							);
						}
					}
					value={this.state.keywords}
					onKeyPress={
						(event) => {
							if (event.key === 'Enter') {
								this.receiveGoClick();
							}
						}
					}
				/>
				<input type="checkbox"
					className="w3-check"
					onChange={
						(event) => {
							this.setState({searchBodyEnabled: event.target.value});
						}
					}
				/>
				<label>Search body</label>
				<div style={{height: "8px"}}></div>				
				<button className="w3-btn w3-black" onClick={() => this.receiveGoClick()}>Go</button>
			</div>
		</div>)
	}

	receiveGoClick() {
		this.props.searchActReceiver(this.state.keywords);
	}

}

export default BugSearchPanel;