import React from 'react';
import CreateAttachmentsView from './AttachmentsView';

class BugEventView extends React.PureComponent {

    render() {
        return (
            <div style={{lineHeight: 1}}>
                <hr/>
                <div className="w3-panel w3-leftbar w3-border-black">
                    {this.getMomentString()}
		        	<div style={{display: "inline-block", minWidth: "4px"}}/>
                    <span className="w3-light-gray" style={{padding: "4px"}}>{this.props.person}</span>
        			<div style={{display: "inline-block", minWidth: "4px"}}/>
                    {this.props.verb}
                </div>
                {this.props.text.trim().length > 0
                    ? this.getVisibleTextPanel()
                    : ""
                }
                {this.getAttachmentsView()}
            </div>
        );
    }

    getMomentString() {
        const date = new Date(this.props.moment);
        const s = (text) => {
            text = "" + text;
            while (text.length < 2)
                text = "0" + text;
            return text;
        };
        return "" + date.getFullYear() + "." + s(date.getMonth() + 1) + "." + s(date.getDay()) + " " + s(date.getHours()) + ":" + s(date.getMinutes());
    }

    getVisibleText() {
        let text = this.props.text.trim();
        text = text.split("\n");
        text = text.map((item) => item.trim());
        let items = [];
        for (let i = 0; i < text.length; i++) {
            items.push(<span key={"" + i + "s"}>{text[i]}</span>);
            items.push(<br key={"" + i + "br"}/>);
        }
        return items;
    }

    getVisibleTextPanel() {
        return (
            <div className="w3-panel" style={{marginTop: 0, lineHeight: 1.2}}>
                {this.getVisibleText()}
            </div>
        );
    }

    getAttachmentsView() {
        if (this.props.attachments != null && this.props.attachments.length > 0) {
            return <CreateAttachmentsView attachments={this.props.attachments}/>;
        } else {
            return "";
        }
    }

} 

export default BugEventView;