import React, { useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import logoutOfAccount from '../../../effects/logoutOfAccount';

const Logout = () => {
    useEffect(() => {
        const logout = async () => {
            await logoutOfAccount();
        };
        logout();
    }, []);

    // TODO: This causes an error when accessing /logout directly
    return <Redirect to='/login' />;
    // window.location = '/login';
};

export default Logout;
