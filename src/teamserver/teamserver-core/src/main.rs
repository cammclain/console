use axum::{
    //routing::{get, post},
    Router, Json,
};

// use serde for deserialization and serialization
use serde::{Deserialize, Serialize};

#[tokio::main]
async fn main() {
    // initialize tracing within the application
    tracing_subscriber::fmt::init();
    

    // initialize the database
    // TODO: implement database initialization

    // build the application
    let app = Router::new().
    
}
