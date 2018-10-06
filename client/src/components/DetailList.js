import React, { Component } from 'react';
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

    render() {
        return (
            <div>
                <h3 className='DetailList-sectionHeading'>Payments</h3>
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
                <h3 className='DetailList-sectionHeading'>Totals</h3>
                <ul className='DetailList'>
                    <DetailListItem
                        name='Total'
                        amount={this.getTotal()}
                    />
                </ul>
            </div>
        );
    }
}

export default DetailList;
