use crate::domain::{NewsLetter, SendSubscribeRequest};
use crate::router::Handler;
use crate::Context;
use crate::{handler, postgres};
use async_trait::async_trait;
use hyper::{Body, Response, StatusCode};
use serde_json::json;

pub struct GetListSubscribesHandler;

#[async_trait]
impl Handler for GetListSubscribesHandler {
    async fn invoke(&self, context: Context) -> Response<Body> {
        let result = handler::get_list_subscribes(context).await;
        Response::new(result.into())
    }
}

pub struct NewsletterSubscribeHandler;

#[async_trait]
impl Handler for NewsletterSubscribeHandler {
    async fn invoke(&self, context: Context) -> Response<Body> {
        handler::newsletter_subscribe(context).await
    }
}

pub struct NewsletterUnsubscribeHandler;

#[async_trait]
impl Handler for NewsletterUnsubscribeHandler {
    async fn invoke(&self, context: Context) -> Response<Body> {
        handler::newsletter_unsubscribe(context).await
    }
}

pub async fn get_list_subscribes(_ctx: Context) -> String {
    match postgres::list().await {
        Ok(newsletters) => {
            serde_json::to_string(&newsletters).unwrap()
        }
        Err(e) => {
            serde_json::to_string(&json!({"error": format!("Database error: {}", e)})).unwrap()
        }
    }
}

pub async fn newsletter_subscribe(mut ctx: Context) -> Response<Body> {
    // Parse the request body
    let body: SendSubscribeRequest = match ctx.body_json().await {
        Ok(v) => v,
        Err(e) => {
            return Response::builder()
                .status(StatusCode::BAD_REQUEST)
                .body(format!("Could not parse JSON: {}", e).into())
                .unwrap();
        }
    };

    // Add email to the database
    if let Err(e) = postgres::add(&body.email).await {
        return Response::builder()
            .status(StatusCode::INTERNAL_SERVER_ERROR)
            .body(format!("Database error: {}", e).into())
            .unwrap();
    }

    // Return success response
    Response::new(
        serde_json::to_string(&json!({
            "message": "Successfully subscribed",
            "email": body.email,
        }))
            .unwrap()
            .into(),
    )
}

pub async fn newsletter_unsubscribe(ctx: Context) -> Response<Body> {
    // Extract the email parameter
    let param = match ctx.params.find("email") {
        Some(email) => email,
        None => {
            return Response::builder()
                .status(StatusCode::BAD_REQUEST)
                .body("Missing email parameter".into())
                .unwrap();
        }
    };

    // Delete the email from the database
    if let Err(e) = postgres::delete(param).await {
        return Response::builder()
            .status(StatusCode::INTERNAL_SERVER_ERROR)
            .body(format!("Database error: {}", e).into())
            .unwrap();
    }

    // Return success response
    Response::new(
        serde_json::to_string(&json!({
            "message": "Successfully unsubscribed",
            "email": param,
        }))
            .unwrap()
            .into(),
    )
}
