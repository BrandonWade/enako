import React from 'react';
import { Link } from 'react-router-dom';
import DetailRow from '../DetailRow';
import RoundButton from '../../atoms/RoundButton';
import './DetailList.css';

const DetailList = props => {
    const getTotal = () => {
        let total = 0.0;

        props.expenses.forEach(expense => {
            total += expense.expense_amount;
        });

        return total;
    };

    const renderPaymentsSection = () => {
        return props.expenses.length > 0 ? (
            <>
                <h4 className='DetailList-sectionHeading'>Expenses</h4>
                <table>
                    <tbody>
                        {props.expenses.map(expense => {
                            return (
                                <DetailRow
                                    key={expense.id}
                                    selectedDate={props.selectedDate}
                                    types={props.types}
                                    categories={props.categories}
                                    expense={expense}
                                />
                            );
                        })}
                    </tbody>
                </table>
            </>
        ) : (
            'No expenses to display.'
        );
    };

    const renderTotalsSection = () => {
        const total = getTotal().toFixed(2);

        return total > 0 ? (
            <div className='DetailList-totalSection'>
                <div className='DetailList-totalText'>Total</div>
                <div className='DetailList-totalAmount'>${total}</div>
            </div>
        ) : null;
    };

    return (
        <div>
            {renderPaymentsSection()}
            <div className='DetailList-addItemContainer'>
                <Link
                    to={{
                        pathname: '/expenses',
                        state: {
                            selectedDate: props.selectedDate,
                            types: props.types,
                            categories: props.categories,
                            type: '',
                            category: '',
                            description: '',
                            amount: 0.0,
                        },
                    }}
                >
                    <RoundButton text='+' />
                </Link>
            </div>
            {renderTotalsSection()}
        </div>
    );
};

export default DetailList;
