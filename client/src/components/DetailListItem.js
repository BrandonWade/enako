import React, { Component } from 'react';
import '../css/DetailListItem.css';

class DetailListItem extends Component {
    render() {
        return (
            <li
                className={`DetailListItem ${this.props.colour ? `Payments__${this.props.colour}` : ''}`}
            >
                <div className='DetailListItem-name'>
                    {this.props.name}
                </div>
                <div className='DetailListItem-amount'>
                    {this.props.amount}
                </div>
            </li>
        );
    }
}

export default DetailListItem;
