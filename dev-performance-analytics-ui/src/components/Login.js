// src/components/Login.js
import React, { useState } from 'react';
import { Container, Box, Typography, TextField, Button, Grid, Link } from '@mui/material';
import GitHubIcon from '@mui/icons-material/GitHub';
import './Login.css';

const Login = ({ setToken }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleStandardLogin = async (e) => {
    e.preventDefault();
    console.log('Standard login with', username, password);
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
