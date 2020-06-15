import React, { useState, useContext } from 'react';
import { format } from 'date-fns';
import { Link, useHistory } from 'react-router-dom';
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
    const history = useHistory();
    const categories = useContext(CategoryContext);
    const expenses = useContext(ExpenseContext);
    const expenseID = parseInt(props.computedMatch.params.id);
    const expense = expenses.find(e => e.id === expenseID) || {};

    const [expenseDate, setExpenseDate] = useState(useContext(SelectedDateContext));
    const [categoryID, setCategoryID] = useState(expense.category_id || 0);
    const [description, setDescription] = useState(expense.description || '');
    const [amount, setAmount] = useState(expense.amount || 0);
    const formattedDate = format(expenseDate, 'yyyy-MM-dd');

    const notFoundRedirect = () => {
        return expenseID && !expense.id ? <AuthenticatedRedirect /> : null;
    };

    const renderHeadingText = () => {
        return expenseID ? 'Editing Expense' : 'New Expense';
    };

    const renderDeleteButton = () => {
        return expenseID ? (
            <Link to='/' onClick={onExpenseDelete}>
                <Button color='red' text='Delete' className='Editor-delete' />
            </Link>
        ) : null;
    };

    const renderSubmitButtonText = () => {
        return expenseID ? 'Save' : 'Create';
    };

    const onExpenseDelete = () => {
        deleteExpense(expenseID);
        props.setExpenses(expenses.filter(e => e.id !== expenseID));
    };

    const onExpenseSubmit = async () => {
        if (categoryID === 0 || description === '' || amount <= 0) {
            return;
        }

        props.setSelectedDate(expenseDate);

        const id = expenseID || 0;
        const data = {
            category_id: parseInt(categoryID),
            description,
            amount: amount * 100,
            expense_date: formattedDate,
        };

        if (id) {
            const index = expenses.findIndex(e => e.id === id);
            const expense = await updateExpense(id, data);
            props.setExpenses([...expenses.slice(0, index), expense, ...expenses.slice(index + 1)]);
        } else {
            const expense = await createExpense(data);
            props.setExpenses([...expenses, expense]);
        }

        history.push('/');
    };

    return (
        <>
            {notFoundRedirect()}
            <div className='Editor'>
                <div className='Editor-content'>
                    <Card heading={renderHeadingText()}>
                        <InputField
                            type='date'
                            label='Date'
                            value={formattedDate}
                            description='Select the date that the expense occurred'
                            onChange={e => setExpenseDate(new Date(`${e.target.value} 00:00:00`))}
                        />
                        <SelectField
                            name='category'
                            label='Category'
                            value={categoryID}
                            description='Choose the most relevant category of expense'
                            onChange={e => setCategoryID(e.target.value)}
                        >
                            <option value=''>-- Select a Category -- </option>
                            {categories.map(c => {
                                return (
                                    <option key={c.id} value={c.id}>
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
                            type='number'
                            name='amount'
                            label='Amount'
                            value={amount}
                            description='Enter the cost of this expense'
                            onChange={e => setAmount(e.target.value)}
                        />
                        <div className='Editor-buttons'>
                            <Link to='/'>
                                <Button text='Cancel' />
                            </Link>
                            <div>
                                {renderDeleteButton()}
                                <Button color='orange' text={renderSubmitButtonText()} onClick={() => onExpenseSubmit()} />
                            </div>
                        </div>
                    </Card>
                </div>
            </div>
        </>
    );
};

export default Editor;
