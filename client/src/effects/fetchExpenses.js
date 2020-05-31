import { fetchFromServer } from '../effects/helpers';

export default async () => {
    return await fetchFromServer('/api/v1/expenses');
};
