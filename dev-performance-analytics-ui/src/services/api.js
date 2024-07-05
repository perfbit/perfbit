import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

const api = {
    login: (username, password) => {
        return axios.post(`${API_URL}/login`, { username, password });
    },
    getRepositories: (token) => {
        return axios.get(`${API_URL}/repos`, { headers: { Authorization: token } });
    },
    getBranches: (token, owner, repo) => {
        return axios.get(`${API_URL}/repos/${owner}/${repo}/branches`, { headers: { Authorization: token } });
    },
    getCommits: (token, owner, repo, branch) => {
        return axios.get(`${API_URL}/repos/${owner}/${repo}/branches/${branch}/commits`, { headers: { Authorization: token } });
    },
};

export default api;
