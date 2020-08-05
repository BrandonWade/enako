export const headToServer = async (url, headers = {}) => {
    const response = await fetch(url, {
        method: 'HEAD',
        headers: {
            ...headers,
        },
    });

    return response;
};

export const fetchFromServer = async (url, headers = {}) => {
    const response = await fetch(url, {
        headers: {
            ...headers,
        },
    });

    return response;
};

export const postToServer = async (url, data, headers = {}) => {
    const csrfToken = sessionStorage.getItem('csrfToken') || '';
    const response = await fetch(url, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'X-Csrf-Token': csrfToken,
        },
        body: JSON.stringify(data),
    });

    return response;
};

export const putToServer = async (url, data, headers = {}) => {
    const csrfToken = sessionStorage.getItem('csrfToken') || '';
    const response = await fetch(url, {
        method: 'PUT',
        credentials: 'same-origin',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'X-Csrf-Token': csrfToken,
        },
        body: JSON.stringify(data),
    });

    return response;
};

export const deleteFromServer = async (url, headers = {}) => {
    const csrfToken = sessionStorage.getItem('csrfToken') || '';
    const response = await fetch(url, {
        method: 'DELETE',
        credentials: 'same-origin',
        headers: {
            ...headers,
            'Content-Type': 'application/json',
            'X-Csrf-Token': csrfToken,
        },
    });

    return response;
};
