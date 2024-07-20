use actix_web::{web, HttpResponse, Responder};
use diesel::prelude::*;
use crate::{models::{User, NewUser}, schema::users::dsl::*};
use crate::AppState;

pub fn init_routes(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::scope("/api")
            .route("/users", web::get().to(get_users))
            .route("/users", web::post().to(create_user))
    );
}

async fn get_users(data: web::Data<AppState>) -> impl Responder {
    use crate::schema::users::dsl::*;

    let connection = &data.pool;

    let results = users
        .load::<User>(connection)
        .expect("Error loading users");

    HttpResponse::Ok().json(results)
}

async fn create_user(data: web::Data<AppState>, new_user: web::Json<NewUser>) -> impl Responder {
    use crate::schema::users;

    let connection = &data.pool;

    let new_user = NewUser {
        username: &new_user.username,
        email: &new_user.email,
    };

    diesel::insert_into(users::table)
        .values(&new_user)
        .execute(connection)
        .expect("Error saving new user");

    HttpResponse::Created().finish()
}
