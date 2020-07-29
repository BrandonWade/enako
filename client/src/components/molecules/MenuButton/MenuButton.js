import React, { useContext } from 'react';
import { Link, useLocation } from 'react-router-dom';
import AuthenticatedContext from '../../../contexts/AuthenticatedContext';
import { CogIcon } from '../../atoms/Icons';
import './MenuButton.scss';

const MenuButton = () => {
    const authenticated = useContext(AuthenticatedContext);
    const location = useLocation();
    const isAccount = location.pathname === '/account';
    const isLogout = location.pathname === '/logout';

    return authenticated && !isAccount && !isLogout ? (
        <Link to='/account'>
            <CogIcon className='MenuButton' />
        </Link>
    ) : null;
};

export default MenuButton;
