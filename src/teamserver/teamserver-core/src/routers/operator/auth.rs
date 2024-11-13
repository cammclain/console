// operator auth router is used to handle requests from operators clients for authentication

use axum::{
    Router,
    routing::get,
};

pub fn operator_auth_router() -> Router {
    Router::new()
        .route("/login", get(login))
        .route("/logout", get(logout))
}

async fn login() -> &'static str {
    "login"
}

async fn logout() -> &'static str {
    "logout"
}