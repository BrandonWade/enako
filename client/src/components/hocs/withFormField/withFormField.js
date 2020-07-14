import React from 'react';
import './withFormField.scss';

const withFormField = BaseComponent => {
    const FormField = props => {
        return (
            <div className={`FormField ${props.formClassName || ''}`}>
                {props.label && <label className='FormField-label'>{props.label}</label>}
                <BaseComponent {...props}>{props.children}</BaseComponent>
                {props.description && <div className='FormField-description'>{props.description}</div>}
            </div>
        );
    };

    return FormField;
};

export default withFormField;
