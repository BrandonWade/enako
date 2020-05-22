import React, { useContext } from 'react';
import moment from 'moment';
import ExpenseContext from '../../../contexts/ExpenseContext';
import DetailList from '../DetailList';
import Card from '../../atoms/Card';
import './Details.css';

const Details = props => {
    const expenses = useContext(ExpenseContext);

    const filterExpenses = (expenses, date) => {
        const compareDate = moment(date).format('YYYY-MM-DD');
        const dailyExpenses = expenses.filter(expense => expense.expense_date === compareDate);
        if (dailyExpenses.length) {
            return dailyExpenses;
        }

        return [];
    };

    return (
        <Card heading={moment(props.selectedDate).format('MMMM Do YYYY')}>
            <DetailList expenses={filterExpenses(expenses, props.selectedDate)} />
        </Card>
    );
};

export default Details;
