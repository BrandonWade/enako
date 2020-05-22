import React, { useState, useContext } from 'react';
import moment from 'moment';
import { Link } from 'react-router-dom';
import TypeContext from '../../../contexts/TypeContext';
import CategoryContext from '../../../contexts/CategoryContext';
import ExpenseContext from '../../../contexts/ExpenseContext';
import Card from '../../atoms/Card';
import Button from '../../atoms/Button';
import InputField from '../../molecules/InputField';
import SelectField from '../../molecules/SelectField';
import './Editor.css';

const Editor = props => {
    const types = useContext(TypeContext);
    const categories = useContext(CategoryContext);
    const expenses = useContext(ExpenseContext);
    const expense = expenses.find(e => e.id === parseInt(props.computedMatch.params.id)) || {};

    const [type, setType] = useState(expense.expense_type || '');
    const [category, setCategory] = useState(expense.expense_category || '');
    const [description, setDescription] = useState(expense.expense_description || '');
    const [amount, setAmount] = useState(expense.expense_amount || 0);

    const renderHeadingText = () => {
        const formattedDate = moment(props.selectedDate).format('MMMM Do, YYYY');
        return props.computedMatch.params.id ? `Editing an expense on ${formattedDate}` : `Creating a new expense on ${formattedDate}`;
    };

    const renderSubmitButtonText = () => {
        return props.computedMatch.params.id ? 'Save' : 'Create';
    };

    return (
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
                            <Button primary text={renderSubmitButtonText()} />
                        </div>
                    </div>
                </Card>
            </div>
        </div>
    );
};

export default Editor;
