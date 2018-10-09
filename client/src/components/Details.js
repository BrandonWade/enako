import React, { Component } from 'react';
import DetailList from './DetailList';
import Card from './Card';
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
                <Card heading={this.props.selectedDate}>
                    <DetailList
                        payments={this.filterPayments(this.props.selectedDate)}
                    />
                </Card>
            </div>
        );
    }
}

export default Details;
