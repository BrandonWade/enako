import React from 'react';
import './withFormField.css';

const withFormField = BaseComponent => {
    const FormField = props => {
        return (
            <section className='FormField'>
                {props.label && <label className='FormField-label'>{props.label}</label>}
                <BaseComponent {...props}>{props.children}</BaseComponent>
                {props.description && <div className='FormField-description'>{props.description}</div>}
            </section>
        );
    };

    return FormField;
};

export default withFormField;
