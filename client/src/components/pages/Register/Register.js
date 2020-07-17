import React, { useState, useContext } from 'react';
import { Link, useHistory } from 'react-router-dom';
import createAccount from '../../../effects/createAccount';
import resetPassword from '../../../effects/resetPassword';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import PasswordField from '../../molecules/PasswordField';
import ValidationRow from '../../atoms/ValidationRow';
import MessageList from '../../organisms/MessageList';
import Button from '../../atoms/Button';
import MessageContext from '../../../contexts/MessageContext';
import { ValidateEmailFormat } from '../../../validators/email';
import { ValidatePasswordLength, ValidatePasswordCharacters, ValidatePasswordsMatch } from '../../../validators/password';
import './Register.scss';

const Register = props => {
    const history = useHistory();
    const [email, setEmail] = useState('foo@bar.net');
    const [password, setPassword] = useState('testpassword123');
    const [confirmPassword, setConfirmPassword] = useState('testpassword123');
    const { setMessages } = useContext(MessageContext);

    const validEmailFormat = ValidateEmailFormat(email);
    const validPasswordLength = ValidatePasswordLength(password);
    const validPasswordCharacters = ValidatePasswordCharacters(password);
    const validPasswordsMatch = ValidatePasswordsMatch(password, confirmPassword);

    const isPasswordValid = validPasswordLength && validPasswordCharacters && validPasswordsMatch;
    const isEnabled = validEmailFormat && isPasswordValid;

    const renderEmail = () => {
        return !props.passwordReset ? (
            <>
                <InputField label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                <div className='Register-validatorRules'>
                    <ValidationRow valid={validEmailFormat} description='Email is well formatted' />
                </div>
            </>
        ) : null;
    };

    const renderSubmit = () => {
        if (props.passwordReset) {
            return <Button color='orange' text='Reset' onClick={onResetPassword} disabled={!isPasswordValid} />;
        } else {
            return <Button color='orange' text='Create' onClick={onCreateAccount} disabled={!isEnabled} />;
        }
    };

    const onResetPassword = async () => {
        setMessages([]);

        const data = {
            password,
            confirm_password: confirmPassword,
        };

        const response = await resetPassword(data);
        if (response?.messages?.length > 0) {
            setMessages(response.messages);
            return;
        }

        history.push('/login');
    };

    const onCreateAccount = async () => {
        setMessages([]);

        const data = {
            email,
            password,
            confirm_password: confirmPassword,
        };

        const response = await createAccount(data);
        if (response?.messages?.length > 0) {
            setMessages(response.messages);
            return;
        }

        history.push('/login');
    };

    return (
        <div className='Register'>
            <div className='Register-content'>
                <Logo />
                <Card className='Register-form'>
                    <MessageList />
                    <div className='Register-formGrid'>
                        {renderEmail()}
                        <PasswordField
                            type='password'
                            label='Password'
                            value={password}
                            autoComplete='current-password'
                            onChange={e => setPassword(e.target.value)}
                        />
                        <div
                            className={`Register-validatorRules ${
                                props.passwordReset ? 'Register-passwordRules--passwordReset' : 'Register-passwordRules'
                            }`}
                        >
                            <ValidationRow valid={validPasswordLength} description='Password is between 15 and 50 characters' />
                            <ValidationRow
                                valid={validPasswordCharacters}
                                description='Password contains only numbers, letters, and valid symbols: ! @ # $ % ^ * _'
                            />
                            <ValidationRow valid={validPasswordsMatch} description='Password and Confirm Password match' />
                        </div>
                        <PasswordField
                            type='password'
                            label='Confirm Password'
                            value={confirmPassword}
                            autoComplete='new-password'
                            onChange={e => setConfirmPassword(e.target.value)}
                        />
                    </div>
                    <div className='Register-buttons'>
                        <Link to='/login'>
                            <Button text='Cancel' />
                        </Link>
                        {renderSubmit()}
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default Register;
