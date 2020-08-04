import { putToServer } from './helpers';

export default async (id, data) => {
    const response = await putToServer(`/api/v1/expenses/${id}`, data);

    switch (response.status) {
        case 200:
            return await response.json();
        case 400:
        case 404:
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue updating that expense. Please refresh the page and try again.',
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
