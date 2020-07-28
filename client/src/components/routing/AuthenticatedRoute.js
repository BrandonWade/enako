import React, { useContext } from 'react';
import AuthenticatedContext from '../../contexts/AuthenticatedContext';
import { Route, Redirect } from 'react-router-dom';

const AuthenticatedRoute = props => {
    const auth = useContext(AuthenticatedContext);

    const renderPropHandler = () => {
        const Page = props.component;
        return auth.authenticated ? <Page {...props} /> : <Redirect to='/login' />;
    };

    // Create props object containing all props except component
    const { component, ...rest } = props;

    return <Route {...rest} render={renderPropHandler} />;
};

export default AuthenticatedRoute;
