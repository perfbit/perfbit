import React, { useState, useEffect } from 'react';
import { Container, Box, Typography, TextField, Button, Grid, Link } from '@mui/material';
import GitHubIcon from '@mui/icons-material/GitHub';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';
import './Login.css';

const Login = ({ setToken }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    // Extract the token from the URL if present
    const urlParams = new URLSearchParams(window.location.search);
    const token = urlParams.get('token');
    if (token) {
      setToken(token);
      navigate('/dashboard');
    }
  }, [setToken, navigate]);

  const handleStandardLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await api.login(username, password);
      setToken(response.data.token);
      navigate('/dashboard'); // Redirect to dashboard or home page
    } catch (error) {
      console.error('Failed to login:', error);
    }
  };

  const handleGitHubLogin = () => {
    window.location.href = 'http://localhost:8080/auth/github/login';
  };

  return (
    <Container component="main" maxWidth="xs" className="container">
      <Box className="box">
        <Typography component="h1" variant="h5">
          Log in
        </Typography>
        <Box component="form" onSubmit={handleStandardLogin} className="form-container">
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            autoFocus
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            sx={{ mt: 3, mb: 2 }}
          >
            Login
          </Button>
          <Grid container className="grid-container">
            <Grid item xs>
              <Link href="#" variant="body2">
                Forgot password?
              </Link>
            </Grid>
            <Grid item>
              <Link href="#" variant="body2">
                {"Don't have an account? Register here"}
              </Link>
            </Grid>
          </Grid>
          <Button
            fullWidth
            variant="outlined"
            startIcon={<GitHubIcon />}
            color="inherit"
            sx={{ mt: 3, mb: 2 }}
            onClick={handleGitHubLogin}
            className="github-button"
          >
            Login with GitHub
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Login;
