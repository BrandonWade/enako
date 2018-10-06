import React, { Component } from 'react';
import moment from 'moment';
import DetailList from './DetailList';
import payments from './SamplePayments';
import '../css/Details.css';

class Details extends Component {
    formatDate = (date) => {
        return moment(date).format('MMMM Do YYYY');
    };

    filterPayments = (date) => {
        const day = payments.find(payment => payment.date === date);
        return day.events;
    };

    render() {
        return (
            <div className='Details'>
                <h1 className='Details-heading'>Breakdown for {this.formatDate(this.props.selectedDate)}</h1>
                <DetailList
                    payments={this.filterPayments(this.props.selectedDate)}
                />
            </div>
        );
    }
}

export default Details;
