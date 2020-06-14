import React from 'react';
import Calendar from '../../molecules/Calendar';
import Details from '../../organisms/Details';
import './Home.scss';

const Home = props => {
    return (
        <div className='Home'>
            <div className='Home-content'>
                <Calendar setSelectedDate={props.setSelectedDate} />
                <Details selectedDate={props.selectedDate} />
            </div>
        </div>
    );
};

export default Home;
