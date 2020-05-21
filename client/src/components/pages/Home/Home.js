import React, { useState, useEffect } from 'react';
import Calendar from '../../molecules/Calendar';
import Details from '../../organisms/Details';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import './Home.css';

const Home = () => {
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [types, setTypes] = useState([
        {
            id: 1,
            type_name: 'general',
        },
        {
            id: 2,
            type_name: 'unnecessary',
        },
        {
            id: 3,
            type_name: 'recurring',
        },
    ]);
    const [categories, setCategories] = useState([
        {
            id: 1,
            category_name: 'food',
        },
        {
            id: 2,
            category_name: 'entertainment',
        },
        {
            id: 3,
            category_name: 'transportation',
        },
        {
            id: 4,
            category_name: 'clothing',
        },
        {
            id: 5,
            category_name: 'technology',
        },
        {
            id: 6,
            category_name: 'health',
        },
    ]);
    const [expenses, setExpenses] = useState([
        {
            id: 1,
            expense_type: 'unnecessary',
            expense_category: 'food',
            expense_description: 'foo',
            expense_amount: 25.0,
            expense_date: '2020-05-20',
        },
        {
            id: 2,
            expense_type: 'unnecessary',
            expense_category: 'technology',
            expense_description: 'bar',
            expense_amount: 100.0,
            expense_date: '2020-05-20',
        },
        {
            id: 3,
            expense_type: 'unnecessary',
            expense_category: 'clothing',
            expense_description: 'baz',
            expense_amount: 75.0,
            expense_date: '2020-05-20',
        },
    ]);

    useEffect(() => {
        const boot = async () => {
            const bootInfo = await fetchBootInfo();

            setTypes(bootInfo.types);
            setCategories(bootInfo.Categories);
            setExpenses(bootInfo.expenses);
        };
        boot();
    }, []);

    return (
        <div className='home'>
            <div className='home__content'>
                <Calendar expenses={expenses} setSelectedDate={setSelectedDate} />
                <Details selectedDate={selectedDate} types={types} categories={categories} expenses={expenses} />
            </div>
        </div>
    );
};

export default Home;
