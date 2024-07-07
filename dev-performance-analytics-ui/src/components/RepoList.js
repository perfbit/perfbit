// src/components/RepoList.js
import React from 'react';
import { Link } from 'react-router-dom';
import './RepoList.css';

const RepoList = ({ repos, token }) => {
  return (
    <div className="repo-list">
      <h2>Repositories</h2>
      <ul>
        {repos.map(repo => (
          <li key={repo.id}>
            <Link to={`/repos/${repo.id}/branches`}>{repo.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default RepoList;
