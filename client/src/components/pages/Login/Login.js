import React from 'react';
import { Link } from 'react-router-dom';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Login.scss';

const Login = () => {
    return (
        <div className='login'>
            <Card heading='Enako' className='login__content'>
                <InputField label='Username' />
                <InputField type='password' label='Password' />
                <Button full color='orange' text='Login' />
                <div className='login__separator'>or</div>
                <Link to='/register'>
                    <Button full color='blue' text='Create Account' />
                </Link>
            </Card>
        </div>
    );
};

export default Login;
