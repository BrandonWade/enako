import React, { useContext } from 'react';
import { Calendar as BigCalendar, dateFnsLocalizer } from 'react-big-calendar';
import format from 'date-fns/format';
import parse from 'date-fns/parse';
import startOfWeek from 'date-fns/startOfWeek';
import getDay from 'date-fns/getDay';
import style from 'react-big-calendar/lib/css/react-big-calendar.css';
import enUS from 'date-fns/locale/en-US';
import ExpenseContext from '../../../contexts/ExpenseContext';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import CalendarDate from '../CalendarDate';
import './Calendar.scss';

const localizer = dateFnsLocalizer({
    format,
    parse,
    startOfWeek,
    getDay,
    locales: {
        'en-US': enUS,
    },
});

const Calendar = ({ setSelectedDate }) => {
    const expenses = useContext(ExpenseContext);
    const selectedDate = useContext(SelectedDateContext);

    return (
        <div className='Calendar'>
            <BigCalendar
                localizer={localizer}
                events={[]}
                views={['month']}
                style={style}
                selectable='ignoreEvents'
                components={{
                    dateCellWrapper: CalendarDate({ expenses, selectedDate }),
                }}
                onSelectSlot={e => setSelectedDate(e.start)}
            />
        </div>
    );
};

export default Calendar;
