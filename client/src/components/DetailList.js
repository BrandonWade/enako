import React, { Component } from 'react';
import { Link } from 'react-router-dom';
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
            <>
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
            </>
        ) : null;
    };

    renderTotalsSection = () =>  {
        const total = this.getTotal();

        return total > 0 ? (
            <>
                <h4 className='DetailList-sectionHeading'>Totals</h4>
                <ul className='DetailList'>
                    <DetailListItem
                        name='Total'
                        amount={total}
                    />
                </ul>
            </>
        ) : (
            'No expenses to display.'
        );
    };

    render() {
        return (
            <div>
                { this.renderPaymentsSection() }
                <Link to='/edit'>
                    <button>+</button>
                </Link>
                { this.renderTotalsSection() }
            </div>
        );
    }
}

export default DetailList;
