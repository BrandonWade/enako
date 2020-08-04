import { postToServer } from '../effects/helpers';

export default async data => {
    const response = await postToServer('/api/v1/accounts/password', data);

    switch (response.status) {
        case 200:
        case 404:
            return await response.json();
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue while processing your request. Please refresh the page and try again.',
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
