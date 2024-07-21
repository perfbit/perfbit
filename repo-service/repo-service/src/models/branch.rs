use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Branch {
    pub name: String,
    pub commit_hash: String,
    pub last_commit_message: String,
    pub last_commit_timestamp: String,
    pub last_commit_author: String,
}
