import React from 'react';
import './ValidationRow.scss';

const ValidationRow = ({ validator, value, description }) => {
    const isValid = validator(value);

    const getValidationStyle = () => {
        if (!value.length) {
            return 'ValidationRow--default';
        }

        return isValid ? 'ValidationRow--valid' : 'ValidationRow--invalid';
    };

    // TODO: Find icons
    const renderIcon = () => {
        if (!value.length) {
            return '?';
        }

        return isValid ? 'v' : 'x';
    };

    return (
        <div className='ValidationRow'>
            <div className={`ValidationRow-icon ${getValidationStyle()}`}>{renderIcon(validator, value)}</div>
            <div className={`ValidationRow-message ${getValidationStyle()}`}>{description}</div>
        </div>
    );
};

export default ValidationRow;
