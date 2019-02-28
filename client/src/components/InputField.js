import React, { Component } from 'react';
import withFormField from './FormField';
import '../css/InputField.css';

class InputField extends Component {
    render() {
        return (
            <input
                type='text'
                name={this.props.name}
                value={this.props.value}
                onChange={this.props.onChange}
            />
        );
    }
}

export default withFormField(InputField);
