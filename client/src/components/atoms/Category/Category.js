import React from 'react';
import './Category.scss';

const Category = ({ type }) => {
    return <div className={`category category--${type}`}>{type}</div>;
};

export default Category;
