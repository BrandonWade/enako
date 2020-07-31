import React, { useContext } from 'react';
import { Redirect } from 'react-router-dom';
import AuthenticatedContext from '../../contexts/AuthenticatedContext';

const AuthenticatedRedirect = ({ authenticatedOnly }) => {
    const authenticated = useContext(AuthenticatedContext);
    const unauthenticatedTarget = !authenticatedOnly ? <Redirect to='/login' /> : null;
    return authenticated ? <Redirect to='/' /> : unauthenticatedTarget;
};

export default AuthenticatedRedirect;
