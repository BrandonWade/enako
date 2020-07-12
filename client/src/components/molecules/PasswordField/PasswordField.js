import React from 'react';
import InputField from '../InputField';
import './PasswordField.scss';

const okayLegth = 20;
const strongLength = 30;

const PasswordField = props => {
    let passwordClass = '';
    let passwordDescription = '';

    if (props.type === 'password') {
        if (props.value.length < okayLegth) {
            passwordClass = 'PasswordField--weak';
            passwordDescription = 'Weak';
        } else if (props.value.length < strongLength) {
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
            name={props.name}
            value={props.value}
            className={passwordClass}
            label={props.label || ''}
            description={passwordDescription || ''}
            autoComplete={props.autoComplete || ''}
            onChange={props.onChange}
        />
    );
};

export default PasswordField;
