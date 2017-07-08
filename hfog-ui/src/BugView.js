import React from 'react';
import { Link } from 'react-router-dom';
import * as AppHeader from "./AppHeader";
import * as Api from './Api';
import AppURL from './AppURL';
import BugEventView from './BugEventView';
import * as AccessKey from './AccessKey';

class BugView extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            title: "",
            events: [],
        };
        document.title = "FB-Bug " + this.props.match.params.bugId;
        this.requestContent();
    }

    render() {
        return (
            <div>
                {AppHeader.AppHeaderPanel()}
                <Link className="w3-btn w3-black" to={AppURL + "/?" + AccessKey.GetURL()}> ← Bug list</Link>
                <div style={{display: "inline-block", minWidth: "8px"}}/>
                Bug 
                <div style={{display: "inline-block", minWidth: "4px"}}/>
                <button className="w3-btn w3-black">{this.props.match.params.bugId}</button>
                <div style={{display: "inline-block", minWidth: "4px"}}/>
                {this.state.title} 
                <br/>
                {this.renderEvents()}
            </div>
        );
    }

    renderEvents() {
        const items = this.state.events.map(
            (event, i) => {
                return (
                    <BugEventView 
                        key={this.props.match.params.bugId + "-" + i}
                        moment={event.Moment}
                        person={event.Person}
                        verb={event.Verb}
                        text={event.Text}
                        html={event.HTML}
                        attachments={event.Attachments}
                    />
                );
            });
        return items;
    }

    requestContent() {
        Api.LoadBug(this.props.match.params.bugId, (data) => this.receiveContent(data));
    }

    receiveContent(data) {
        this.setState({
            title: data.Title,
            events: data.Events,
        });
    }

}

export default BugView;