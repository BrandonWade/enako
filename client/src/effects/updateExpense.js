import { putToServer } from './helpers';
import { handleResponseError } from './responses';

export default async (id, data) => {
    const response = await putToServer(`/api/v1/expenses/${id}`, data);

    switch (response.status) {
        case 200:
            return await response.json();
        default:
            return handleResponseError(response);
    }
};
