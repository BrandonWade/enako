import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import Home from '../Home';
import Login from '../Login';
import Editor from '../Editor';
import './App.css';

const App = () => {
    const authenticated = true; // TODO: For testing

    return (
        <BrowserRouter>
            <Switch>
                <Route path='/login' component={Login} />
                <AuthenticatedRoute path='/' exact={true} authenticated={authenticated} component={Home} />
                <AuthenticatedRoute path='/expenses' exact={true} authenticated={authenticated} component={Editor} />
                <AuthenticatedRoute path='/expenses/:id' authenticated={authenticated} component={Editor} />
                <AuthenticatedRedirect authenticated={authenticated} />
            </Switch>
        </BrowserRouter>
    );
};

export default App;
