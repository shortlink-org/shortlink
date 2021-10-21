use crate::{Context};
use hyper::{Response, Body, Method, StatusCode};
use crate::domain::SendSubscribeRequest;
use crate::postgres;
use futures::executor::block_on;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
struct NewsLetter {
    _id: i32,
    email: String
}

pub async fn get_list_subscribes(ctx: Context) -> String {
    let mut client = postgres::new().await.unwrap();

    let rows = client.query("SELECT id, email FROM shortlink.newsletters", &[]).await;

    let mut newsletters = Vec::new();

    for row in rows.unwrap().as_slice() {
        newsletters.push(NewsLetter{
            _id: 0,
            email: row.get(1),
        });
    }

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

    let mut client = postgres::new().await.unwrap();
    client.execute("INSERT INTO shortlink.newsletters (email) VALUES ($1)", &[&body.email]).await.ok();

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

    let mut client = postgres::new().await.unwrap();
    client.execute("DELETE FROM shortlink.newsletters WHERE email=$1", &[&param]).await.ok();

    Response::new(format!("{}", serde_json::to_string(&NewsLetter{
        _id: 0,
        email: param.into(),
    }).unwrap()).into())
}
