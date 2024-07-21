use serde::Serialize;

#[derive(Serialize)]
pub struct Repository {
    pub id: String,
    pub name: String,
}
