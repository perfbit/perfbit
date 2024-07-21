use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
pub struct ConnectGitRequest {
    pub provider: String,
    pub access_token: String,
}
