import React from 'react';
import './Category.css';

const Category = ({ type }) => {
    return <div className={`category category--${type}`}>{type}</div>;
};

export default Category;
