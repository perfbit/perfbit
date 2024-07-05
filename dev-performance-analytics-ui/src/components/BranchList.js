// src/components/BranchList.js
import React, { useState, useEffect } from 'react';
import api from '../services/api';
import CommitList from './CommitList';

const BranchList = ({ repo, token }) => {
    const [branches, setBranches] = useState([]);
    const [selectedBranch, setSelectedBranch] = useState(null);

    useEffect(() => {
        const fetchBranches = async () => {
            try {
                const response = await api.getBranches(token, repo.owner.login, repo.name);
                setBranches(response.data);
            } catch (error) {
                console.error('Failed to fetch branches:', error);
            }
        };
        fetchBranches();
    }, [repo, token]);

    return (
        <div>
            <h3>Branches in {repo.name}</h3>
            <ul>
                {branches.map((branch) => (
                    <li key={branch.name} onClick={() => setSelectedBranch(branch)}>
                        {branch.name}
                    </li>
                ))}
            </ul>
            {selectedBranch && <CommitList repo={repo} branch={selectedBranch} token={token} />}
        </div>
    );
};

export default BranchList;
