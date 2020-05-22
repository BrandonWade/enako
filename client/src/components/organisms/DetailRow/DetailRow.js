import React from 'react';
import { Link } from 'react-router-dom';
import Category from '../../atoms/Category';
import './DetailRow.css';

const DetailRow = ({ expense, selectedDate }) => {
    return (
        <tr>
            <td>
                <Link
                    to={{
                        pathname: `/expenses/${expense.id}`,
                        state: {
                            selectedDate: selectedDate,
                        },
                    }}
                >
                    <button>Edit</button>
                </Link>
            </td>
            <td>
                <Category type={expense.expense_category} />
            </td>
            <td className='detail-row__description'>{expense.expense_description}</td>
            <td className='detail-row__amount'>${expense.expense_amount.toFixed(2)}</td>
        </tr>
    );
};

export default DetailRow;
