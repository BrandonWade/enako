import React from 'react';
import './Button.scss';

const Button = ({ color = '', full = '', disabled = false, className = '', onClick = () => {}, text = '' }) => {
    const colorClass = color ? `Button--${color}` : '';
    const fullClass = full ? 'Button--full' : '';
    const isDisabled = disabled ?? false;

    return (
        <button className={`Button ${colorClass} ${fullClass} ${className}`} onClick={onClick} disabled={isDisabled}>
            {text}
        </button>
    );
};

export default Button;
