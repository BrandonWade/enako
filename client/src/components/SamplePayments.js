const payments = [
    {
        date: 'October 6th 2018',
        events: [
            {
                id: 1,
                payment_category: 'unnecessary expense',
                description: 'went out for lunch',
                amount: 16.80,
            },
            {
                id: 2,
                payment_category: 'recurring',
                description: 'paid phone bill for next 2 months',
                amount: 120.58,
            },
            {
                id: 3,
                payment_category: 'unnecessary expense',
                description: 'went to a movie',
                amount: 11.50,
            },
            {
                id: 4,
                payment_category: 'recurring',
                description: 'crunchyroll subscription',
                amount: 8.98,
            },
        ],
    },
    {
        date: 'October 7th 2018',
        events: [
            {
                id: 5,
                payment_category: 'unnecessary expense',
                description: 'went out for wings',
                amount: 13.29,
            },
        ],
    },
    {
        date: 'October 8th 2018',
        events: [
            {
                id: 6,
                payment_category: 'expense',
                description: 'bought The Phoenix Project',
                amount: 19.99,
            },
            {
                id: 7,
                payment_category: 'expense',
                description: 'bought new wow expansion',
                amount: 100.00,
            },
        ],
    }
];

export default payments;
