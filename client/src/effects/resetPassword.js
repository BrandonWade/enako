import { postToServer } from './helpers';

export default async data => {
    const response = await postToServer('/api/v1/accounts/password/reset', data);

    switch (response.status) {
        case 200:
            return;
        case 404:
            return await response.json();
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue while logging in. Please refresh the page and try again.',
                    },
                ],
            };
        default:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Uh oh, something unexpected happened. Please refresh the page and try again.',
                    },
                ],
            };
    }
};
