import React, { useState, useEffect } from 'react';
import api from '../services/api';
import RepoList from './RepoList';

const Dashboard = ({ token }) => {
    const [repos, setRepos] = useState([]);

    useEffect(() => {
        const fetchRepos = async () => {
            try {
                const response = await api.getRepositories(token);
                setRepos(response.data);
            } catch (error) {
                console.error('Failed to fetch repositories:', error);
            }
        };
        fetchRepos();
    }, [token]);

    return (
        <div>
            <h2>Dashboard</h2>
            <RepoList repos={repos} token={token} />
        </div>
    );
};

export default Dashboard;
