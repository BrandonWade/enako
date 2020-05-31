import React from 'react';
import './Button.scss';

const Button = props => {
    const color = props.color ? `button--${props.color}` : '';
    const full = props.full ? 'button--full' : '';

    return (
        <button className={`button ${color} ${full} ${props.className || ''}`} onClick={props.onClick}>
            {props.text}
        </button>
    );
};

export default Button;
