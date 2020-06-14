import React, { useContext } from 'react';
import CategoryContext from '../../../contexts/CategoryContext';
import './Category.scss';

const Category = ({ id }) => {
    const categories = useContext(CategoryContext);
    const category = categories.find(c => c.id === id);

    return <div className={`Category Category--${category.name}`}>{category.name}</div>;
};

export default Category;
