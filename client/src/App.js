import React, { Component } from 'react';
import Calendar from './components/Calendar';
import Details from './components/Details';
import moment from 'moment';

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: moment().format('MMMM Do YYYY'),
        };
    }

    setSelectedDate = (selectedDate) => {
        this.setState({
            selectedDate,
        });
    };

    render() {
        return (
            <div className='App'>
                <div className='App-content'>
                    <Calendar
                        setSelectedDate={this.setSelectedDate}
                    />
                    <Details
                        selectedDate={this.state.selectedDate}
                    />
                </div>
            </div>
        );
    }
}

export default App;
