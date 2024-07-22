use axum::{extract::Extension, response::IntoResponse, Json};
use serde::{Deserialize, Serialize};
use sqlx::PgPool;
use reqwest::Client;

#[derive(Deserialize)]
struct GitHubCredentials {
    access_token: String,
    token_type: String,
}

#[derive(Deserialize, Serialize, Debug)]
struct GitHubRepo {
    id: i32,
    name: String,
    full_name: String,
    description: Option<String>,
    html_url: String,
    created_at: String,
    updated_at: String,
    private: bool,
    owner: GitHubOwner,
}

#[derive(Deserialize, Serialize, Debug)]
struct GitHubOwner {
    login: String,
}

pub async fn fetch_repositories(Extension(pool): Extension<PgPool>, Extension(client): Extension<Client>) -> impl IntoResponse {
    // Get GitHub credentials from auth service
    let auth_service_url = "http://localhost:8081/auth/github/credentials";
    let credentials: GitHubCredentials = client
        .get(auth_service_url)
        .send()
        .await
        .expect("Failed to fetch credentials from auth service")
        .json()
        .await
        .expect("Failed to parse credentials response");

    // Fetch repositories from GitHub using credentials
    let repos_url = "https://api.github.com/user/repos";
    let repos: Vec<GitHubRepo> = client
        .get(repos_url)
        .header("Authorization", format!("{} {}", credentials.token_type, credentials.access_token))
        .header("User-Agent", "Axum-Rust-App")
        .send()
        .await
        .expect("Failed to fetch repositories from GitHub")
        .json()
        .await
        .expect("Failed to parse repositories response");

    // Persist repositories in the database
    for repo in repos {
        sqlx::query!(
            r#"
            INSERT INTO repositories (id, name, description, url, created_at, updated_at, owner, private)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
            ON CONFLICT (id) DO UPDATE
            SET name = EXCLUDED.name,
                description = EXCLUDED.description,
                url = EXCLUDED.url,
                created_at = EXCLUDED.created_at,
                updated_at = EXCLUDED.updated_at,
                owner = EXCLUDED.owner,
                private = EXCLUDED.private
            "#,
            repo.id,
            repo.name,
            repo.description,
            repo.html_url,
            repo.created_at,
            repo.updated_at,
            repo.owner.login,
            repo.private
        )
        .execute(&pool)
        .await
        .expect("Failed to insert/update repository data");
    }

    Json(repos)
}
