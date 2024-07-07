// src/components/Dashboard.js
import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import api from '../services/api';
import './Dashboard.css';

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
        <div className="dashboard">
            <h2>Dashboard</h2>
            <nav>
                <ul>
                    <li><Link to="/repos">Repositories</Link></li>
                    <li><Link to="/metrics">Performance Metrics</Link></li>
                </ul>
            </nav>
            <RepoList repos={repos} token={token} />
        </div>
    );
};

export default Dashboard;
