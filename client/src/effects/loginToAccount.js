import { postToServer } from './helpers';
import { handleResponseError } from './responses';

export default async data => {
    const response = await postToServer('/api/v1/login', data);

    switch (response.status) {
        case 200:
            return;
        case 401:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
