import { postToServer } from './helpers';

export default async data => {
    const response = await postToServer('/api/v1/accounts/password/change', data);

    switch (response.status) {
        case 200:
            return await response.json();
        case 500:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'Hmm looks like there was an issue while changing your password. Please refresh the page and try again.',
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
