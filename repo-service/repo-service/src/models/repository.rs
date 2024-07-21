use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Repository {
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    pub url: String,
    pub created_at: String,
    pub updated_at: String,
    pub owner: String,
    pub private: bool,
}
