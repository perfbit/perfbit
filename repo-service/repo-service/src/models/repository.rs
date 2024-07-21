use serde::{Deserialize, Serialize};
use sqlx::FromRow;

#[derive(Serialize, Deserialize, FromRow, Debug)]
pub struct Repository {
    pub id: i32,
    pub name: String,
    pub description: Option<String>,
    pub url: String,
    pub created_at: String,
    pub updated_at: String,
    pub owner: String,
    pub private: bool,
}
