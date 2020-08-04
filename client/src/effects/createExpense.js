import { postToServer } from './helpers';

export default async data => {
    const response = await postToServer('/api/v1/expenses', data);

    switch (response.status) {
        case 201:
            return await response.json();
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue while creating your expense. Please refresh the page and try again.',
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
