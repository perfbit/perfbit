use axum::{
    routing::{get, post},
    Router,
};
use crate::handlers::{git::connect_git_account, repository::fetch_repositories, branch::fetch_branches, commit::fetch_commit_data};

pub fn create_routes() -> Router {
    Router::new()
        .route("/api/v1/git/connect", post(connect_git_account))
        .route("/api/v1/repositories", get(fetch_repositories))
        .route("/api/v1/repositories/:repoId/branches", get(fetch_branches))
        .route("/api/v1/repositories/:repoId/branches/:branchName/commits", get(fetch_commit_data))
}
