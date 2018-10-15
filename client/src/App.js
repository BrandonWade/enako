import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import AuthenticatedRoute from './components/routing/AuthenticatedRoute';
import AuthenticatedRedirect from './components/routing/AuthenticatedRedirect';
import Home from './components/pages/Home';
import Login from './components/pages/Login';
import Editor from './components/pages/Editor';

class App extends Component {
    render() {
        const authenticated = true; // TODO: For testing

        return (
            <BrowserRouter>
                <Switch>
                    <Route
                        path='/login'
                        component={Login}
                    />
                    <AuthenticatedRoute
                        path='/'
                        exact={true}
                        authenticated={authenticated}
                        component={Home}
                    />
                    <AuthenticatedRoute
                        path='/create'
                        authenticated={authenticated}
                        component={Editor}
                    />
                    <AuthenticatedRoute
                        path='/edit'
                        authenticated={authenticated}
                        component={Editor}
                    />
                    <AuthenticatedRedirect
                        authenticated={authenticated}
                    />
                </Switch>
            </BrowserRouter>
        );
    }
}

export default App;
