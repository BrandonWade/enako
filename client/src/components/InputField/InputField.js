import React from 'react';
import withFormField from '../withFormField';
import './InputField.css';

const InputField = (props) => {
    return <input type='text' name={props.name} value={props.value} onChange={props.onChange} />;
};

export default withFormField(InputField);
