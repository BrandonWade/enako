import React, { Component } from 'react';
import Calendar from './components/Calendar';
import Details from './components/Details';

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: 'September 29', 
        };
    }

    render() {
        return (
            <div className='App'>
                <div className='App-content'>
                <Calendar />
                    <Details
                        selectedDate={this.state.selectedDate}
                    />
                </div>
            </div>
        );
    }
}

export default App;
