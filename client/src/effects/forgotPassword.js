import { postToServer } from '../effects/helpers';
import { handleResponseError } from './responses';

export default async data => {
    const response = await postToServer('/api/v1/accounts/password', data);

    switch (response.status) {
        case 200:
        case 404:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
