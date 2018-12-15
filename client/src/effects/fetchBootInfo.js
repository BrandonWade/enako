import { fetchFromServer } from './helpers';

const fetchBootInfo = async () => {
    const types = await fetchFromServer('/api/v1/types');
    const categories = await fetchFromServer('/api/v1/categories');
    const expenses = await fetchFromServer('/api/v1/expenses');

    if (types.errors || categories.errors) {
        return {
            errors: [types.errors, categories.errors],
        }
    }

    return {
        types,
        categories,
        expenses,
    };
}

export default fetchBootInfo;
