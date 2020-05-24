import React from 'react';
import { format } from 'date-fns';
import './CalendarDate.scss';

const calculateTotal = (date, expenses) => {
    const compareDate = format(date, 'yyy-MM-dd');
    const total = expenses.reduce((total, expense) => (expense.expense_date === compareDate ? total + expense.amount : total), 0);

    return total.toFixed(2);
};

const CalendarDate = ({ expenses }) => props => {
    const total = calculateTotal(props.value, expenses);
    const className = `${props.children.props.className} calendar-date ${total > 0 ? 'u-negative' : 'u-positive'}`;

    return <div className={className}>{`$${total}`}</div>;
};

export default CalendarDate;
