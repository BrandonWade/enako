import React, { Component } from 'react';
import DetailList from './DetailList';
import Card from './Card';
import expenses from '../data/SampleExpenses';
import '../css/Details.css';

class Details extends Component {
    filterExpenses = (date) => {
        const day = expenses.find(expense => expense.date === date);
        if (day) {
            return day.expenses;
        }

        return [];
    };

    render() {
        return (
            <div className='Details'>
                <Card heading={this.props.selectedDate}>
                    <DetailList
                        selectedDate={this.props.selectedDate}
                        types={this.props.types}
                        categories={this.props.categories}
                        expenses={this.filterExpenses(this.props.selectedDate)}
                    />
                </Card>
            </div>
        );
    }
}

export default Details;
