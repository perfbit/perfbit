use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Commit {
    pub commit_hash: String,
    pub author_id: String,
    pub author_name: String,
    pub author_email: String,
    pub message: String,
    pub timestamp: String,
    pub changes: Changes,
    pub url: String,
    pub repository_id: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Changes {
    pub added: Vec<String>,
    pub modified: Vec<String>,
    pub deleted: Vec<String>,
}
