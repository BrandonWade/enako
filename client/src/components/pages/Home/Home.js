import React, { useState, useEffect } from 'react';
import Calendar from '../../Calendar';
import Details from '../../Details';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import './Home.css';

const Home = () => {
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [types, setTypes] = useState([]);
    const [categories, setCategories] = useState([]);
    const [expenses, setExpenses] = useState([]);

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
        <div className='Home'>
            <div className='Home-content'>
                <Calendar expenses={expenses} setSelectedDate={setSelectedDate} />
                <Details selectedDate={selectedDate} types={types} categories={categories} expenses={expenses} />
            </div>
        </div>
    );
};

export default Home;
