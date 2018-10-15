import React, { Component } from 'react';
import Card from '../components/Card';
import InputField from '../components/InputField';
import Button from '../components/Button';
import '../css/Login.css';

class Login extends Component {
    render() {
        return (
            <div className='Login'>
                <div className='Login-content'>
                    <Card heading='Enako'>
                        <InputField label='Username' />
                        <InputField label='Password' />
                        <div className='Login-formButtons'>
                            <Button
                                main={true}
                                text='Submit'
                            />
                        </div>
                    </Card>
                </div>
            </div>
        );
    }
}

export default Login;
