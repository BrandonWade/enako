import React from 'react';
import { format } from 'date-fns';
import './CalendarDate.scss';

const calculateTotal = (date, expenses) => {
    const compareDate = format(date, 'yyy-MM-dd');
    const total = expenses.reduce((total, expense) => (expense.expense_date === compareDate ? total + expense.amount : total), 0);

    return total.toFixed(2);
};

const CalendarDate = ({ expenses, selectedDate }) => props => {
    const total = calculateTotal(props.value, expenses);
    const selected = props.value.getTime() === selectedDate.getTime() ? 'calendar-date--selected' : '';
    const negative = total > 0 ? 'u-negative' : '';
    const className = `${props.children.props.className} calendar-date ${negative} ${selected}`;

    return <div className={className}>{total > 0 ? `$${total}` : ''}</div>;
};

export default CalendarDate;
