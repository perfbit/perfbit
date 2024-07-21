use serde::Serialize;

#[derive(Serialize)]
pub struct Commit {
    pub commit_hash: String,
    pub author_id: String,
    pub message: String,
    pub timestamp: String,
    pub changes: Changes,
}

#[derive(Serialize)]
pub struct Changes {
    pub added: u32,
    pub modified: u32,
    pub deleted: u32,
}
