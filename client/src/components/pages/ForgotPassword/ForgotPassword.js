import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import { ValidateEmailFormat } from '../../../validators/email';
import forgotPassword from '../../../effects/forgotPassword';
import './ForgotPassword.scss';

const ForgotPassword = () => {
    const [email, setEmail] = useState('foo@bar.net');

    const isInvalid = !ValidateEmailFormat(email);

    return (
        <div className='ForgotPassword'>
            <div className='ForgotPassword-content'>
                <Logo />
                <Card className='ForgotPassword-form'>
                    <p className='ForgotPassword-instructions'>
                        Enter your email below and we'll send a password reset link to your email.
                    </p>
                    <InputField label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                    <div className='ForgotPassword-buttons'>
                        <Link to='/login'>
                            <Button text='Cancel' />
                        </Link>
                        <Button color='orange' text='Send' onClick={() => forgotPassword({ email })} disabled={isInvalid} />
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default ForgotPassword;
