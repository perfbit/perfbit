// src/App.js
import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import './styles/App.css';

function App() {
    const [token, setToken] = useState(null);

    return (
        <Router>
            <div className="App">
                <Switch>
                    <Route path="/login">
                        <Login setToken={setToken} />
                    </Route>
                    <Route path="/">
                        {token ? <Dashboard token={token} /> : <Login setToken={setToken} />}
                    </Route>
                </Switch>
            </div>
        </Router>
    );
}

export default App;
