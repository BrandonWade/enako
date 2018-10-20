import React, { Component } from 'react';
import '../css/Category.css';

class Category extends Component {
    render() {
        return (
            <div className={`Category Category--${this.props.type}`}>
                {this.props.type}
            </div>
        );
    }
}

export default Category;
