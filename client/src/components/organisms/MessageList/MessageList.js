import React, { useContext, useEffect } from 'react';
import MessageContext from '../../../contexts/MessageContext';
import Message from '../../molecules/Message';
import './MessageList.scss';

const MessageList = () => {
    const { messages, setMessages } = useContext(MessageContext);

    useEffect(() => {
        return () => {
            if (messages.length > 0) {
                setMessages([]);
            }
        };
    }, [messages, setMessages]);

    return (
        <>
            {messages.map(m => (
                <Message key={m.text} type={m.type} text={m.text} />
            ))}
        </>
    );
};

export default MessageList;
