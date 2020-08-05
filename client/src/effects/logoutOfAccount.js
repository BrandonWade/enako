import { fetchFromServer } from './helpers';
import { handleResponseError } from './responses';

export default async () => {
    const response = await fetchFromServer('/api/v1/logout');

    switch (response.status) {
        case 204:
            return;
        default:
            return handleResponseError(response);
    }
};
