import React from 'react';
import './Card.css';

const Card = props => {
    return (
        <div className={`card ${props.className || ''}`}>
            {props.heading && <h2 className='card__heading'>{props.heading}</h2>}
            {props.children}
        </div>
    );
};

export default Card;
