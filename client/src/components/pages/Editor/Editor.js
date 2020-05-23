import React, { useState, useContext } from 'react';
import { format } from 'date-fns';
import { Link } from 'react-router-dom';
import createExpense from '../../../effects/createExpense';
import updateExpense from '../../../effects/updateExpense';
import deleteExpense from '../../../effects/deleteExpense';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import TypeContext from '../../../contexts/TypeContext';
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
    const types = useContext(TypeContext);
    const categories = useContext(CategoryContext);
    const expenses = useContext(ExpenseContext);
    const expense = expenses.find(e => e.id === parseInt(props.computedMatch.params.id)) || {};

    const [type, setType] = useState(expense.expense_type || '');
    const [category, setCategory] = useState(expense.expense_category || '');
    const [description, setDescription] = useState(expense.expense_description || '');
    const [amount, setAmount] = useState(expense.expense_amount || 0);

    const notFoundRedirect = () => {
        return props.computedMatch.params.id && !expense.id ? <AuthenticatedRedirect /> : null;
    };

    const renderHeadingText = () => {
        const formattedDate = format(selectedDate, 'MMMM do yyyy');
        return props.computedMatch.params.id ? `Editing an expense on ${formattedDate}` : `Creating a new expense on ${formattedDate}`;
    };

    const renderDeleteButton = () => {
        return props.computedMatch.params.id ? <Button text='Delete' className='editor__delete button--red' onClick={onExpenseDelete} /> : null;
    };

    const renderSubmitButtonText = () => {
        return props.computedMatch.params.id ? 'Save' : 'Create';
    };

    const onExpenseDelete = () => {
        const id = props.computedMatch.params.id;
        deleteExpense(id);
    };

    const onExpenseSubmit = () => {
        const id = props.computedMatch.params.id || 0;
        const data = {
            type,
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
                            name='type'
                            label='Type'
                            value={type}
                            description='Choose the most relevant type of expense'
                            onChange={e => setType(e.target.value)}
                        >
                            <option value=''>-- Select a Type -- </option>
                            {types.map(t => {
                                return (
                                    <option key={t.id} value={t.type_name}>
                                        {t.type_name}
                                    </option>
                                );
                            })}
                        </SelectField>
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
                                    <option key={c.id} value={c.category_name}>
                                        {c.category_name}
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
