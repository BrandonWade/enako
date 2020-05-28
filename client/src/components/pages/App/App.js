import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import AuthenticatedContext from '../../../contexts/AuthenticatedContext';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import CategoryContext from '../../../contexts/CategoryContext';
import ExpenseContext from '../../../contexts/ExpenseContext';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Home from '../Home';
import Login from '../Login';
import Editor from '../Editor';
import './App.scss';

const App = () => {
    const authenticated = true; // TODO: For testing
    const [selectedDate, setSelectedDate] = useState(new Date());

    const [categories, setCategories] = useState([]);
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {
        const boot = async () => {
            const bootInfo = await fetchBootInfo();

            setCategories(bootInfo.categories);
            setExpenses(bootInfo.expenses);
        };
        boot();
    }, []);

    return (
        <AuthenticatedContext.Provider value={authenticated}>
            <SelectedDateContext.Provider value={selectedDate}>
                <CategoryContext.Provider value={categories}>
                    <ExpenseContext.Provider value={expenses}>
                        <BrowserRouter>
                            <Switch>
                                <Route path='/login' component={Login} />
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
