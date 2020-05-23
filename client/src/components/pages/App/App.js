import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import AuthenticatedContext from '../../../contexts/AuthenticatedContext';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Home from '../Home';
import Login from '../Login';
import Editor from '../Editor';
import './App.scss';

const App = () => {
    const authenticated = true; // TODO: For testing
    const [selectedDate, setSelectedDate] = useState(new Date());

    const [, setTypes] = useState();
    const [, setCategories] = useState();
    const [, setExpenses] = useState();

    useEffect(() => {
        const boot = async () => {
            const bootInfo = await fetchBootInfo();

            setTypes(bootInfo.types);
            setCategories(bootInfo.Categories);
            setExpenses(bootInfo.expenses);
        };
        boot();
    }, []);

    return (
        <AuthenticatedContext.Provider value={authenticated}>
            <SelectedDateContext.Provider value={selectedDate}>
                <BrowserRouter>
                    <Switch>
                        <Route path='/login' component={Login} />
                        <AuthenticatedRoute path='/' exact selectedDate={selectedDate} setSelectedDate={setSelectedDate} component={Home} />
                        <AuthenticatedRoute path='/expenses' exact component={Editor} />
                        <AuthenticatedRoute path='/expenses/:id' selectedDate={selectedDate} component={Editor} />
                        <AuthenticatedRedirect />
                    </Switch>
                </BrowserRouter>
            </SelectedDateContext.Provider>
        </AuthenticatedContext.Provider>
    );
};

export default App;
