import { fetchFromServer } from '../effects/helpers';
import { handleResponseError } from './responses';

export default async () => {
    const response = await fetchFromServer('/api/v1/expenses');

    switch (response.status) {
        case 200:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
