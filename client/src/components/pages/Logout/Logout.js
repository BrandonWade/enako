import { useEffect } from 'react';
import logoutOfAccount from '../../../effects/logoutOfAccount';

const Logout = () => {
    useEffect(() => {
        const logout = async () => {
            await logoutOfAccount();
        };
        logout();
    }, []);

    window.location = '/login';

    return null;
};

export default Logout;
