import React, { useContext } from 'react';
import MessageContext from '../../../contexts/MessageContext';
import Message from '../../molecules/Message';
import './MessageList.scss';

const MessageList = () => {
    const messages = useContext(MessageContext);

    return (
        <>
            {messages.errors.map(e => (
                <Message key={e} type='error' text={e} />
            ))}
            {messages.messages.map(m => (
                <Message key={m} type='info' text={m} />
            ))}
        </>
    );
};

export default MessageList;
