mod handlers;
mod models;
mod routes;

use dotenv::dotenv;
use sqlx::postgres::PgPoolOptions;
use std::net::SocketAddr;
use routes::create_routes;
use reqwest::Client;
use axum::AddExtensionLayer;

#[tokio::main]
async fn main() {
    dotenv().ok();

    let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await
        .expect("Failed to create pool.");

    let client = Client::new();

    let app = create_routes()
        .layer(AddExtensionLayer::new(pool))
        .layer(AddExtensionLayer::new(client));

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("Listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}
