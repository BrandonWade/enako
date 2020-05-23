import { postToServer } from './helpers';

export default async data => {
    return await postToServer('/api/v1/expenses', data);
};
