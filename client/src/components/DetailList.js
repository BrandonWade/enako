import React, { Component } from 'react';
import DetailListItem from './DetailListItem';
import '../css/DetailList.css';

class DetailList extends Component {
    filterPayments(type) {
        return this.props.payments.filter(payment => payment.type === type);
    }

    getPaymentTotals() {
        const totals = {
            'one-time': 0.0,
            'recurring': 0.0,
            'total': 0.0,
        };

        this.props.payments.forEach((payment) => {
            totals[payment.type] += payment.amount;
            totals['total'] += payment.amount;
        });

        return totals;
    }

    render() {
        return (
            <div>
                <h3 className='DetailList-sectionHeading'>One-Time Payments</h3>
                <ul className='DetailList'>
                    {
                        this.filterPayments('one-time').map(payment => {
                            return (
                                <DetailListItem
                                    key={payment.id}
                                    name={payment.description}
                                    amount={payment.amount}
                                    colour={payment.colour}
                                />
                            );
                        })
                    }
                </ul>
                <h3 className='DetailList-sectionHeading'>Recurring Payments</h3>
                <ul className='DetailList'>
                    {
                        this.filterPayments('recurring').map(payment => {
                            return (
                                <DetailListItem
                                    key={payment.id}
                                    name={payment.description}
                                    amount={payment.amount}
                                    colour={payment.colour}
                                />
                            );
                        })
                    }
                </ul>
                <h3 className='DetailList-sectionHeading'>Totals</h3>
                <ul className='DetailList'>
                    {
                        Object.entries(this.getPaymentTotals()).map(total => {
                            return (
                                <DetailListItem
                                    key={total[0]}
                                    name={total[0]}
                                    amount={total[1]}
                                />
                            );
                        })
                    }
                </ul>
            </div>
        );
    }
}

export default DetailList;
