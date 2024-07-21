use axum::{extract::Path, extract::Extension, response::IntoResponse, Json};
use crate::models::branch::Branch;
use sqlx::PgPool;

pub async fn fetch_branches(Path(repo_id): Path<i32>, Extension(pool): Extension<PgPool>) -> impl IntoResponse {
    let branches: Vec<Branch> = sqlx::query_as!(
        Branch,
        r#"
        SELECT id, repository_id, name, commit_hash, last_commit_message, last_commit_timestamp, last_commit_author
        FROM branches
        WHERE repository_id = $1
        "#,
        repo_id
    )
    .fetch_all(&pool)
    .await
    .expect("Failed to fetch branches");

    Json(branches)
}
