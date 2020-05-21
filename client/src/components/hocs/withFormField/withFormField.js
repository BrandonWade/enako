import React from 'react';
import './withFormField.css';

const withFormField = BaseComponent => {
    const FormField = props => {
        return (
            <div className='form-field'>
                {props.label && <label className='form-field__label'>{props.label}</label>}
                <BaseComponent {...props}>{props.children}</BaseComponent>
                {props.description && <div className='form-field__description'>{props.description}</div>}
            </div>
        );
    };

    return FormField;
};

export default withFormField;
