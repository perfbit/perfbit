use axum::{extract::Path, response::IntoResponse};
use crate::models::commit::{Commit, Changes};

pub async fn fetch_commit_data(Path((repo_id, branch_name)): Path<(String, String)>) -> impl IntoResponse {
    // Fetch commit data logic here
    let commits = vec![
        Commit {
            commit_hash: "abc123".into(),
            author_id: "1".into(),
            message: "Initial commit".into(),
            timestamp: "2024-07-21T12:34:56Z".into(),
            changes: Changes { added: 10, modified: 2, deleted: 1 },
        },
    ];
    axum::Json(commits)
}
