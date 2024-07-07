import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

const api = {
    login: (username, password) => {
        return axios.post(`${API_URL}/login`, { username, password });
    },
    getRepositories: (token) => {
        return axios.get(`${API_URL}/repos`, { headers: { Authorization: `Bearer ${token}` } });
    },
    getBranches: (token, repoId) => {
        return axios.get(`${API_URL}/repos/${repoId}/branches`, { headers: { Authorization: `Bearer ${token}` } });
    },
    getCommits: (token, repoId, branch) => {
        return axios.get(`${API_URL}/repos/${repoId}/branches/${branch}/commits`, { headers: { Authorization: `Bearer ${token}` } });
    },
    getPerformanceMetrics: (token) => {
        return axios.get(`${API_URL}/metrics`, { headers: { Authorization: `Bearer ${token}` } });
    }
};

export default api;
