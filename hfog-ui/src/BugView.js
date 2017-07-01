import React from 'react';
import { Link } from 'react-router-dom';
import * as AppHeader from "./AppHeader.js";
import * as Api from './Api';
import AppURL from './AppURL';

class BugView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            title: "",
        };
        document.title = "Bug";
        this.requestContent();
    }

    render() {
        return (
            <div>
                {AppHeader.AppHeaderPanel()}
                <Link className="w3-btn w3-black" to={AppURL + "/"}> ← Bug list</Link>
                <div style={{display: "inline-block", minWidth: "8px"}}/>
                Bug 
                <div style={{display: "inline-block", minWidth: "4px"}}/>
                <button className="w3-btn w3-black">{this.props.match.params.bugId}</button>
                <div style={{display: "inline-block", minWidth: "4px"}}/>
                {this.state.title} 
            </div>
        );
    }

    requestContent() {
        Api.LoadBug(this.props.match.params.bugId, (data) => this.receiveContent(data));
    }

    receiveContent(data) {
        console.log(data);
        this.setState({
            title: data.Title,
        });
    }

}

export default BugView;