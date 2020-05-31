import React from 'react';
import { Link } from 'react-router-dom';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Register.scss';

const Register = () => {
    return (
        <div className='register'>
            <Card className='register__content' heading='Create Account'>
                <InputField label='Username' />
                <InputField label='Email' />
                <InputField type='password' label='Password' />
                <InputField type='password' label='Confirm Password' />
                <div className='register__buttons'>
                    <Link to='/login'>
                        <Button text='Cancel' />
                    </Link>
                    <Button color='orange' text='Create' />
                </div>
            </Card>
        </div>
    );
};

export default Register;
