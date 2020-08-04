import { deleteFromServer } from './helpers';

export default async id => {
    await deleteFromServer(`/api/v1/expenses/${id}`);
};
