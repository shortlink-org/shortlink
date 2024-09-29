use crate::context::Context;
use futures::executor::block_on;
use hyper::{
    service::{make_service_fn, service_fn},
    Body, Request, Response, Server,
};
use router::Router;
use std::net::SocketAddr;
use std::sync::Arc;

use pyroscope::PyroscopeAgent;
use pyroscope_pprofrs::{pprof_backend, PprofConfig};

mod context;
mod domain;
mod handler;
mod postgres;
mod router;

type Error = Box<dyn std::error::Error + Send + Sync + 'static>;

#[tokio::main]
pub async fn main() {
    pretty_env_logger::init();

    // Init postgres
    let future = postgres::run_migrations();
    block_on(future);

    // Create Pyroscope Agent
    // TODO: Use env variable for Pyroscope server
    let agent = PyroscopeAgent::builder("http://localhost:4040", "newsletter")
        .backend(pprof_backend(PprofConfig::new().sample_rate(100)))
        .build()?;

    // Start Agent
    let agent_running = agent.start()?;

    // Routing
    let mut router: Router = Router::new();
    router.get("/api/newsletters", Box::new(handler::get_list_subscribes));
    router.post("/api/newsletter", Box::new(handler::newsletter_subscribe));
    router.get(
        "/api/newsletter/unsubscribe/:email",
        Box::new(handler::newsletter_unsubscribe),
    );

    let shared_router = Arc::new(router);
    let new_service = make_service_fn(move |_| {
        let router_capture = shared_router.clone();
        async { Ok::<_, Error>(service_fn(move |req| route(router_capture.clone(), req))) }
    });

    // We'll bind to 127.0.0.1:7070
    let addr = SocketAddr::from(([127, 0, 0, 1], 7070));
    let server = Server::bind(&addr).serve(new_service);
    println!("Listening on http://{}", addr);

    // And now add a graceful shutdown signal...
    let graceful = server.with_graceful_shutdown(shutdown_signal());

    // Run this server for... forever!
    if let Err(e) = graceful.await {
        eprintln!("server error: {}", e);

        // Stop Agent
        let agent_ready = agent_running.stop()?;
        agent_ready.shutdown();
    }
}

async fn route(router: Arc<Router>, req: Request<hyper::Body>) -> Result<Response<Body>, Error> {
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
