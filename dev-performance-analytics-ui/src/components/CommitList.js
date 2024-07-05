// src/components/CommitList.js
import React, { useState, useEffect } from 'react';
import api from '../services/api';

const CommitList = ({ repo, branch, token }) => {
    const [commits, setCommits] = useState([]);

    useEffect(() => {
        const fetchCommits = async () => {
            try {
                const response = await api.getCommits(token, repo.owner.login, repo.name, branch.name);
                setCommits(response.data);
            } catch (error) {
                console.error('Failed to fetch commits:', error);
            }
        };
        fetchCommits();
    }, [repo, branch, token]);

    return (
        <div>
            <h3>Commits in {branch.name} branch</h3>
            <ul>
                {commits.map((commit) => (
                    <li key={commit.sha}>
                        {commit.commit.message}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default CommitList;
