import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import { ValidateUsernameLength, ValidateUsernameCharacters } from '../../../validators/username';
import forgotPassword from '../../../effects/forgotPassword';
import './ForgotPassword.scss';

const ForgotPassword = () => {
    const [username, setUsername] = useState('foobar');

    const isInvalid = !(ValidateUsernameLength(username) && ValidateUsernameCharacters(username));

    return (
        <div className='ForgotPassword'>
            <div className='ForgotPassword-content'>
                <Logo />
                <Card className='ForgotPassword-form'>
                    <p className='ForgotPassword-instructions'>Enter your username below and we'll send a password reset link to your email.</p>
                    <InputField label='Username' value={username} onChange={e => setUsername(e.target.value)} />
                    <div className='ForgotPassword-buttons'>
                        <Link to='/login'>
                            <Button text='Cancel' />
                        </Link>
                        <Button color='orange' text='Send' onClick={() => forgotPassword({ username })} disabled={isInvalid} />
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default ForgotPassword;
