import { putToServer } from './helpers';

export default async (id, data) => {
    return await putToServer(`/api/v1/expenses/${id}`, data);
};
