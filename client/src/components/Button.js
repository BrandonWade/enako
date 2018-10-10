import React, { Component } from 'react';
import '../css/Button.css';

class Button extends Component {
    getButtonStyles = () => {
        return this.props.main ? ' Button__main' : '';
    };

    render() {
        return (
            <button className={`Button ${this.getButtonStyles()}`}>
                {this.props.text}
            </button>
        );
    }
}

export default Button;
