import React, { Component } from 'react';
import { Calendar as BigCalendar, momentLocalizer } from 'react-big-calendar';
import moment from 'moment';
import style from 'react-big-calendar/lib/css/react-big-calendar.css';
import CalendarDate from './CalendarDate';
import '../css/Calendar.css';

const localizer = momentLocalizer(moment);
const views = ['month'];

class Calendar extends Component {
    onDateSelected = (evt) => {
        this.props.setSelectedDate(evt.start);
    };

    render() {
        const { expenses } = this.props;

        return (
            <div className='Calendar'>
                <BigCalendar
                    localizer={localizer}
                    events={[]}
                    views={views}
                    style={style}
                    selectable='ignoreEvents'
                    components={{
                        dateCellWrapper: CalendarDate({ expenses }),
                    }}
                    onSelectSlot={this.onDateSelected}
                />
            </div>
        );
    }
}

export default Calendar;
