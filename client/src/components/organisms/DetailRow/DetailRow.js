import React from 'react';
import { Link } from 'react-router-dom';
import Category from '../../atoms/Category';
import './DetailRow.scss';

const DetailRow = ({ expense }) => {
    return (
        <tr>
            <td>
                <Link to={`/expenses/${expense.id}`}>
                    <button>Edit</button>
                </Link>
            </td>
            <td>
                <Category type={expense.category} />
            </td>
            <td className='detail-row__description'>{expense.description}</td>
            <td className='detail-row__amount'>${expense.amount.toFixed(2)}</td>
        </tr>
    );
};

export default DetailRow;
