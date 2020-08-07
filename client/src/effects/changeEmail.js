import { putToServer } from './helpers';
import { handleResponseError } from './responses';

export default async data => {
    const response = await putToServer('/api/v1/accounts/email/change', data);

    switch (response.status) {
        case 200:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
