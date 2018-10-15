import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import DetailListItem from './DetailListItem';
import RoundButton from './RoundButton';
import '../css/DetailList.css';

class DetailList extends Component {
    getTotal() {
        let total = 0.0;

        this.props.expenses.forEach((expense) => {
            total += expense.amount;
        });

        return total;
    }

    renderPaymentsSection = () => {
        return this.props.expenses.length > 0 ? (
            <>
                <h4 className='DetailList-sectionHeading'>Expenses</h4>
                <ul className='DetailList'>
                    {
                        this.props.expenses.map(expense => {
                            return (
                                <DetailListItem
                                    key={expense.id}
                                    name={expense.description}
                                    amount={expense.amount}
                                    colour={expense.colour}
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
                    <Link to='/create'>
                        <RoundButton text='+' />
                    </Link>
                </div>
                { this.renderTotalsSection() }
            </div>
        );
    }
}

export default DetailList;
