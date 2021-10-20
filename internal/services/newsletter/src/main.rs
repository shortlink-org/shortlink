use bytes::Bytes;
use std::{net::SocketAddr};

use hyper::{
    body::to_bytes,
    service::{make_service_fn, service_fn},
    Body, Request, Server,
};
use route_recognizer::Params;
use router::Router;
use std::sync::Arc;
use crate::context::Context;

mod handler;
mod router;
mod context;

type Response = hyper::Response<hyper::Body>;
type Error = Box<dyn std::error::Error + Send + Sync + 'static>;

#[tokio::main]
pub async fn main() {
    pretty_env_logger::init();

    let mut router: Router = Router::new();
    router.get("/api/newsletters", Box::new(handler::get_list_subscribes));
    router.post("/api/newsletter", Box::new(handler::newsletter_subscribe));
    router.delete("/api/newsletter/unsubscribe/:email", Box::new(handler::newsletter_unsubscribe));

    let shared_router = Arc::new(router);
    let new_service = make_service_fn(move |_| {
        let router_capture = shared_router.clone();
        async {
            Ok::<_, Error>(service_fn(move |req| {
                route(router_capture.clone(), req)
            }))
        }
    });

    // We'll bind to 127.0.0.1:3000
    let addr = SocketAddr::from(([127, 0, 0, 1], 7070));
    let server = Server::bind(&addr).serve(new_service);
    println!("Listening on http://{}", addr);

    // And now add a graceful shutdown signal...
    let graceful = server.with_graceful_shutdown(shutdown_signal());

    // Run this server for... forever!
    if let Err(e) = graceful.await {
        eprintln!("server error: {}", e);
    }
}

async fn route(
    router: Arc<Router>,
    req: Request<hyper::Body>,
) -> Result<Response, Error> {
    let found_handler = router.route(req.uri().path(), req.method());
    let resp = found_handler
        .handler
        .invoke(Context::new(req, found_handler.params))
        .await;
    Ok(resp)
}

async fn shutdown_signal() {
    // Wait for the CTRL+C signal
    tokio::signal::ctrl_c()
        .await
        .expect("failed to install CTRL+C signal handler");
}

