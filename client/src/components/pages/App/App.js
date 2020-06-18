import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import retrieveCSRFToken from '../../../effects/retrieveCSRFToken';
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
    const [authenticated, setAuthenticated] = useState(document.cookie.includes('enako-session'));
    const [selectedDate, setSelectedDate] = useState(new Date());

    const [categories, setCategories] = useState([]);
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {
        const Boot = async () => {
            // TODO: Try and figure out how to make this run only once (different useEffect? check context value?)
            // TODO: Save CSRF token from header to Context
            await retrieveCSRFToken();

            if (!authenticated) {
                return;
            }

            const categories = await fetchCategories();
            const expenses = await fetchExpenses();

            // TODO: Implement proper error handling
            if (categories.errors || expenses.errors) {
                console.error(categories);
                console.error(expenses);
                return;
            }

            setCategories(categories);
            setExpenses(expenses);
        };
        Boot();
    }, [authenticated]);

    return (
        <AuthenticatedContext.Provider value={authenticated}>
            <SelectedDateContext.Provider value={selectedDate}>
                <CategoryContext.Provider value={categories}>
                    <ExpenseContext.Provider value={expenses}>
                        <BrowserRouter>
                            <Switch>
                                <Route
                                    path='/login'
                                    render={() => (
                                        <Login setAuthenticated={setAuthenticated} setCategories={setCategories} setExpenses={setExpenses} />
                                    )}
                                />
                                <Route path='/register' component={Register} />
                                <AuthenticatedRoute path='/' exact component={Home} selectedDate={selectedDate} setSelectedDate={setSelectedDate} />
                                <AuthenticatedRoute
                                    path='/expenses'
                                    exact
                                    component={Editor}
                                    setExpenses={setExpenses}
                                    setSelectedDate={setSelectedDate}
                                />
                                <AuthenticatedRoute
                                    path='/expenses/:id'
                                    component={Editor}
                                    setExpenses={setExpenses}
                                    setSelectedDate={setSelectedDate}
                                />
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
