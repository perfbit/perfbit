// components/GitHubLoginButton.tsx

import React from 'react';

const GitHubLoginButton: React.FC = () => {
    const handleGitHubLogin = () => {
        window.location.href = 'http://localhost:8081/auth/github';
    };

    return (
        <button
            type="button"
            onClick={handleGitHubLogin}
            className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-gray-800 hover:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-900"
        >
            Log In with GitHub
        </button>
    );
};

export default GitHubLoginButton;
