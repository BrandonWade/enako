import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import fetchCategories from '../../../effects/fetchCategories';
import fetchExpenses from '../../../effects/fetchExpenses';
import AuthenticatedContext from '../../../contexts/AuthenticatedContext';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import CategoryContext from '../../../contexts/CategoryContext';
import ExpenseContext from '../../../contexts/ExpenseContext';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Home from '../Home';
import Register from '../Register';
import Login from '../Login';
import Editor from '../Editor';
import './App.scss';

const App = () => {
    const authenticated = true; // TODO: For testing
    const [selectedDate, setSelectedDate] = useState(new Date());

    const [categories, setCategories] = useState([]);
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {
        const Boot = async () => {
            const categories = await fetchCategories();
            const expenses = await fetchExpenses();

            setCategories(categories);
            setExpenses(expenses);
        };
        Boot();
    }, []);

    return (
        <AuthenticatedContext.Provider value={authenticated}>
            <SelectedDateContext.Provider value={selectedDate}>
                <CategoryContext.Provider value={categories}>
                    <ExpenseContext.Provider value={expenses}>
                        <BrowserRouter>
                            <Switch>
                                <Route path='/login' component={Login} />
                                <Route path='/register' component={Register} />
                                <AuthenticatedRoute path='/' exact component={Home} selectedDate={selectedDate} setSelectedDate={setSelectedDate} />
                                <AuthenticatedRoute path='/expenses' exact component={Editor} setExpenses={setExpenses} />
                                <AuthenticatedRoute path='/expenses/:id' component={Editor} setExpenses={setExpenses} />
                                <AuthenticatedRedirect />
                            </Switch>
                        </BrowserRouter>
                    </ExpenseContext.Provider>
                </CategoryContext.Provider>
            </SelectedDateContext.Provider>
        </AuthenticatedContext.Provider>
    );
};

export default App;
