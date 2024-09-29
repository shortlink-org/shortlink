use crate::domain::{NewsLetter, SendSubscribeRequest};
use crate::router::Handler;
use crate::Context;
use crate::{handler, postgres};
use async_trait::async_trait;
use hyper::{Body, Response, StatusCode};

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
    let newsletters = postgres::list().await.unwrap();
    serde_json::to_string(&newsletters).unwrap().to_string()
}

pub async fn newsletter_subscribe(mut ctx: Context) -> Response<Body> {
    let body: SendSubscribeRequest = match ctx.body_json().await {
        Ok(v) => v,
        Err(e) => {
            return hyper::Response::builder()
                .status(StatusCode::BAD_REQUEST)
                .body(format!("could not parse JSON: {}", e).into())
                .unwrap()
        }
    };

    postgres::add(&body.email).await.unwrap();

    Response::new(
        serde_json::to_string(&NewsLetter {
                _id: 0,
                email: body.email,
            })
            .unwrap().to_string()
        .into(),
    )
}

pub async fn newsletter_unsubscribe(ctx: Context) -> Response<Body> {
    let param = ctx.params.find("email").unwrap_or("empty");

    postgres::delete(param).await.unwrap();

    Response::new(
        serde_json::to_string(&NewsLetter {
                _id: 0,
                email: param.into(),
            })
            .unwrap().to_string()
        .into(),
    )
}
