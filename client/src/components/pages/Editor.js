import React, { Component } from 'react';
import moment from 'moment';
import { Link } from 'react-router-dom';
import Card from '../Card';
import InputField from '../InputField';
import SelectField from '../SelectField';
import Button from '../Button';
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
        const formattedDate = moment(this.props.selectedDate).format('MMMM Do YYYY');
        return this.props.computedMatch.params.id ? (
            `Editing expense on ${formattedDate}`
        ) : (
            `Creating new expense on ${formattedDate}`
        );
    };

    renderSubmitButtonText = () => {
        return this.props.computedMatch.params.id ? 'Save' : 'Create';
    };

    render() {
        const { types, categories } = this.props.location.state;
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
                                                value={type.type_name}
                                            >
                                                {type.type_name}
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
                                                value={category.category_name}
                                            >
                                                {category.category_name}
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
