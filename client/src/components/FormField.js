import React, { Component } from 'react';
import '../css/FormField.css';

const withFormField = (BaseComponent) => {
    class FormField extends Component {
        render() {
            return (
                <section className='FormField'>
                    <label className='FormField-label'>
                        {this.props.label}
                    </label>
                    <BaseComponent>
                        {this.props.children}
                    </BaseComponent>
                    <div className='FormField-description'>
                        {this.props.description}
                    </div>
                </section>
            );
        }
    }

    return FormField;
}

export default withFormField;
