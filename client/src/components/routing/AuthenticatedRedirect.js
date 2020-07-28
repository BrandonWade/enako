import React, { useContext } from 'react';
import { Redirect } from 'react-router-dom';
import AuthenticatedContext from '../../contexts/AuthenticatedContext';

const AuthenticatedRedirect = () => {
    const auth = useContext(AuthenticatedContext);
    return auth.authenticated ? <Redirect to='/' /> : <Redirect to='/login' />;
};

export default AuthenticatedRedirect;
