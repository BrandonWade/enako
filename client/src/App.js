import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Home from './pages/Home';
import Editor from './pages/Editor';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <Switch>
                    <Route exact path='/' component={Home} />
                    <Route path='/create' component={Editor} />
                    <Route path='/edit' component={Editor} />
                </Switch>
            </BrowserRouter>
        );
    }
}

export default App;
