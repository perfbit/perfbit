use axum::{extract::Path, extract::Extension, response::IntoResponse, Json};
use crate::models::commit::{Commit, Changes};
use sqlx::PgPool;

pub async fn fetch_commit_data(Path((repo_id, branch_name)): Path<(i32, String)>, Extension(pool): Extension<PgPool>) -> impl IntoResponse {
    let commits: Vec<Commit> = sqlx::query_as!(
        Commit,
        r#"
        SELECT 
            id, 
            repository_id, 
            branch_id, 
            commit_hash, 
            author_id, 
            author_name, 
            author_email, 
            message, 
            timestamp, 
            url, 
            jsonb_build_object('added', changes->'added', 'modified', changes->'modified', 'deleted', changes->'deleted') as changes
        FROM commits
        WHERE repository_id = $1 AND branch_id = (
            SELECT id FROM branches WHERE name = $2 AND repository_id = $1
        )
        "#,
        repo_id,
        branch_name
    )
    .fetch_all(&pool)
    .await
    .expect("Failed to fetch commits");

    Json(commits)
}
