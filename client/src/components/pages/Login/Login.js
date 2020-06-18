import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import loginToAccount from '../../../effects/loginToAccount';
import fetchCategories from '../../../effects/fetchCategories';
import fetchExpenses from '../../../effects/fetchExpenses';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Login.scss';

const Login = props => {
    const history = useHistory();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const onLogin = async () => {
        const data = {
            username,
            password,
        };

        const response = await loginToAccount(data);
        if (response.errors) {
            console.error(response); // TODO: Implement proper error handling
            return;
        }

        props.setAuthenticated(true);
        history.push('/');
    };

    return (
        <div className='Login'>
            <Card heading='Enako' className='Login-content'>
                <InputField type='text' label='Username' value={username} onChange={e => setUsername(e.target.value)} />
                <InputField type='password' label='Password' value={password} onChange={e => setPassword(e.target.value)} />
                <Button full color='orange' text='Login' onClick={onLogin} />
                <div className='Login-separator'>or</div>
                <Link to='/register'>
                    <Button full color='blue' text='Create Account' />
                </Link>
            </Card>
        </div>
    );
};

export default Login;
