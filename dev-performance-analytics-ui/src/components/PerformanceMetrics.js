// src/components/PerformanceMetrics.js
import React, { useEffect, useState } from 'react';
import api from '../services/api';
import './PerformanceMetrics.css';

const PerformanceMetrics = ({ token }) => {
  const [metrics, setMetrics] = useState([]);

  useEffect(() => {
    const fetchMetrics = async () => {
      try {
        const response = await api.getPerformanceMetrics(token);
        setMetrics(response.data);
      } catch (error) {
        console.error('Failed to fetch performance metrics:', error);
      }
    };
    fetchMetrics();
  }, [token]);

  return (
    <div className="performance-metrics">
      <h2>Performance Metrics</h2>
      <ul>
        {metrics.map(metric => (
          <li key={metric.id}>
            <p>{metric.name}</p>
            <small>Commits: {metric.commits}</small>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default PerformanceMetrics;
