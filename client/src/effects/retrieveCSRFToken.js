import { headToServer } from './helpers';

export default async () => {
    return await headToServer('/api/v1/csrf');
};
