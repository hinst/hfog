import React from 'react';
import { Link } from 'react-router-dom';
import * as AppHeader from "./AppHeader.js";

class BugView extends React.Component {

    constructor(props) {
        super(props);
        document.title = "Bug";
    }

    render() {
        return (
            <div>
                {AppHeader.AppHeaderPanel()}
                <Link to="/"> ← List</Link>
            </div>
        );
    }

}

export default BugView;