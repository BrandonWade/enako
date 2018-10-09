import React, { Component } from 'react';
import Calendar from '../components/Calendar';
import Details from '../components/Details';
import moment from 'moment';
import '../css/Home.css';

class Home extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: moment().format('MMMM Do YYYY'),
        };
    }

    setSelectedDate = (date) => {
        this.setState({
            selectedDate: moment(date).format('MMMM Do YYYY'),
        });
    };

    render() {
        return (
            <div className='Home'>
                <div className='Home-content'>
                    <Calendar setSelectedDate={this.setSelectedDate} />
                    <Details selectedDate={this.state.selectedDate} />
                </div>
            </div>
        );
    }
}

export default Home;
