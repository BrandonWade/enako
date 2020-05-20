import React from 'react';
import './RoundButton.css';

const RoundButton = ({ text }) => {
    return (
        <div className='RoundButton'>
            <span className='RoundButton-text'>{text}</span>
        </div>
    );
};

export default RoundButton;
