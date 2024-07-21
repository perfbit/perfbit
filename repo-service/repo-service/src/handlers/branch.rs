use axum::{extract::Path, response::IntoResponse};
use crate::models::branch::Branch;

pub async fn fetch_branches(Path(repo_id): Path<String>) -> impl IntoResponse {
    // Fetch branches logic here
    let branches = vec![
        Branch { name: "main".into() },
        Branch { name: "develop".into() },
    ];
    axum::Json(branches)
}
