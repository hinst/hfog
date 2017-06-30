import React from 'react';
import { Link } from 'react-router-dom';
import * as AppHeader from "./AppHeader.js";
import * as Api from './Api';

class BugView extends React.Component {

    constructor(props) {
        super(props);
        document.title = "Bug";
    }

    render() {
        return (
            <div>
                {AppHeader.AppHeaderPanel()}
                <Link className="w3-btn w3-black" to="/"> ← Bug list</Link>
                <div style={{display: "inline-block", minWidth: "10px"}}/>
                Bug {this.props.match.params.bugId}
            </div>
        );
    }

    requestContent() {
        Api.LoadBug(this.props.match.params.bugId, (data) => this.receiveContent(data));
    }

    receiveContent(data) {
    }

}

export default BugView;