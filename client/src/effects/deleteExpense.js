import { deleteFromServer } from './helpers';
import { handleResponseError } from './responses';

export default async id => {
    const response = await deleteFromServer(`/api/v1/expenses/${id}`);

    switch (response.status) {
        case 204:
            return;
        default:
            return handleResponseError(response);
    }
};
