import React from 'react';
import withFormField from '../../hocs/withFormField';
import './InputField.scss';

const InputField = props => {
    return <input type='text' name={props.name} value={props.value} onChange={props.onChange} />;
};

export default withFormField(InputField);
