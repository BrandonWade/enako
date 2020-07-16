import React, { useContext } from 'react';
import MessageContext from '../../../contexts/MessageContext';
import Message from '../../molecules/Message';
import './MessageList.scss';

const MessageList = () => {
    const messages = useContext(MessageContext);

    return (
        <>
            {messages.map(m => (
                <Message key={m.text} type={m.type} text={m.text} />
            ))}
        </>
    );
};

export default MessageList;
