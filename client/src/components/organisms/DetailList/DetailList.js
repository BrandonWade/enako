import React from 'react';
import { Link } from 'react-router-dom';
import DetailRow from '../DetailRow';
import Button from '../../atoms/Button';
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
                <h4 className='detail-list__section-heading'>Expenses</h4>
                <table>
                    <tbody>
                        {props.expenses.map(expense => {
                            return <DetailRow key={expense.id} selectedDate={props.selectedDate} expense={expense} />;
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
            <div className='detail-list__total-section'>
                <div>Total</div>
                <div>${total}</div>
            </div>
        ) : null;
    };

    return (
        <div>
            {renderPaymentsSection()}
            {renderTotalsSection()}
            <div className='detail-list__add-container'>
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
                    <Button primary text='Add' />
                </Link>
            </div>
        </div>
    );
};

export default DetailList;
