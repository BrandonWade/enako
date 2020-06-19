import { headToServer } from './helpers';

export default async () => {
    const response = await headToServer('/api/v1/csrf');
    const headers = Array.from(response.headers.entries());
    const tokenPair = headers.filter(h => h[0] === 'x-csrf-token');

    if (tokenPair.length !== 1) {
        return { errors: ['error retrieving CSRF token'] };
    }

    const [, token] = tokenPair[0];
    return { token };
};
