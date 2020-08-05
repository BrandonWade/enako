import { headToServer } from './helpers';
import { handleResponseError } from './responses';

export default async () => {
    const response = await headToServer('/api/v1/csrf');
    const headers = Array.from(response.headers.entries());
    const tokenPair = headers.filter(h => h[0] === 'x-csrf-token');

    if (tokenPair.length !== 1) {
        return {
            messages: [
                {
                    type: 'error',
                    text: 'There was an error establishing your session. Please clear your browser cache and try again.',
                },
            ],
        };
    }

    const [, token] = tokenPair[0];
    return { token };
};
