import { postToServer } from './helpers';
import { handleResponseError } from './responses';

export default async data => {
    const response = await postToServer('/api/v1/accounts', data);

    switch (response.status) {
        case 201:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
