import React from 'react';
import './Category.css';

const Category = ({ type }) => {
    return <div className={`Category Category--${type}`}>{type}</div>;
};

export default Category;
