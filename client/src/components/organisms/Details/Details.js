import React, { useContext } from 'react';
import { format } from 'date-fns';
import ExpenseContext from '../../../contexts/ExpenseContext';
import DetailList from '../DetailList';
import Card from '../../atoms/Card';
import './Details.scss';

const Details = props => {
    const expenses = useContext(ExpenseContext);

    const filterExpenses = (expenses, date) => {
        const compareDate = format(date, 'yyyy-MM-dd');
        const dailyExpenses = expenses.filter(expense => expense.expense_date === compareDate);
        if (dailyExpenses.length) {
            return dailyExpenses;
        }

        return [];
    };

    return (
        <Card heading={format(props.selectedDate, 'MMMM do yyyy')}>
            <DetailList expenses={filterExpenses(expenses, props.selectedDate)} />
        </Card>
    );
};

export default Details;
