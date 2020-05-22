import React from 'react';
import { Link } from 'react-router-dom';
import Category from '../../atoms/Category';
import './DetailRow.css';

const DetailRow = props => {
    return (
        <tr>
            <td>
                <Link
                    to={{
                        pathname: `/expenses/${props.expense.id}`,
                        state: {
                            selectedDate: props.selectedDate,
                            ...props.expense,
                        },
                    }}
                >
                    <button>Edit</button>
                </Link>
            </td>
            <td>
                <Category type={props.expense.expense_category} />
            </td>
            <td className='detail-row__description'>{props.expense.expense_description}</td>
            <td className='detail-row__amount'>${props.expense.expense_amount.toFixed(2)}</td>
        </tr>
    );
};

export default DetailRow;
