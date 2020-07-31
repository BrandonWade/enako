import React, { useState, useContext } from 'react';
import { Link } from 'react-router-dom';
import { ValidateEmailFormat } from '../../../validators/email';
import MessageContext from '../../../contexts/MessageContext';
import forgotPassword from '../../../effects/forgotPassword';
import InputField from '../../molecules/InputField';
import MessageList from '../../organisms/MessageList';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import Button from '../../atoms/Button';
import './ForgotPassword.scss';

const ForgotPassword = () => {
    const [email, setEmail] = useState('foo@bar.net');
    const { setMessages } = useContext(MessageContext);

    const isInvalid = !ValidateEmailFormat(email);

    const onSubmit = async () => {
        const response = await forgotPassword({ email });
        if (response?.messages?.length > 0) {
            setMessages(response.messages);
            return;
        }
    };

    return (
        <div className='ForgotPassword'>
            <div className='ForgotPassword-content'>
                <Logo />
                <Card className='ForgotPassword-form'>
                    <MessageList />
                    <p className='ForgotPassword-instructions'>
                        Enter your email below and we'll send a password reset link to your email.
                    </p>
                    <InputField label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                    <div className='ForgotPassword-buttons'>
                        <Link to='/login'>
                            <Button text='Cancel' />
                        </Link>
                        <Button color='orange' text='Send' onClick={onSubmit} disabled={isInvalid} />
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default ForgotPassword;
