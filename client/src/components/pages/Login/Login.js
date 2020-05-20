import React from 'react';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Login.css';

const Login = () => {
    return (
        <div className='Login'>
            <div className='Login-content'>
                <Card heading='Enako'>
                    <InputField label='Username' />
                    <InputField label='Password' />
                    <div className='Login-formButtons'>
                        <Button main={true} text='Submit' />
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default Login;
