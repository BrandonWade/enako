import React, { Component } from 'react';
import Calendar from '../Calendar';
import Details from '../Details';
import fetchBootInfo from '../../effects/fetchBootInfo';
import moment from 'moment';
import '../../css/Home.css';

class Home extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: moment().format('MMMM Do YYYY'),
            types: [],
            categories: [],
        };
    }

    componentDidMount = async () => {
        const bootInfo = await fetchBootInfo();

        this.setState({
            ...bootInfo,
        });
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
                    <Details
                        selectedDate={this.state.selectedDate}
                        types={this.state.types}
                        categories={this.state.categories}
                    />
                </div>
            </div>
        );
    }
}

export default Home;
