import React, { useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import createAccount from '../../../effects/createAccount';
import Card from '../../atoms/Card';
import InputField from '../../molecules/InputField';
import Button from '../../atoms/Button';
import './Register.scss';

const Register = () => {
    const history = useHistory();
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');

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
        <div className='register'>
            <Card className='register__content' heading='Create Account'>
                <InputField label='Username' value={username} onChange={e => setUsername(e.target.value)} />
                <InputField label='Email' value={email} onChange={e => setEmail(e.target.value)} />
                <InputField type='password' label='Password' value={password} onChange={e => setPassword(e.target.value)} />
                <InputField type='password' label='Confirm Password' value={confirmPassword} onChange={e => setConfirmPassword(e.target.value)} />
                <div className='register__buttons'>
                    <Link to='/login'>
                        <Button text='Cancel' />
                    </Link>
                    <Button color='orange' text='Create' onClick={onCreateAccount} />
                </div>
            </Card>
        </div>
    );
};

export default Register;
