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
                <Category id={expense.category_id} />
            </td>
            <td className='DetailRow-description'>{expense.description}</td>
            <td className='DetailRow-amount'>${expense.amount.toFixed(2)}</td>
        </tr>
    );
};

export default DetailRow;
