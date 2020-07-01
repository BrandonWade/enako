import React from 'react';
import './Button.scss';

const Button = props => {
    const color = props.color ? `Button--${props.color}` : '';
    const full = props.full ? 'Button--full' : '';
    const isDisabled = props.disabled ?? false;

    return (
        <button className={`Button ${color} ${full} ${props.className || ''}`} onClick={props.onClick} disabled={isDisabled}>
            {props.text}
        </button>
    );
};

export default Button;
