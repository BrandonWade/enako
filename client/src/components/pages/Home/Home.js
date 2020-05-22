import React, { useState, useEffect } from 'react';
import Calendar from '../../molecules/Calendar';
import Details from '../../organisms/Details';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import './Home.css';

const Home = () => {
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [types, setTypes] = useState();
    const [categories, setCategories] = useState();
    const [expenses, setExpenses] = useState();

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
                <Calendar setSelectedDate={setSelectedDate} />
                <Details selectedDate={selectedDate} />
            </div>
        </div>
    );
};

export default Home;
