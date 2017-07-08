import React from 'react';

export function CreateAttachmentsView(props) {
    var attachmentBars = props.attachments.map((attachmentData) => {
        return <span style={{marginLeft: "4px"}}>| {attachmentData.FileName}</span>
    });
    return <div className="w3-panel">
        {attachmentBars}
    </div>
}