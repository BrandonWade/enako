import { fetchFromServer } from './helpers';

export default async () => {
    const response = await fetchFromServer('/api/v1/logout');

    switch (response.status) {
        case 204:
            return;
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue while logging out. Please refresh the page and try again.',
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
