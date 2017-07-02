import React from 'react';
import * as Api from "./Api";
import BugSearchPanel from "./BugSearchPanel";
import BugListView from "./BugListView";
import AppURL from './AppURL';

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			bugs: [],
      		sortAscending: false,
			searchPanelVisible: false,
			filterString: "",
		}
	  this.requestBugs();
	}

	render() {
		return (<div>
			{this.state.searchPanelVisible
				?(
					<div className="overlay" style={{ zIndex: 1, display: this.getSearchPanelDisplay() }}>
						<div className="overlay-content">
							<BugSearchPanel 
								keywords={this.state.filterString}
								backClickReceiver={ () => this.receiveSearchPanelClickBack() }
								searchActReceiver={ (keywords) => this.receiveSearchAct(keywords) }
							>
							</BugSearchPanel>
						</div>
					</div>
				)
				: ""
			}
			<div className="w3-panel">
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>{this.state.sortAscending ? "▲" : "▼"}</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={() => this.receiveRefreshClick()}>Refresh</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={ () => this.receiveSearchClick() }>Search</button>
				<span style={{marginLeft: "4px"}}></span>
				Showing {this.state.bugs.length} items
				<span style={{marginLeft: "4px"}}></span>
				{this.state.filterString.length > 0
					? (
						<button 
							className="w3-btn w3-black" 
							onClick={() => {
								this.setState({filterString: "", bugs: []}, () => this.requestBugs());
							}}
						>
						Clear filter
						</button>
					)
					: ""}
			</div>
			<BugListView 
				bugs={this.state.bugs} 
				sortAscending={this.state.sortAscending}
				itemClickReceiver={(bugId) => this.receiveItemClick(bugId)}
			/>
		</div>);
	}

	getSearchPanelDisplay() {
		return this.state.searchPanelVisible ? null : "none";
	}

	changeSortDirection() {
		this.setState({sortAscending: ! this.state.sortAscending});
	}

	receiveSearchClick() {
		this.setState({searchPanelVisible: true});
	}

	receiveSearchPanelClickBack() {
		this.setState({searchPanelVisible: false});
	}

	requestBugs() {
		if (this.state.filterString.length === 0)
			Api.LoadBugList(data => this.receiveBugs(data));
		else {
			console.log(this.state.filterString);
			Api.LoadBugListFiltered(this.state.filterString, data => this.receiveBugs(data));
		}
	}

	receiveBugs(data) {
		this.setState({bugs: data});
	}

	receiveSearchAct(keywords) {
		this.setState(
			{
				searchPanelVisible: false,
				filterString: keywords,
				bugs: [],
			},
			() => this.requestBugs());
	}

	receiveRefreshClick() {
		this.setState({bugs: []});
		this.requestBugs();
	}

	receiveItemClick(bugId) {
		window.open(AppURL + "/bug/" + bugId);
	}

}

export default BugListPanel;