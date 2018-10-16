import React, { Component } from 'react';
import withFormField from './FormField';

class SelectField extends Component {
    render() {
        return (
            <select
                name={this.props.name}
                value={this.props.value}
                onChange={this.props.onChange}
            >
                {this.props.children}
            </select>
        );
    }
}

export default withFormField(SelectField);
