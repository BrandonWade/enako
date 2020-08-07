import React, { useState, useEffect } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import retrieveCSRFToken from '../../../effects/retrieveCSRFToken';
import fetchCategories from '../../../effects/fetchCategories';
import fetchExpenses from '../../../effects/fetchExpenses';
import AuthenticatedContext from '../../../contexts/AuthenticatedContext';
import SelectedDateContext from '../../../contexts/SelectedDateContext';
import CategoryContext from '../../../contexts/CategoryContext';
import ExpenseContext from '../../../contexts/ExpenseContext';
import MessageContext from '../../../contexts/MessageContext';
import AuthenticatedRoute from '../../routing/AuthenticatedRoute';
import AuthenticatedRedirect from '../../routing/AuthenticatedRedirect';
import MenuButton from '../../molecules/MenuButton';
import Home from '../Home';
import Register from '../Register';
import Login from '../Login';
import ForgotPassword from '../ForgotPassword';
import ChangePassword from '../ChangePassword';
import ChangeEmail from '../ChangeEmail';
import Editor from '../Editor';
import Account from '../Account';
import Logout from '../Logout';
import './App.scss';

const App = () => {
    const [messages, setMessages] = useState([]);
    const [authenticated, setAuthenticated] = useState(document.cookie.includes('enako-session'));
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [categories, setCategories] = useState([]);
    const [expenses, setExpenses] = useState([]);

    useEffect(() => {
        const fetchCSRF = async () => {
            const csrfToken = await retrieveCSRFToken();
            if (csrfToken?.messages?.length > 0) {
                setMessages(csrfToken.messages);
                return;
            }

            sessionStorage.setItem('csrfToken', csrfToken.token);
        };
        fetchCSRF();
    }, []);

    useEffect(() => {
        const fetchUserData = async () => {
            if (!authenticated) {
                return;
            }

            const categories = await fetchCategories();
            if (categories?.messages?.length > 0) {
                setMessages(categories.messages);
                return;
            }

            const expenses = await fetchExpenses();
            if (expenses?.messages?.length > 0) {
                setMessages(expenses.messages);
                return;
            }

            setCategories(categories);
            setExpenses(expenses);
        };
        fetchUserData();
    }, [authenticated]);

    return (
        <MessageContext.Provider value={{ messages, setMessages }}>
            <AuthenticatedContext.Provider value={authenticated}>
                <SelectedDateContext.Provider value={selectedDate}>
                    <CategoryContext.Provider value={categories}>
                        <ExpenseContext.Provider value={expenses}>
                            <BrowserRouter>
                                <MenuButton />
                                <Switch>
                                    <Route
                                        path='/login'
                                        render={() => (
                                            <Login setAuthenticated={setAuthenticated} setCategories={setCategories} setExpenses={setExpenses} />
                                        )}
                                    />
                                    <Route path='/password' exact render={() => <ForgotPassword />} />
                                    <Route path='/password/reset' render={() => <Register passwordReset={true} />} />
                                    <Route path='/register' component={() => <Register passwordReset={false} />} />
                                    <AuthenticatedRoute
                                        path='/'
                                        exact
                                        component={Home}
                                        selectedDate={selectedDate}
                                        setSelectedDate={setSelectedDate}
                                    />
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
                                    <AuthenticatedRoute path='/account' exact component={Account} />
                                    <AuthenticatedRoute path='/account/password' component={ChangePassword} />
                                    <AuthenticatedRoute path='/account/email' component={ChangeEmail} />
                                    <AuthenticatedRoute path='/logout' component={Logout} />
                                    <AuthenticatedRedirect />
                                </Switch>
                            </BrowserRouter>
                        </ExpenseContext.Provider>
                    </CategoryContext.Provider>
                </SelectedDateContext.Provider>
            </AuthenticatedContext.Provider>
        </MessageContext.Provider>
    );
};

export default App;
