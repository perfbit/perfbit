use axum::{
    routing::{get, post},
    Router,
};
use std::net::SocketAddr;

mod handlers;
mod models;
mod routes;

use std::net::SocketAddr;
use routes::create_routes;

#[tokio::main]
async fn main() {
    let app = create_routes();

    let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
    println!("Listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}