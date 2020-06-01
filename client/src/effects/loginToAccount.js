import { postToServer } from './helpers';

export default async data => {
    postToServer('/api/v1/login', data);
};
