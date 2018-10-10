import React, { Component } from 'react';
import withFormField from './FormField';

class SelectField extends Component {
    render() {
        return (
            <select>
                {this.props.children}
            </select>
        );
    }
}

export default withFormField(SelectField);
