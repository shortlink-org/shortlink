use tonic::{Request, Response, Status};
use async_trait::async_trait;

#[derive(Default)]
pub struct MyNewsletterService;

// Implement the generated NewsletterService trait
#[async_trait]
impl NewsletterService for MyNewsletterService {
    async fn get(
        &self,
        request: Request<GetRequest>,
    ) -> Result<Response<GetResponse>, Status> {
        let email = request.into_inner().email;
        let response = GetResponse {
            email: email.clone(),
            active: true,  // Example response
        };

        println!("Received Get request for email: {}", email);

        Ok(Response::new(response))
    }

    async fn subscribe(
        &self,
        request: Request<SubscribeRequest>,
    ) -> Result<Response<()>, Status> {
        let email = request.into_inner().email;

        println!("Subscribed to newsletter: {}", email);

        Ok(Response::new(()))
    }

    async fn un_subscribe(
        &self,
        request: Request<UnSubscribeRequest>,
    ) -> Result<Response<()>, Status> {
        let email = request.into_inner().email;

        println!("Unsubscribed from newsletter: {}", email);

        Ok(Response::new(()))
    }

    async fn list(
        &self,
        _request: Request<()>,
    ) -> Result<Response<ListResponse>, Status> {
        // Example list response
        let newsletters = vec![];

        let response = ListResponse { newsletters };

        println!("Returning list of newsletters");

        Ok(Response::new(response))
    }

    async fn update_status(
        &self,
        request: Request<UpdateStatusRequest>,
    ) -> Result<Response<()>, Status> {
        let emails = request.into_inner().emails;
        let active = request.into_inner().active;

        for email in emails {
            println!("Updated status for email: {}, active: {}", email, active);
        }

        Ok(Response::new(()))
    }

    async fn delete(
        &self,
        request: Request<DeleteRequest>,
    ) -> Result<Response<()>, Status> {
        let emails = request.into_inner().emails;

        for email in emails {
            println!("Deleted newsletter for email: {}", email);
        }

        Ok(Response::new(()))
    }
}
