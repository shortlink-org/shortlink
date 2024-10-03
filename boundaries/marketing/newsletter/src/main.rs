use std::net::SocketAddr;
use tonic::{transport::Server};
use infrastructure::rpc::newsletter::v1::newsletter_service_server::NewsletterServiceServer;
use infrastructure::rpc::newsletter::v1::api::MyNewsletterService;

mod infrastructure;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Define the address for the gRPC server
    let addr = SocketAddr::from(([127, 0, 0, 1], 50051));
    println!("NewsletterService gRPC Server listening on {}", addr);

    // Create the newsletter service
    let newsletter_service = MyNewsletterService::default();

    // Start the gRPC server
    println!("Starting NewsletterService gRPC server on {}", addr);
    Server::builder()
        .add_service(NewsletterServiceServer::new(newsletter_service))
        .serve(addr)
        .await?;

    Ok(())
}
