import React from 'react';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Login.scss';

const Login = () => {
    return (
        <div className='login'>
            <Card heading='Enako' className='login__content'>
                <InputField label='Username' />
                <InputField label='Password' />
                <Button primary text='Submit' className='login__button' />
            </Card>
        </div>
    );
};

export default Login;
