import React, { Component } from 'react';
import BigCalendar from 'react-big-calendar';
import moment from 'moment';
import style from 'react-big-calendar/lib/css/react-big-calendar.css';
import '../css/Calendar.css';

const localizer = BigCalendar.momentLocalizer(moment);
const views = ['month'];

class Calendar extends Component {
  render() {
    return (
      <div className='Calendar'>
        <BigCalendar
            localizer={localizer}
            events={[]}
            views={views}
            style={style}
        />
      </div>
    );
  }
}

export default Calendar;
