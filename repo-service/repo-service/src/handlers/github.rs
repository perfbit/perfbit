use axum::{
    extract::{Extension, Json},
    response::IntoResponse,
    http::StatusCode,
};
use serde::{Deserialize, Serialize};
use sqlx::PgPool;
use reqwest::Client;
use std::env;

#[derive(Deserialize)]
pub struct GitHubCallbackRequest {
    pub code: String,
    pub state: String,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct GitHubAccessTokenResponse {
    access_token: String,
    scope: String,
    token_type: String,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct GitHubUser {
    login: String,
    id: i32,
    node_id: String,
    avatar_url: String,
    url: String,
    name: Option<String>,
    company: Option<String>,
    blog: Option<String>,
    location: Option<String>,
    email: Option<String>,
}

pub async fn handle_github_callback(
    Json(payload): Json<GitHubCallbackRequest>,
    Extension(pool): Extension<PgPool>,
    Extension(client): Extension<Client>,
) -> impl IntoResponse {
    let client_id = env::var("GITHUB_CLIENT_ID").expect("GITHUB_CLIENT_ID must be set");
    let client_secret = env::var("GITHUB_CLIENT_SECRET").expect("GITHUB_CLIENT_SECRET must be set");

    // Exchange code for access token
    let token_res = client.post("https://github.com/login/oauth/access_token")
        .header("Accept", "application/json")
        .form(&[
            ("client_id", client_id.as_str()),
            ("client_secret", client_secret.as_str()),
            ("code", payload.code.as_str()),
            ("state", payload.state.as_str()),
        ])
        .send()
        .await
        .expect("Failed to send request for access token")
        .json::<GitHubAccessTokenResponse>()
        .await
        .expect("Failed to parse access token response");

    // Use the access token to fetch user data
    let user_res = client.get("https://api.github.com/user")
        .header("Authorization", format!("token {}", token_res.access_token))
        .header("User-Agent", "Axum-Rust-App")
        .send()
        .await
        .expect("Failed to send request for user data")
        .json::<GitHubUser>()
        .await
        .expect("Failed to parse user data");

    // Persist user data in the database
    sqlx::query!(
        r#"
        INSERT INTO github_users (login, github_id, node_id, avatar_url, url, name, company, blog, location, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        ON CONFLICT (github_id) DO UPDATE
        SET login = EXCLUDED.login,
            node_id = EXCLUDED.node_id,
            avatar_url = EXCLUDED.avatar_url,
            url = EXCLUDED.url,
            name = EXCLUDED.name,
            company = EXCLUDED.company,
            blog = EXCLUDED.blog,
            location = EXCLUDED.location,
            email = EXCLUDED.email
        "#,
        user_res.login,
        user_res.id,
        user_res.node_id,
        user_res.avatar_url,
        user_res.url,
        user_res.name,
        user_res.company,
        user_res.blog,
        user_res.location,
        user_res.email
    )
    .execute(&pool)
    .await
    .expect("Failed to insert/update user data");

    (StatusCode::OK, Json(user_res))
}
