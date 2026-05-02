const API_BASE = 'http://localhost:8080/api';

async function request(endpoint, options = {}) {
    const url = `${API_BASE}${endpoint}`;
    
    const config = {
        headers: {
            'Content-Type': 'application/json',
            ...options.headers,
        },
        credentials: 'include',
        ...options,
    };

    const response = await fetch(url, config);
    const data = await response.json();
    
    if (!response.ok) {
        throw new Error(data.error || 'เกิดข้อผิดพลาด');
    }
    
    return data;
}

export const api = {
    // Auth
    register: (body) => request('/auth/register', { method: 'POST', body: JSON.stringify(body) }),
    login: (body) => request('/auth/login', { method: 'POST', body: JSON.stringify(body) }),
    logout: () => request('/auth/logout', { method: 'POST' }),
    me: () => request('/auth/me'),

    // Activities
    getActivities: (params = '') => request(`/activities${params ? '?' + params : ''}`),
    getActivity: (id) => request(`/activities/${id}`),

    // Registrations
    registerActivity: (body) => request('/registrations', { method: 'POST', body: JSON.stringify(body) }),
    getMyRegistrations: () => request('/registrations/me'),
    cancelRegistration: (id) => request(`/registrations/${id}`, { method: 'DELETE' }),
    checkRegistration: (activityId) => request(`/registrations/check/${activityId}`),
};
