import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import DetailRow from './DetailRow';
import RoundButton from './RoundButton';
import '../css/DetailList.css';

class DetailList extends Component {
    getTotal() {
        let total = 0.0;

        this.props.expenses.forEach((expense) => {
            total += expense.expense_amount;
        });

        return total;
    }

    renderPaymentsSection = () => {
        return this.props.expenses.length > 0 ? (
            <>
                <h4 className='DetailList-sectionHeading'>Expenses</h4>
                <table>
                    <tbody>
                        {
                            this.props.expenses.map(expense => {
                                return (
                                    <DetailRow
                                        key={expense.id}
                                        selectedDate={this.props.selectedDate}
                                        types={this.props.types}
                                        categories={this.props.categories}
                                        expense={expense}
                                    />
                                );
                            })
                        }
                    </tbody>
                </table>
            </>
        ) : (
            'No expenses to display.'
        );
    };

    renderTotalsSection = () => {
        const total = this.getTotal().toFixed(2);

        return total > 0 ? (
            <div className='DetailList-totalSection'>
                <div className='DetailList-totalText'>Total</div>
                <div className='DetailList-totalAmount'>${total}</div>
            </div>
        ) : null;
    };

    render() {
        return (
            <div>
                { this.renderPaymentsSection() }
                <div className='DetailList-addItemContainer'>
                    <Link
                        to={{
                            pathname: '/expenses',
                            state: {
                                selectedDate: this.props.selectedDate,
                                types: this.props.types,
                                categories: this.props.categories,
                                type: '',
                                category: '',
                                description: '',
                                amount: 0.00,
                            },
                        }}
                    >
                        <RoundButton text='+' />
                    </Link>
                </div>
                { this.renderTotalsSection() }
            </div>
        );
    }
}

export default DetailList;
