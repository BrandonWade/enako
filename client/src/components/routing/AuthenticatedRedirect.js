import React, { useContext } from 'react';
import { Redirect } from 'react-router-dom';
import AuthenticatedContext from '../../contexts/AuthenticatedContext';

const AuthenticatedRedirect = () => {
    const authenticated = useContext(AuthenticatedContext);
    return authenticated ? <Redirect to='/' /> : <Redirect to='/login' />;
};

export default AuthenticatedRedirect;
