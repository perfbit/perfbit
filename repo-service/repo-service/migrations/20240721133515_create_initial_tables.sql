-- Create table for repositories
CREATE TABLE repositories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    owner VARCHAR(255) NOT NULL,
    private BOOLEAN NOT NULL
);

-- Create table for branches
CREATE TABLE branches (
    id SERIAL PRIMARY KEY,
    repository_id INTEGER NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    commit_hash VARCHAR(255) NOT NULL,
    last_commit_message TEXT NOT NULL,
    last_commit_timestamp TIMESTAMP NOT NULL,
    last_commit_author VARCHAR(255) NOT NULL
);

-- Create table for commits
CREATE TABLE commits (
    id SERIAL PRIMARY KEY,
    repository_id INTEGER NOT NULL REFERENCES repositories(id) ON DELETE CASCADE,
    branch_id INTEGER NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
    commit_hash VARCHAR(255) NOT NULL,
    author_id VARCHAR(255) NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    author_email VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    url VARCHAR(255) NOT NULL,
    changes JSONB NOT NULL
);

CREATE TABLE github_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    github_id INTEGER NOT NULL UNIQUE,
    node_id VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    url VARCHAR(255),
    name VARCHAR(255),
    company VARCHAR(255),
    blog VARCHAR(255),
    location VARCHAR(255),
    email VARCHAR(255)
);
