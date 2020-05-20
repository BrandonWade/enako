import React from 'react';
import { Route, Redirect } from 'react-router-dom';

const AuthenticatedRoute = props => {
    const renderPropHandler = () => {
        const Page = props.component;

        return props.authenticated ? <Page {...props} /> : <Redirect to='/login' />;
    };

    // Create props object containing all props except component
    const { component, ...rest } = props;

    return <Route {...rest} render={renderPropHandler} />;
};

export default AuthenticatedRoute;
