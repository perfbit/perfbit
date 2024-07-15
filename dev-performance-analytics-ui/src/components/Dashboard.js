import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Container, Typography, AppBar, Toolbar, Button, Grid, Paper } from '@mui/material';
import api from '../services/api';
import RepoList from './RepoList';
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
    <Container maxWidth="lg" className="dashboard">
      <AppBar position="static" className="app-bar">
        <Toolbar>
          <Typography variant="h6" className="title">
            Perfbit Dashboard
          </Typography>
          <Button color="inherit" component={Link} to="/repos">Repositories</Button>
          <Button color="inherit" component={Link} to="/metrics">Performance Metrics</Button>
        </Toolbar>
      </AppBar>
      <Grid container spacing={3} className="content">
        <Grid item xs={12}>
          <Paper className="paper">
            <Typography variant="h5">Repositories</Typography>
            <RepoList repos={repos} token={token} />
          </Paper>
        </Grid>
        <Grid item xs={12}>
          <Paper className="paper">
            <Typography variant="h5">Performance Metrics</Typography>
            {/* Placeholder for Performance Metrics Component */}
            <Typography variant="body1">Metrics will be displayed here.</Typography>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default Dashboard;
