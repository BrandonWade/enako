import React from 'react';
import './Card.scss';

const Card = ({ className = '', heading = '', children = [] }) => {
    return (
        <div className={`Card ${className}`}>
            {heading && <h2 className='Card-heading'>{heading}</h2>}
            {children}
        </div>
    );
};

export default Card;
