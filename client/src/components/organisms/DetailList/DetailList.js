import React from 'react';
import { Link } from 'react-router-dom';
import DetailRow from '../DetailRow';
import Button from '../../atoms/Button';
import './DetailList.scss';

const DetailList = ({ expenses }) => {
    const total = expenses.reduce((total, expense) => (total += expense.amount), 0);

    const renderPaymentsSection = () => {
        return expenses.length > 0 ? (
            <>
                <h4 className='DetailList-sectionHeading'>Expenses</h4>
                <table>
                    <tbody>
                        {expenses.map(e => (
                            <DetailRow key={e.id} expense={e} />
                        ))}
                    </tbody>
                </table>
            </>
        ) : (
            <p className='DetailList-message'>No expenses to display.</p>
        );
    };

    const renderTotalsSection = () => {
        return total > 0 ? (
            <div className='DetailList-totalSection'>
                <div>Total</div>
                <div>${total.toFixed(2)}</div>
            </div>
        ) : null;
    };

    return (
        <div>
            {renderPaymentsSection()}
            {renderTotalsSection()}
            <div className='DetailList-addContainer'>
                <Link to='/expenses'>
                    <Button color='orange' text='Add' />
                </Link>
            </div>
        </div>
    );
};

export default DetailList;
