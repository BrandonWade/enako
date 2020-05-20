import React from 'react';
import moment from 'moment';
import DetailList from '../DetailList';
import Card from '../Card';
import './Details.css';

const Details = (props) => {
    const filterExpenses = (date) => {
        const compareDate = moment(date).format('YYYY-MM-DD');
        const expenses = props.expenses.filter((expense) => expense.expense_date === compareDate);
        if (expenses.length) {
            return expenses;
        }

        return [];
    };

    return (
        <div className='Details'>
            <Card heading={moment(props.selectedDate).format('MMMM Do YYYY')}>
                <DetailList
                    selectedDate={props.selectedDate}
                    types={props.types}
                    categories={props.categories}
                    expenses={filterExpenses(props.selectedDate)}
                />
            </Card>
        </div>
    );
};

export default Details;
