import React from 'react';
import * as Api from "./Api";
import BugSearchPanel from "./BugSearchPanel";
import BugListView from "./BugListView";
import AppURL from './AppURL';
import * as AccessKey from './AccessKey';

class BugListPanel extends React.Component {

	constructor(props) {
		super(props);
		this.state = {
			bugs: [],
			sortedBugs: [],
      		sortAscending: false,
			searchPanelVisible: false,
			filterString: "",
			pageNumber: 0,
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
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>◄</button>
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>►</button>
				<span style={{marginLeft: "4px"}}></span>
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
				bugs={this.state.sortedBugs} 
				sortAscending={this.state.sortAscending}
				itemClickReceiver={(bugId) => this.receiveItemClick(bugId)}
			/>
		</div>);
	}

	getSearchPanelDisplay() {
		return this.state.searchPanelVisible ? null : "none";
	}

	changeSortDirection() {
		this.setState({sortAscending: ! this.state.sortAscending},
			() => this.sortBugs());
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
		this.setState({bugs: data}, () => this.sortBugs());
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
		window.open(AppURL + "/bug/" + bugId + "?" + AccessKey.GetURL());
	}

	changePage(delta) {
		let pageNumber = this.state.pageNumber;
		pageNumber = pageNumber + delta;
	}

	checkPageNumberValid(pageNumber) {

	}

	getSortedBugs() {
		const sortAscending = this.state.sortAscending;
		const bugs = this.state.bugs.slice();
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

	sortBugs() {
		this.setState({sortedBugs: this.getSortedBugs()});
	}

}

export default BugListPanel;