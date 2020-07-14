import React, { useState } from 'react';
import './Message.scss';

const Message = props => {
    const [visible, setVisible] = useState(true);

    const onDismiss = () => {
        setVisible(false);
    };

    return (
        <>
            {visible ? (
                <div className={`Message ${props.className || ''} Message--${props.type || ''}`}>
                    <button className='Message-dismissButton' onClick={onDismiss}>
                        &times;
                    </button>
                    <p className='Message-content'>{props.text || ''}</p>
                </div>
            ) : null}
        </>
    );
};

export default Message;
