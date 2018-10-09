import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import moment from 'moment';
import types from '../data/SampleTypes';
import categories from '../data/SampleCategories';
import Card from '../components/Card';
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
                            <section className='Editor-formSection'>
                                <label className='Editor-fieldLabel'>
                                    Type
                                </label>
                                <select>
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
                                </select>
                                <div className='Editor-fieldDescription'>
                                    Choose the most relevant type of expense
                                </div>
                            </section>
                            <section className='Editor-formSection'>
                                <label className='Editor-fieldLabel'>
                                    Category
                                </label>
                                <select>
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
                                </select>
                                <div className='Editor-fieldDescription'>
                                    Choose the most relevant category of expense
                                </div>
                            </section>
                            <section className='Editor-formSection'>
                                <label className='Editor-fieldLabel'>
                                    Description
                                </label>
                                <input type='text' />
                                <div className='Editor-fieldDescription'>
                                    Give a brief description of this expense
                                </div>
                            </section>
                            <section className='Editor-formSection'>
                                <label className='Editor-fieldLabel'>
                                    Amount
                                </label>
                                <input type='text' />
                                <div className='Editor-fieldDescription'>
                                    Enter the cost of this expense
                                </div>
                            </section>
                        </div>
                        <div className='Editor-formButtons'>
                            <Link to='/'>
                                <button>
                                    Cancel
                                </button>
                            </Link>
                            <button>
                                {this.renderSubmitButtonText()}
                            </button>
                        </div>
                    </Card>
                </div>
            </div>
        );
    }
}

export default Editor;
