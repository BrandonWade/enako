import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';

class AuthenticatedRedirect extends Component {
    render() {
        return (
            <>
                {
                    this.props.authenticated
                        ? <Redirect to='/' />
                        : <Redirect to='/login' />
                }
            </>
        );
    }
}

export default AuthenticatedRedirect;
