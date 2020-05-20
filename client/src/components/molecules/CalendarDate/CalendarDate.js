import React from 'react';
import moment from 'moment';
import './CalendarDate.css';

const calculateTotal = (date, expenses) => {
    const compareDate = moment(date).format('YYYY-MM-DD');
    const total = expenses.reduce((total, expense) => {
        if (expense.expense_date === compareDate) {
            return total + expense.expense_amount;
        }

        return total;
    }, 0);

    return (total || 0).toFixed(2);
};

const CalendarDate = ({ expenses }) => (props) => {
    const total = calculateTotal(props.value, expenses);
    const className = `${props.children.props.className} CalendarDate ${total > 0 ? 'u-negative' : 'u-positive'}`;

    return <div className={className}>{`$${total}`}</div>;
};

export default CalendarDate;
