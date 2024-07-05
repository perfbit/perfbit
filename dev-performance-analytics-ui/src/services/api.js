// src/services/api.js
import axios from 'axios';

const API_URL = 'http://localhost:8080/api/v1';

const login = (username, password) => {
    return axios.post(`${API_URL}/login`, { username, password });
};

const getRepositories = (token) => {
    return axios.get(`${API_URL}/repos`, { headers: { Authorization: token } });
};

const getBranches = (token, owner, repo) => {
    return axios.get(`${API_URL}/repos/${owner}/${repo}/branches`, { headers: { Authorization: token } });
};

const getCommits = (token, owner, repo, branch) => {
    return axios.get(`${API_URL}/repos/${owner}/${repo}/branches/${branch}/commits`, { headers: { Authorization: token } });
};

export default {
    login,
    getRepositories,
    getBranches,
    getCommits
};
