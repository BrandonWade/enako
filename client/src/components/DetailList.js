import React, { Component, Fragment } from 'react';
import DetailListItem from './DetailListItem';
import '../css/DetailList.css';

class DetailList extends Component {
    getTotal() {
        let total = 0.0;

        this.props.payments.forEach((payment) => {
            total += payment.amount;
        });

        return total;
    }

    renderPaymentsSection = () => {
        return this.props.payments.length > 0 ? (
            <Fragment>
                <h4 className='DetailList-sectionHeading'>Payments</h4>
                <ul className='DetailList'>
                    {
                        this.props.payments.map(payment => {
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
            </Fragment>
        ) : null;
    };

    renderTotalsSection = () =>  {
        const total = this.getTotal();

        return total > 0 ? (
            <Fragment>
                <h4 className='DetailList-sectionHeading'>Totals</h4>
                <ul className='DetailList'>
                    <DetailListItem
                        name='Total'
                        amount={total}
                    />
                </ul>
            </Fragment>
        ) : (
            'No expenses to display.'
        );
    };

    render() {
        return (
            <div>
                { this.renderPaymentsSection() }
                { this.renderTotalsSection() }
            </div>
        );
    }
}

export default DetailList;
