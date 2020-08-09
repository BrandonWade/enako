import { postToServer } from './helpers';
import { handleResponseError } from './responses';

export default async () => {
    const response = await postToServer('/api/v1/accounts/email/change');

    switch (response.status) {
        case 200:
            return;
        default:
            return handleResponseError(response);
    }
};
