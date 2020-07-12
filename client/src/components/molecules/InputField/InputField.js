import React from 'react';
import withFormField from '../../hocs/withFormField';
import './InputField.scss';

const InputField = props => {
    return (
        <input
            type={props.type || 'text'}
            name={props.name}
            value={props.value}
            step={props.type === 'number' ? '0.01' : ''}
            className={`${props.className || ''}`}
            description={props.description || ''}
            autoComplete={props.autoComplete || ''}
            onChange={props.onChange}
        />
    );
};

export default withFormField(InputField);
