import React from 'react';
import { CheckIcon, CrossIcon } from '../Icons';
import './ValidationRow.scss';

const ValidationRow = ({ valid, description }) => {
    const validClass = valid ? 'u-valid' : 'u-invalid';
    const validIcon = valid ? <CheckIcon className={validClass} /> : <CrossIcon className={validClass} />;

    return (
        <div className='ValidationRow'>
            <div className={`ValidationRow-icon ${validClass}`}>{validIcon}</div>
            <div className={`ValidationRow-message ${validClass}`}>{description}</div>
        </div>
    );
};

export default ValidationRow;
