import React from 'react';
import withFormField from '../../hocs/withFormField';
import './SelectField.css';

const SelectField = props => {
    return (
        <select name={props.name} value={props.value} onChange={props.onChange}>
            {props.children}
        </select>
    );
};

export default withFormField(SelectField);
