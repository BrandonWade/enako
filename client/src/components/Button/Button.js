import React from 'react';
import './Button.css';

const Button = (props) => {
    const getButtonStyles = () => {
        return props.main ? ' Button__main' : '';
    };

    return <button className={`Button ${getButtonStyles()}`}>{props.text}</button>;
};

export default Button;
