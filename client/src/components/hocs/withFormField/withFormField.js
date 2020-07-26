import React from 'react';
import './withFormField.scss';

const withFormField = BaseComponent => {
    const FormField = props => {
        const { formClassName = '', label = '', children = [], description = '' } = props;
        return (
            <div className={`FormField ${formClassName}`}>
                {label && <label className='FormField-label'>{label}</label>}
                <BaseComponent {...props}>{children}</BaseComponent>
                {description && <div className='FormField-description'>{description}</div>}
            </div>
        );
    };

    return FormField;
};

export default withFormField;
