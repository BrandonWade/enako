import { fetchFromServer } from './helpers';

export default async () => {
    const categories = await fetchFromServer('/api/v1/categories');
    const expenses = await fetchFromServer('/api/v1/expenses');

    if (categories.errors || expenses.errors) {
        return {
            errors: [...categories.errors, ...expenses.errors],
        };
    }

    return {
        categories,
        expenses,
    };
};
