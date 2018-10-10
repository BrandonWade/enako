import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import moment from 'moment';
import Card from '../components/Card';
import InputField from '../components/InputField';
import SelectField from '../components/SelectField';
import Button from '../components/Button';
import types from '../data/SampleTypes';
import categories from '../data/SampleCategories';
import '../css/Editor.css';

class Editor extends Component {
    constructor(props) {
        super(props);

        this.state = {
            date: moment().format('MMMM Do YYYY'), // TODO: Pass in date as prop
            expense: {},
        };
    }

    renderHeadingText = () => {
        return this.state.expense.id ? (
            `Editing expense on ${this.state.date}`
        ) : (
            `Creating new expense on ${this.state.date}`
        );
    };

    renderSubmitButtonText = () => {
        return this.state.expense.id ? 'Save' : 'Create';
    };

    render() {
        return (
            <div className='Editor'>
                <div className='Editor-content'>
                    <Card heading={this.renderHeadingText()}>
                        <div className='Editor-form'>
                            <SelectField
                                label='Type'
                                description='Choose the most relevant type of expense'
                            >
                                <option value=''>-- Select a Type -- </option>
                                {
                                    types.map((type) => {
                                        return (
                                            <option
                                                key={type.id}
                                                value={type.value}
                                            >
                                                {type.text}
                                            </option>
                                        );
                                    })
                                }
                            </SelectField>
                            <SelectField
                                label='Category'
                                description='Choose the most relevant category of expense'
                            >
                                <option value=''>-- Select a Category -- </option>
                                {
                                    categories.map((category) => {
                                        return (
                                            <option
                                                key={category.id}
                                                value={category.value}
                                            >
                                                {category.text}
                                            </option>
                                        );
                                    })
                                }
                            </SelectField>
                            <InputField
                                label='Description'
                                description='Give a brief description of this expense'
                            />
                            <InputField
                                label='Amount'
                                description='Enter the cost of this expense'
                            />
                        </div>
                        <div className='Editor-formButtons'>
                            <Link to='/'>
                                <Button text='Cancel' />
                            </Link>
                            <Button
                                main={true}
                                text={this.renderSubmitButtonText()}
                            />
                        </div>
                    </Card>
                </div>
            </div>
        );
    }
}

export default Editor;
