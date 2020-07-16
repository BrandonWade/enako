import { createContext } from 'react';

export default createContext([
    {
        text: 'Heads up! A password reset link was sent to the email associated with your account: foo@bar.net',
        type: 'error',
    },
    {
        text: 'Heads up! A password reset link was sent to the email associated with your account: foo@bar.net',
        type: 'info',
    },
]);
