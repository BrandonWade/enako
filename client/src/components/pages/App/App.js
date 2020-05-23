import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import fetchBootInfo from '../../../effects/fetchBootInfo';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Home from '../Home';
import Login from '../Login';
import Editor from '../Editor';
import './App.scss';

const App = () => {
    const [selectedDate, setSelectedDate] = useState(new Date());
    const authenticated = true; // TODO: For testing

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
        <SelectedDateContext.Provider value={selectedDate}>
            <BrowserRouter>
                <Switch>
                    <Route path='/login' component={Login} />
                    <AuthenticatedRoute
                        path='/'
                        exact={true}
                        authenticated={authenticated}
                        selectedDate={selectedDate}
                        setSelectedDate={setSelectedDate}
                        component={Home}
                    />
                    <AuthenticatedRoute path='/expenses' exact={true} authenticated={authenticated} component={Editor} />
                    <AuthenticatedRoute path='/expenses/:id' authenticated={authenticated} selectedDate={selectedDate} component={Editor} />
                    <AuthenticatedRedirect authenticated={authenticated} />
                </Switch>
            </BrowserRouter>
        </SelectedDateContext.Provider>
    );
};

export default App;
