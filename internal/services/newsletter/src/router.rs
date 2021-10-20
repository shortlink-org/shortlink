use crate::{Context};
use async_trait::async_trait;
use futures::future::Future;
use hyper::{Response, Body, Method, StatusCode};
use route_recognizer::{Match, Params, Router as InternalRouter};
use std::collections::HashMap;

#[async_trait]
pub trait Handler: Send + Sync + 'static {
    async fn invoke(&self, context: Context) -> Response<Body>;
}

#[async_trait]
impl<F: Send + Sync + 'static, Fut> Handler for F
where
    F: Fn(Context) -> Fut,
    Fut: Future + Send + 'static,
    Fut::Output: IntoResponse,
{
    async fn invoke(&self, context: Context) -> Response<Body> {
        (self)(context).await.into_response()
    }
}

pub struct RouterMatch<'a> {
    pub handler: &'a dyn Handler,
    pub params: Params,
}

pub struct Router {
    method_map: HashMap<Method, InternalRouter<Box<dyn Handler>>>,
}

impl Router {
    pub fn new() -> Router {
        Router {
            method_map: HashMap::default(),
        }
    }

    pub fn get(&mut self, path: &str, handler: Box<dyn Handler>) {
        self.method_map
            .entry(Method::GET)
            .or_insert_with(InternalRouter::new)
            .add(path, handler)
    }

    pub fn post(&mut self, path: &str, handler: Box<dyn Handler>) {
        self.method_map
            .entry(Method::POST)
            .or_insert_with(InternalRouter::new)
            .add(path, handler)
    }

    pub fn delete(&mut self, path: &str, handler: Box<dyn Handler>) {
        self.method_map
            .entry(Method::DELETE)
            .or_insert_with(InternalRouter::new)
            .add(path, handler)
    }

    pub fn route(&self, path: &str, method: &Method) -> RouterMatch<'_> {
        if let Some(Match { handler, params }) = self
            .method_map
            .get(method)
            .and_then(|r| r.recognize(path).ok())
        {
            RouterMatch {
                handler: &**handler,
                params,
            }
        } else {
            RouterMatch {
                handler: &not_found_handler,
                params: Params::new(),
            }
        }
    }
}

async fn not_found_handler(_cx: Context) -> Response<Body> {
    hyper::Response::builder()
        .status(StatusCode::NOT_FOUND)
        .body("NOT FOUND".into())
        .unwrap()
}

pub trait IntoResponse: Send + Sized {
    fn into_response(self) -> Response<Body>;
}

impl IntoResponse for Response<Body> {
    fn into_response(self) -> Response<Body> {
        self
    }
}

impl IntoResponse for &'static str {
    fn into_response(self) -> Response<Body> {
        Response::new(self.into())
    }
}

impl IntoResponse for String {
    fn into_response(self) -> Response<Body> {
        Response::new(self.into())
    }
}
