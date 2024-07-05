// src/App.js
import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import './styles/App.css';

function App() {
  const [token, setToken] = useState(null);

  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/login" element={<Login setToken={setToken} />} />
          <Route path="/" element={token ? <Dashboard token={token} /> : <Login setToken={setToken} />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
