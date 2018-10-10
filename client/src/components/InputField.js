import React, { Component } from 'react';
import withFormField from './FormField';

class InputField extends Component {
    render() {
        return (
            <input type='text' />
        );
    }
}

export default withFormField(InputField);
