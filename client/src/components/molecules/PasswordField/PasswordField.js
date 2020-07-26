import React from 'react';
import InputField from '../InputField';
import './PasswordField.scss';

const okayLength = 20;
const strongLength = 30;

const PasswordField = ({ type = 'password', value = '', name = '', label = '', autoComplete = '', onChange = () => {} }) => {
    let passwordClass = '';
    let passwordDescription = '';

    if (type === 'password') {
        if (value.length < okayLength) {
            passwordClass = 'PasswordField--weak';
            passwordDescription = 'Weak';
        } else if (value.length < strongLength) {
            passwordClass = 'PasswordField--okay';
            passwordDescription = 'Okay';
        } else {
            passwordClass = 'PasswordField--strong';
            passwordDescription = 'Strong';
        }
    }

    return (
        <InputField
            type='password'
            name={name}
            value={value}
            className={passwordClass}
            label={label}
            description={passwordDescription}
            autoComplete={autoComplete}
            onChange={onChange}
        />
    );
};

export default PasswordField;
