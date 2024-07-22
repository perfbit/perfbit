use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Serialize, Deserialize, FromRow, Debug)]
pub struct Branch {
    pub id: i32,
    pub repository_id: i32,
    pub name: String,
    pub commit_hash: String,
    pub last_commit_message: String,
    pub last_commit_timestamp: String,
    pub last_commit_author: String,
}
