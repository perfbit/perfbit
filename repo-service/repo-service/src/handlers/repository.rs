use axum::response::IntoResponse;
use crate::models::repository::Repository;

pub async fn fetch_repositories() -> impl IntoResponse {
    // Fetch repositories logic here
    let repositories = vec![
        Repository { id: "1".into(), name: "Repo1".into() },
        Repository { id: "2".into(), name: "Repo2".into() },
    ];
    axum::Json(repositories)
}
