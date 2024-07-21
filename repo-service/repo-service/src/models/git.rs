use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug)]
pub struct ConnectGitRequest {
    pub provider: String,
    pub access_token: String,
    pub user_id: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct GitProvider {
    pub provider: String,
    pub base_url: String,
    pub access_token: String,
    pub refresh_token: Option<String>,
    pub expires_at: Option<String>,
}
