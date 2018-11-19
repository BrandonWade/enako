import React, { Component } from 'react';
import '../css/CalendarDate.css';

class CalendarDate extends Component {
    render() {
        const total = 19.99; // TODO: Hardcoded for testing
        const className = `${this.props.children.props.className} CalendarDate ${total > 0 ? 'u-negative' : 'u-positive'}`

        return (
            <div
                className={className}
            >
                $19.99
            </div>
        );
    }
}

export default CalendarDate;
