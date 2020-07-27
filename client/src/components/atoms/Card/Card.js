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

export const CardSection = ({ heading = '', description = '', children = [] }) => {
    return (
        <section className='CardSection'>
            <h6 className='CardSection-heading'>{heading}</h6>
            <div className='CardSection-description'>{description}</div>
            {children}
        </section>
    );
};

Card.Section = CardSection;

export default Card;
