import React, { useState, useContext } from 'react';
import { format } from 'date-fns';
import { Link } from 'react-router-dom';
import createExpense from '../../../effects/createExpense';
import updateExpense from '../../../effects/updateExpense';
import deleteExpense from '../../../effects/deleteExpense';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import CategoryContext from '../../../contexts/CategoryContext';
import ExpenseContext from '../../../contexts/ExpenseContext';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Card from '../../atoms/Card';
import Button from '../../atoms/Button';
import InputField from '../../molecules/InputField';
import SelectField from '../../molecules/SelectField';
import './Editor.scss';

const Editor = props => {
    const selectedDate = useContext(SelectedDateContext);
    const categories = useContext(CategoryContext);
    const expenses = useContext(ExpenseContext);
    const expenseID = parseInt(props.computedMatch.params.id);
    const expense = expenses.find(e => e.id === expenseID) || {};

    const [category, setCategory] = useState(expense.category || '');
    const [description, setDescription] = useState(expense.description || '');
    const [amount, setAmount] = useState(expense.amount || 0);

    const notFoundRedirect = () => {
        return expenseID && !expense.id ? <AuthenticatedRedirect /> : null;
    };

    const renderHeadingText = () => {
        const formattedDate = format(selectedDate, 'MMMM do yyyy');
        return expenseID ? `Editing an expense on ${formattedDate}` : `Creating a new expense on ${formattedDate}`;
    };

    const renderDeleteButton = () => {
        return expenseID ? <Button text='Delete' className='editor__delete button--red' onClick={onExpenseDelete} /> : null;
    };

    const renderSubmitButtonText = () => {
        return expenseID ? 'Save' : 'Create';
    };

    const onExpenseDelete = () => {
        deleteExpense(expenseID);
    };

    const onExpenseSubmit = () => {
        const id = expenseID || 0;
        const data = {
            category,
            description,
            amount,
        };

        if (id) {
            updateExpense(id, data);
        } else {
            createExpense(data);
        }
    };

    return (
        <>
            {notFoundRedirect()}
            <div className='editor'>
                <div className='editor__content'>
                    <Card heading={renderHeadingText()}>
                        <SelectField
                            name='category'
                            label='Category'
                            value={category}
                            description='Choose the most relevant category of expense'
                            onChange={e => setCategory(e.target.value)}
                        >
                            <option value=''>-- Select a Category -- </option>
                            {categories.map(c => {
                                return (
                                    <option key={c.id} value={c.name}>
                                        {c.name}
                                    </option>
                                );
                            })}
                        </SelectField>
                        <InputField
                            name='description'
                            label='Description'
                            value={description}
                            description='Give a brief description of this expense'
                            onChange={e => setDescription(e.target.value)}
                        />
                        <InputField
                            name='amount'
                            label='Amount'
                            value={amount}
                            description='Enter the cost of this expense'
                            onChange={e => setAmount(e.target.value)}
                        />
                        <div className='editor__buttons'>
                            <Link to='/'>
                                <Button text='Cancel' />
                            </Link>
                            <div>
                                {renderDeleteButton()}
                                <Button primary text={renderSubmitButtonText()} onClick={() => onExpenseSubmit()} />
                            </div>
                        </div>
                    </Card>
                </div>
            </div>
        </>
    );
};

export default Editor;
