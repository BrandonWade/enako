import React, { Component } from 'react';
import Calendar from '../Calendar';
import Details from '../Details';
import fetchBootInfo from '../../effects/fetchBootInfo';
import '../../css/Home.css';

class Home extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: new Date(),
            types: [],
            categories: [],
            expenses: [],
        };
    }

    componentDidMount = async () => {
        const bootInfo = await fetchBootInfo();

        this.setState({
            ...bootInfo,
        });
    }

    setSelectedDate = (selectedDate) => {
        this.setState({
            selectedDate,
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
                        expenses={this.state.expenses}
                    />
                </div>
            </div>
        );
    }
}

export default Home;
