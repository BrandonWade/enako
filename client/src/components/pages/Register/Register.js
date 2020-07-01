import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import createAccount from '../../../effects/createAccount';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import ValidationRow from '../../atoms/ValidationRow';
import Button from '../../atoms/Button';
import './Register.scss';
import { ValidateUsernameLength, ValidateUsernameCharacters } from '../../../validators/username';
import { ValidateEmailFormat } from '../../../validators/email';
import { ValidatePasswordLength, ValidatePasswordCharacters, ValidatePasswordsMatch } from '../../../validators/password';

const Register = () => {
    const history = useHistory();
    const [username, setUsername] = useState('foobar');
    const [email, setEmail] = useState('foo@bar.net');
    const [password, setPassword] = useState('testpassword123');
    const [confirmPassword, setConfirmPassword] = useState('testpassword123');

    const onCreateAccount = async () => {
        const data = {
            username,
            email,
            password,
            confirm_password: confirmPassword,
        };

        const response = await createAccount(data);
        if (response.errors) {
            console.error(response); // TODO: Implement proper error handling
            return;
        }

        history.push('/login');
    };

    return (
        <div className='Register'>
            <div className='Register-content'>
                <Logo />
                <Card className='Register-form'>
                    <div className='Register-formGrid'>
                        <InputField label='Username' value={username} onChange={e => setUsername(e.target.value)} />
                        <div className='Register-validatorRules'>
                            <ValidationRow valid={ValidateUsernameLength(username)} description='Username is between 5 and 32 characters' />
                            <ValidationRow valid={ValidateUsernameCharacters(username)} description='Username contains only numbers and letters' />
                        </div>
                        <InputField label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                        <div className='Register-validatorRules'>
                            <ValidationRow valid={ValidateEmailFormat(email)} description='Email is well formatted' />
                        </div>
                        <InputField type='password' label='Password' value={password} onChange={e => setPassword(e.target.value)} />
                        <div className='Register-validatorRules Register-passwordRules'>
                            <ValidationRow valid={ValidatePasswordLength(password)} description='Password is between 15 and 50 characters' />
                            <ValidationRow
                                valid={ValidatePasswordCharacters(password)}
                                description='Password contains only numbers, letters, and valid symbols: ! @ # $ % ^ * _'
                            />
                            <ValidationRow
                                valid={ValidatePasswordsMatch(password, confirmPassword)}
                                description='Password and Confirm Password match'
                            />
                        </div>
                        <InputField
                            type='password'
                            label='Confirm Password'
                            value={confirmPassword}
                            onChange={e => setConfirmPassword(e.target.value)}
                        />
                    </div>
                    <div className='Register-buttons'>
                        <Link to='/login'>
                            <Button text='Cancel' />
                        </Link>
                        <Button color='orange' text='Create' onClick={onCreateAccount} />
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default Register;
