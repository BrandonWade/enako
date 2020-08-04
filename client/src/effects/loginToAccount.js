import { postToServer } from './helpers';

export default async data => {
    const response = await postToServer('/api/v1/login', data);

    switch (response.status) {
        case 200:
            return;
        case 401:
            return await response.json(); // TODO: This represents multiple cases that need handling
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
