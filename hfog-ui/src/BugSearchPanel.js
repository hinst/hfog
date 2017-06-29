import React from 'react';
import AppHeader from "./AppHeader.js";

class BugSearchPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			keywords: "",
		};
	}

	render() {
		return (<div className="w3-container">
			{AppHeader()}
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={this.props.backClickReceiver}> ← Search</button>
				<div style={{height: "8px"}}></div>
				<label><b>Keywords:</b></label>
				<br/>
				<input 
					type="text"
					ref={(thing) => { this.keywordsField = thing; }} 
					className="w3-input w3-border" 
					onChange={
						(event) => {
							this.setState(
								{keywords: event.target.value}
							);
						}
					}
					onKeyPress={
						(event) => {
							if (event.key === 'Enter') {
								this.receiveGoClick();
							}
						}
					}
				/>
				<div style={{height: "8px"}}></div>				
				<button className="w3-btn w3-black" onClick={() => this.receiveGoClick()}>Go</button>
			</div>
		</div>)
	}

	receiveGoClick() {
		this.props.searchActReceiver(this.state.keywords);
	}

	componentDidMount() {
		this.keywordsField.focus();
		console.log("?");
	}

}

export default BugSearchPanel;