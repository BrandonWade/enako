import React, { Component } from 'react';
import Card from '../Card';
import InputField from '../InputField';
import Button from '../Button';
import '../../css/Login.css';

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
