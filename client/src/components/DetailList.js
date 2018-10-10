import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import DetailListItem from './DetailListItem';
import RoundButton from './RoundButton';
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
        const total = this.getTotal();

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
        ) : (
            'No expenses to display.'
        );
    };

    renderTotalsSection = () => {
        const total = this.getTotal();

        return total > 0 ? (
            <div className='DetailList-totals'>
                <h4 className='DetailList-sectionHeading'>Totals</h4>
                <ul className='DetailList'>
                    <DetailListItem
                        name='Total'
                        amount={this.getTotal()}
                    />
                </ul>
            </div>
        ) : null;
    };

    render() {
        return (
            <div>
                { this.renderPaymentsSection() }
                <div className='DetailList-addItemContainer'>
                    <Link to='/edit'>
                        <RoundButton text='+' />
                    </Link>
                </div>
                { this.renderTotalsSection() }
            </div>
        );
    }
}

export default DetailList;
