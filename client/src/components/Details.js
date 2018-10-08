import React, { Component } from 'react';
import DetailList from './DetailList';
import payments from '../data/SamplePayments';
import '../css/Details.css';

class Details extends Component {
    filterPayments = (date) => {
        const day = payments.find(payment => payment.date === date);
        if (day) {
            return day.expenses;
        }

        return [];
    };

    render() {
        return (
            <div className='Details'>
                <h2 className='Details-heading'>{this.props.selectedDate}</h2>
                <DetailList
                    payments={this.filterPayments(this.props.selectedDate)}
                />
            </div>
        );
    }
}

export default Details;
