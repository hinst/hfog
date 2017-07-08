import React from 'react';
import FileURL from './FileURL';
import ApiURL from './ApiURL';
import * as AccessKey from './AccessKey';

export function CreateAttachmentsView(props) {
    var attachmentBars = props.attachments.map((attachmentData, attachmentDataIndex) => {
        return (
            <span style={{marginRight: "8px"}}>
                <a 
                    href={ApiURL + "/getAtt?" + AccessKey.GetURL() + "&key=" + encodeURIComponent(attachmentData.KeyURL)}
                >
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