import React from 'react';
import { Redirect } from 'react-router-dom';

const AuthenticatedRedirect = ({ authenticated }) => {
    return <>{authenticated ? <Redirect to='/' /> : <Redirect to='/login' />}</>;
};

export default AuthenticatedRedirect;
