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
			bugs: null,
			sortedBugs: null,
      		sortAscending: false,
			searchPanelVisible: false,
			filterString: "",
			pageNumber: 0,
			pageSize: 100,
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
				<button className="w3-btn w3-black" onClick={() => this.changePage(-1)}>◄</button>
				<span className="w3-button w3-black" onClick={() => this.receivePageNumberClick()}>{this.state.pageNumber + 1}</span>
				<button className="w3-btn w3-black" onClick={() => this.changePage(1)}>►</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={() => this.changeSortDirection()}>{this.state.sortAscending ? "▲" : "▼"}</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={() => this.receiveRefreshClick()}>Refresh</button>
				<span style={{marginLeft: "4px"}}></span>
				<button className="w3-btn w3-black" onClick={ () => this.receiveSearchClick() }>Search</button>
				<span style={{marginLeft: "4px"}}></span>
				Showing {this.state.bugs != null ? this.state.bugs.length : "no"} items
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
				bugs={this.getVisibleBugs()} 
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
		this.setState({bugs: null, sortedBugs: null});
		if (this.state.filterString.length === 0)
			Api.LoadBugList(data => this.receiveBugs(data));
		else {
			Api.LoadBugListFiltered(this.state.filterString, data => this.receiveBugs(data));
		}
	}

	receiveBugs(data) {
		this.setState({bugs: data, sortedBugs: null}, () => this.sortBugs());
	}

	receiveSearchAct(keywords) {
		this.setState(
			{
				searchPanelVisible: false,
				filterString: keywords,
				bugs: null, sortedBugs: null,
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
		if (this.checkPageNumberValid(pageNumber)) {
			this.setState({pageNumber: pageNumber});
		}
	}

	checkPageNumberValid(pageNumber) {
		return (this.state.sortedBugs != null) && (0 <= pageNumber) && (pageNumber * this.state.pageSize < this.state.sortedBugs.length);
	}

	getSortedBugs() {
		const sortAscending = this.state.sortAscending;
		const bugs = this.state.bugs != null
			? this.state.bugs.slice()
			: null;
		if (bugs != null) {
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
		}
		return bugs;
	}

	sortBugs() {
		this.setState({pageNumber: 0, sortedBugs: this.getSortedBugs()});
	}

	getVisibleBugs() {
		if (this.state.sortedBugs != null) {
			let start = this.state.pageNumber * this.state.pageSize;
			let end = start + this.state.pageSize;
			return this.state.sortedBugs.slice(start, end);
		} else {
			return null;
		}
	}

	receivePageNumberClick() {
		let pageNumber = this.state.pageNumber;
		if (pageNumber > 0) {
			this.setState({pageNumber: 0});
		} else {
			while (this.checkPageNumberValid(pageNumber + 1))
				pageNumber++;
			this.setState({pageNumber: pageNumber});
		}
	}

}

export default BugListPanel;