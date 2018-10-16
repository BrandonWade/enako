import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import '../css/DetailRow.css';

class DetailRow extends Component {
    render() {
        return (
            <tr className='DetailRow'>
                <td className='DetailRow-edit'>
                    <Link
                        to={{
                            pathname: `/expenses/${this.props.expense.id}`,
                            state: {
                                selectedDate: this.props.selectedDate,
                                ...this.props.expense,
                            },
                        }}
                    >
                        <button>
                            Edit
                        </button>
                    </Link>
                </td>
                <td className='DetailRow-category'>
                    [category]
                </td>
                <td className='DetailRow-description'>
                    {this.props.expense.description}
                </td>
                <td className='DetailRow-amount'>
                    {this.props.expense.amount.toFixed(2)}
                </td>
            </tr>
        );
    }
}

export default DetailRow;
