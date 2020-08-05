export const handleResponseError = async status => {
    switch (status) {
        case 403:
            return {
                messages: [
                    {
                        type: 'error',
                        text: 'It looks like your session is no longer valid. Please refresh the page and try again.',
                    },
                ],
            };
        case 500:
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
