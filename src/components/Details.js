import React, { Component } from 'react';
import DetailList from './DetailList';
import '../css/Details.css';

const payments = [
    {
        id: 1,
        type: 'one-time',
        description: 'went out for lunch',
        colour: '#ffb3ba',
        amount: '16.80',
    },
    {
        id: 2,
        type: 'recurring',
        description: 'paid phone bill for next 2 months',
        colour: '#bae1ff',
        amount: '120.58',
    },
    {
        id: 3,
        type: 'one-time',
        description: 'went to a movie',
        colour: '#ffdfba',
        amount: '11.50',
    }
];

class Details extends Component {
    render() {
        return (
            <div className='Details'>
                <h1 className='Details-heading'>Breakdown for {this.props.selectedDate}</h1>
                <DetailList
                    payments={payments}
                />
            </div>
        );
    }
}

export default Details;
