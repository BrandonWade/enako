import React, { Component } from 'react';

class DetailList extends Component {
    render() {
        return (
            <ul className='Details-list'>
                {
                    this.props.payments.map(payment => {
                        return (
                            <li key={payment.id}>
                                {payment.description}
                            </li>
                        );
                    })
                }
            </ul>
        );
    }
}

export default DetailList;
