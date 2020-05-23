import React, { useContext } from 'react';
import AuthenticatedContext from '../../contexts/AuthenticatedContext';
import { Redirect } from 'react-router-dom';

const AuthenticatedRedirect = () => {
    const authenticated = useContext(AuthenticatedContext);
    return authenticated ? <Redirect to='/' /> : <Redirect to='/login' />;
};

export default AuthenticatedRedirect;
