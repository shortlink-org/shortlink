use crate::{Context};
use hyper::{Response, Body, StatusCode};
use crate::domain::{NewsLetter, SendSubscribeRequest};
use crate::postgres;

pub async fn get_list_subscribes(ctx: Context) -> String {
    let newsletters = postgres::list().await.unwrap();
    format!("{}", serde_json::to_string(&newsletters).unwrap())
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

    Response::new(format!("{}", serde_json::to_string(&NewsLetter{
        _id: 0,
        email: body.email,
    }).unwrap()).into())
}

pub async fn newsletter_unsubscribe(ctx: Context) -> Response<Body> {
    let param = match ctx.params.find("email") {
        Some(v) => v,
        None => "empty",
    };

    postgres::delete(param).await.unwrap();

    Response::new(format!("{}", serde_json::to_string(&NewsLetter{
        _id: 0,
        email: param.into(),
    }).unwrap()).into())
}
