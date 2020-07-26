import React from 'react';
import './Logo.scss';

const Logo = ({ className = '' }) => {
    return <h1 className={`Logo ${className}`}>Enako</h1>;
};

export default Logo;
