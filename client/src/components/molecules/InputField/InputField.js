import React from 'react';
import withFormField from '../../hocs/withFormField';
import './InputField.scss';

const InputField = ({
    type = 'text',
    name = '',
    value = '',
    className = '',
    description = '',
    autoComplete = '',
    onChange = () => {},
}) => {
    return (
        <input
            type={type}
            name={name}
            value={value}
            step={type === 'number' ? '0.01' : ''}
            className={`${className}`}
            description={description}
            autoComplete={autoComplete}
            onChange={onChange}
        />
    );
};

export default withFormField(InputField);
