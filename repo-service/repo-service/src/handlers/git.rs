use axum::{extract::Json, http::StatusCode, response::IntoResponse};
use crate::models::git::ConnectGitRequest;

pub async fn connect_git_account(Json(payload): Json<ConnectGitRequest>) -> impl IntoResponse {
    // OAuth2 connection logic here
    (StatusCode::OK, Json("Git account connected successfully"))
}
