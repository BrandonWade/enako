import React from 'react';
import './Button.scss';

const Button = props => {
    const getButtonType = () => {
        return props.primary ? ' button--primary' : '';
    };

    return (
        <button className={`button ${getButtonType()} ${props.className || ''}`} onClick={props.onClick}>
            {props.text}
        </button>
    );
};

export default Button;
