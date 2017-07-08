import React from 'react';
import FileURL from './FileURL';

export function CreateAttachmentsView(props) {
    var attachmentBars = props.attachments.map((attachmentData, attachmentDataIndex) => {
        return (
            <span style={{marginRight: "8px"}}>
                <a href="#">
                    <img 
                        alt="[F]"
                        src={FileURL + "/if_323-Document_2124302.svg"} width="16" height="16"
                    />
                    <span style={{marginLeft: "4px"}}/>
                    {attachmentData.FileName}
                </a>
            </span>
        );
    });
    return <div className="w3-panel">
        {attachmentBars}
    </div>
}

export default CreateAttachmentsView;