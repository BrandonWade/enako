import { createContext } from 'react';

export default createContext({
    authenticated: false,
    setAuthenticated: () => {}, // TODO: Remove this?
});
