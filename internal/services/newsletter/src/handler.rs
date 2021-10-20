use crate::{Context};
use hyper::{Response, Body, Method, StatusCode};
use crate::domain::SendSubscribeRequest;

pub async fn get_list_subscribes(ctx: Context) -> String {
    format!("get list subscribes: []")
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

    Response::new(
        format!(
            "add newsletter subscribes: {} and active: {}",
            body.email, body.active
        )
        .into(),
    )
}

pub async fn newsletter_unsubscribe(ctx: Context) -> String {
    let param = match ctx.params.find("email") {
        Some(v) => v,
        None => "empty",
    };
    format!("newsletter/unsubscribe: {}", param)
}
