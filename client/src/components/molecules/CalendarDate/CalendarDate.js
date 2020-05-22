import React from 'react';
import moment from 'moment';
import './CalendarDate.css';

const calculateTotal = (date, expenses) => {
    const compareDate = moment(date).format('YYYY-MM-DD');
    const total = expenses.reduce((total, expense) => (expense.expense_date === compareDate ? total + expense.expense_amount : total), 0);

    return total.toFixed(2);
};

const CalendarDate = ({ expenses }) => props => {
    const total = calculateTotal(props.value, expenses);
    const className = `${props.children.props.className} calendar-date ${total > 0 ? 'u-negative' : 'u-positive'}`;

    return <div className={className}>{`$${total}`}</div>;
};

export default CalendarDate;
