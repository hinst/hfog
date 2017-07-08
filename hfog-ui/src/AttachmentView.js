import React from 'react';

function CreateAttachmentView(props) {
    var attachmentBars = props.attachments.map((attachmentData) => {
        return <span style={{marginLeft: "4px"}}>| {attachment.FileName}</span>
    });
    return <div className="w3-panel">
        {attachmentBars}
    </div>
}