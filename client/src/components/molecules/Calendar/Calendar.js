import React from 'react';
import { Calendar as BigCalendar, momentLocalizer } from 'react-big-calendar';
import moment from 'moment';
import style from 'react-big-calendar/lib/css/react-big-calendar.css';
import CalendarDate from '../CalendarDate';
import './Calendar.css';

const localizer = momentLocalizer(moment);
const views = ['month'];

const Calendar = ({ expenses, setSelectedDate }) => {
    const onDateSelected = evt => {
        setSelectedDate(evt.start);
    };

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
                onSelectSlot={onDateSelected}
            />
        </div>
    );
};

export default Calendar;
