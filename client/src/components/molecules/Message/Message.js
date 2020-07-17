import React, { useState, useContext } from 'react';
import MessageContext from '../../../contexts/MessageContext';
import './Message.scss';

const Message = props => {
    const [visible, setVisible] = useState(true);
    const { messages, setMessages } = useContext(MessageContext);

    const onDismiss = () => {
        setVisible(false);
        setMessages(messages.filter(m => m.text !== props.text));
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
