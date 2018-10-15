import React, { Component } from 'react';
import { Route, Redirect } from 'react-router-dom';

class AuthenticatedRoute extends Component {
    renderPropHandler = () => {
        const Page = this.props.component;

        return (
            this.props.authenticated
                ? <Page {...this.props} />
                : <Redirect to='/login' />
        );
    };

    render() {
        // Create props object containing all props except component
        const { component, ...props } = this.props;

        return (
            <Route
                {...props}
                render={this.renderPropHandler}
            />
        );
    }
}

export default AuthenticatedRoute;
