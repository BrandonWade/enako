import React, { useState, useContext } from 'react';
import MessageContext from '../../../contexts/MessageContext';
import './Message.scss';

const Message = ({ text = '', className = '', type = '' }) => {
    const [visible, setVisible] = useState(true);
    const { messages, setMessages } = useContext(MessageContext);

    const onDismiss = () => {
        setVisible(false);
        setMessages(messages.filter(m => m.text !== text));
    };

    return (
        <>
            {visible ? (
                <div className={`Message ${className} Message--${type}`}>
                    <button className='Message-dismissButton' onClick={onDismiss}>
                        &times;
                    </button>
                    <p className='Message-content'>{text}</p>
                </div>
            ) : null}
        </>
    );
};

export default Message;
