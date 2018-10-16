import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Card from '../Card';
import InputField from '../InputField';
import SelectField from '../SelectField';
import Button from '../Button';
import types from '../../data/SampleTypes';
import categories from '../../data/SampleCategories';
import '../../css/Editor.css';

class Editor extends Component {
    constructor(props) {
        super(props);

        this.state = {
            selectedDate: this.props.location.state.selectedDate,
            type: this.props.location.state.type,
            category: this.props.location.state.category,
            description: this.props.location.state.description,
            amount: this.props.location.state.amount,
        };
    }

    onFieldChange = (evt) => {
        this.setState({
            [evt.target.name]: evt.target.value,
        });
    };

    renderHeadingText = () => {
        return this.props.computedMatch.params.id ? (
            `Editing expense on ${this.state.selectedDate}`
        ) : (
            `Creating new expense on ${this.state.selectedDate}`
        );
    };

    renderSubmitButtonText = () => {
        return this.props.computedMatch.params.id ? 'Save' : 'Create';
    };

    render() {
        return (
            <div className='Editor'>
                <div className='Editor-content'>
                    <Card heading={this.renderHeadingText()}>
                        <div className='Editor-form'>
                            <SelectField
                                name='type'
                                label='Type'
                                value={this.state.type}
                                description='Choose the most relevant type of expense'
                                onChange={this.onFieldChange}
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
                                name='category'
                                label='Category'
                                value={this.state.category}
                                description='Choose the most relevant category of expense'
                                onChange={this.onFieldChange}
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
                                name='description'
                                label='Description'
                                value={this.state.description}
                                description='Give a brief description of this expense'
                                onChange={this.onFieldChange}
                            />
                            <InputField
                                name='amount'
                                label='Amount'
                                value={this.state.amount}
                                description='Enter the cost of this expense'
                                onChange={this.onFieldChange}
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
