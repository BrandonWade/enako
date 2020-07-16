import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import loginToAccount from '../../../effects/loginToAccount';
import Logo from '../../atoms/Logo';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import MessageList from '../../organisms/MessageList';

import Button from '../../atoms/Button';
import './Login.scss';

const Login = props => {
    const history = useHistory();
    const [email, setEmail] = useState('foo@bar.net');
    const [password, setPassword] = useState('testpassword123');

    const onLogin = async () => {
        const data = {
            email,
            password,
        };

        const response = await loginToAccount(data);
        if (response?.messages?.length > 0) {
            props.setMessages(response.messages);
            return;
        }

        props.setAuthenticated(true);
        history.push('/');
    };

    return (
        <div className='Login'>
            <div className='Login-content'>
                <Logo />
                <Card className='Login-form'>
                    <MessageList />
                    <InputField type='text' label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                    <InputField
                        type='password'
                        label='Password'
                        value={password}
                        formClassName='Login-password'
                        autoComplete='current-password'
                        onChange={e => setPassword(e.target.value)}
                    />
                    <div className='Login-forgotPassword'>
                        <Link to='/password'>
                            <span>Forgot your password?</span>
                        </Link>
                    </div>
                    <Button full color='orange' text='Login' onClick={onLogin} />
                    <p className='Login-createAccount'>
                        Don't have an account yet?
                        <Link to='/register'>
                            <span className='Login-createAccountLink'>Sign up!</span>
                        </Link>
                    </p>
                </Card>
            </div>
        </div>
    );
};

export default Login;
