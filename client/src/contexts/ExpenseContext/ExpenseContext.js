import { createContext } from 'react';

export default createContext([
    {
        id: 1,
        expense_type: 'unnecessary',
        expense_category: 'food',
        expense_description: 'foo',
        expense_amount: 25.0,
        expense_date: '2020-05-22',
    },
    {
        id: 2,
        expense_type: 'unnecessary',
        expense_category: 'technology',
        expense_description: 'bar',
        expense_amount: 100.0,
        expense_date: '2020-05-22',
    },
    {
        id: 3,
        expense_type: 'unnecessary',
        expense_category: 'clothing',
        expense_description: 'baz',
        expense_amount: 75.0,
        expense_date: '2020-05-22',
    },
]);
