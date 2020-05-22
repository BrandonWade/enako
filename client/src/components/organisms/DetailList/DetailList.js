import React from 'react';
import { Link } from 'react-router-dom';
import DetailRow from '../DetailRow';
import Button from '../../atoms/Button';
import './DetailList.css';

const DetailList = ({ expenses, selectedDate }) => {
    const total = expenses.reduce((total, expense) => (total += expense.expense_amount), 0);

    const renderPaymentsSection = () => {
        return expenses.length > 0 ? (
            <>
                <h4 className='detail-list__section-heading'>Expenses</h4>
                <table>
                    <tbody>
                        {expenses.map(expense => {
                            return <DetailRow key={expense.id} selectedDate={selectedDate} expense={expense} />;
                        })}
                    </tbody>
                </table>
            </>
        ) : (
            <p className='detail-list__message'>No expenses to display.</p>
        );
    };

    const renderTotalsSection = () => {
        return total > 0 ? (
            <div className='detail-list__total-section'>
                <div>Total</div>
                <div>${total.toFixed(2)}</div>
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
                            selectedDate: selectedDate,
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
