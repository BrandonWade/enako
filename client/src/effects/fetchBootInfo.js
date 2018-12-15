import { fetchFromServer } from './helpers';

const fetchBootInfo = async () => {
    const types = await fetchFromServer('/api/v1/types');
    const categories = await fetchFromServer('/api/v1/categories');

    if (types.errors || categories.errors) {
        return {
            errors: [types.errors, categories.errors],
        }
    }

    return {
        types,
        categories,
    };
}

export default fetchBootInfo;
