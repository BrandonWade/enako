import { fetchFromServer } from './helpers';

export default async () => {
    return await fetchFromServer('/api/v1/logout');
};
