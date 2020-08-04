import { fetchFromServer } from '../effects/helpers';

export default async () => {
    const response = await fetchFromServer('/api/v1/categories');

    switch (response.status) {
        case 200:
            return await response.json();
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue fetching the expense categories. Please refresh the page and try again.',
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
