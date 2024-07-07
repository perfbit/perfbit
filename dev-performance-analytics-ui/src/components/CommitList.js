// src/components/CommitList.js
import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import api from '../services/api';
import './CommitList.css';

const CommitList = ({ token }) => {
  const { id, branch } = useParams();
  const [commits, setCommits] = useState([]);

  useEffect(() => {
    const fetchCommits = async () => {
      try {
        const response = await api.getCommits(token, id, branch);
        setCommits(response.data);
      } catch (error) {
        console.error('Failed to fetch commits:', error);
      }
    };
    fetchCommits();
  }, [id, branch, token]);

  return (
    <div className="commit-list">
      <h2>Commits</h2>
      <ul>
        {commits.map(commit => (
          <li key={commit.sha}>
            <p>{commit.commit.message}</p>
            <small>By {commit.commit.author.name} on {new Date(commit.commit.author.date).toLocaleString()}</small>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CommitList;
