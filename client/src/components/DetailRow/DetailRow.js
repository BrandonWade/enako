import React from 'react';
import { Link } from 'react-router-dom';
import Category from '../Category';
import './DetailRow.css';

const DetailRow = (props) => {
    return (
        <tr className='DetailRow'>
            <td className='DetailRow-edit'>
                <Link
                    to={{
                        pathname: `/expenses/${props.expense.id}`,
                        state: {
                            selectedDate: props.selectedDate,
                            types: props.types,
                            categories: props.categories,
                            ...props.expense,
                        },
                    }}
                >
                    <button>Edit</button>
                </Link>
            </td>
            <td className='DetailRow-category'>
                <Category type={props.expense.expense_category} />
            </td>
            <td className='DetailRow-description'>{props.expense.expense_description}</td>
            <td className='DetailRow-amount'>${props.expense.expense_amount.toFixed(2)}</td>
        </tr>
    );
};

export default DetailRow;
