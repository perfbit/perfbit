import React, { useState } from 'react';
import BranchList from './BranchList';

const RepoList = ({ repos, token }) => {
    const [selectedRepo, setSelectedRepo] = useState(null);

    return (
        <div>
            <h3>Repositories</h3>
            <ul>
                {repos.map((repo) => (
                    <li key={repo.id} onClick={() => setSelectedRepo(repo)}>
                        {repo.name}
                    </li>
                ))}
            </ul>
            {selectedRepo && <BranchList repo={selectedRepo} token={token} />}
        </div>
    );
};

export default RepoList;
