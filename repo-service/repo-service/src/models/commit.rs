use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Serialize, Deserialize, FromRow, Debug)]
pub struct Commit {
    pub id: i32,
    pub repository_id: i32,
    pub branch_id: i32,
    pub commit_hash: String,
    pub author_id: String,
    pub author_name: String,
    pub author_email: String,
    pub message: String,
    pub timestamp: String,
    pub url: String,
    pub changes: Changes,
}

#[derive(Serialize, Deserialize, sqlx::Type, Debug)]
pub struct Changes {
    pub added: Vec<String>,
    pub modified: Vec<String>,
    pub deleted: Vec<String>,
}
