import React from 'react';
import './Card.scss';

const Card = props => {
    return (
        <div className={`Card ${props.className || ''}`}>
            {props.heading && <h2 className='Card-heading'>{props.heading}</h2>}
            {props.children}
        </div>
    );
};

export default Card;
