import { postToServer } from '../effects/helpers';

export default async data => {
    return await postToServer('/api/v1/accounts/password', data);
};
