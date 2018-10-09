import React, { Component } from 'react';
import '../css/Card.css';

class Card extends Component {
    render() {
        return (
            <div className='Card'>
                {this.props.heading && <h2 className='Card-heading'>{this.props.heading}</h2>}
                {this.props.children}
            </div>
        );
    }
}

export default Card;
