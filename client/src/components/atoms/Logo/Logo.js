import React from 'react';
import './Logo.scss';

const Logo = props => {
    return <h1 className={`Logo ${props.className || ''}`}>Enako</h1>;
};

export default Logo;
