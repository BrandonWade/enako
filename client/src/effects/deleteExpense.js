import { deleteFromServer } from './helpers';

export default async (id, data) => {
    return await deleteFromServer(`/api/v1/expenses/${id}`, data);
};
