import React from 'react';
import './Card.css';

const Card = (props) => {
    return (
        <div className='Card'>
            {props.heading && <h2 className='Card-heading'>{props.heading}</h2>}
            {props.children}
        </div>
    );
};

export default Card;
