export const API_CONFIG = {
    BASE_URL: import.meta.env.PUBLIC_API_URL,

    ENDPOINTS: {
        HEALTH: '/api/health',
        CONTACT: '/api/contact',
    },
};

export const getApiUrl = (endpoint: string): string => {
    return `${API_CONFIG.BASE_URL}${endpoint}`;
};
