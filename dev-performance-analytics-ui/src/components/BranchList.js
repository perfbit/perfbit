// src/components/BranchList.js
import React, { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import api from '../services/api';
import './BranchList.css';

const BranchList = ({ token }) => {
  const { id } = useParams();
  const [branches, setBranches] = useState([]);

  useEffect(() => {
    const fetchBranches = async () => {
      try {
        const response = await api.getBranches(token, id);
        setBranches(response.data);
      } catch (error) {
        console.error('Failed to fetch branches:', error);
      }
    };
    fetchBranches();
  }, [id, token]);

  return (
    <div className="branch-list">
      <h2>Branches</h2>
      <ul>
        {branches.map(branch => (
          <li key={branch.name}>
            <Link to={`/repos/${id}/branches/${branch.name}/commits`}>{branch.name}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BranchList;
