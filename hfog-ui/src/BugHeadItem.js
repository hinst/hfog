import React from 'react';

class BugHeadItem extends React.PureComponent {

	render() {
		return (<div className="w3-panel">
			<button 
				className="w3-btn w3-black" 
				onClick={()=>this.receiveClick}
			>
				{this.props.number}
			</button>
			{this.props.title}
		</div>)
	}

	receiveClick() {
		this.props.clickReceiver(this.props.number);
	}

}

export default BugHeadItem;