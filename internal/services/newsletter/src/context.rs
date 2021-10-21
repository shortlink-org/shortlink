use std::sync::Arc;
use bytes::Bytes;
use hyper::{
    body::to_bytes,
    Body, Request
};
use route_recognizer::Params;
use tokio_postgres::{Client, GenericClient};

#[derive(Debug)]
pub struct Context {
    pub req: Request<Body>,
    pub params: Params,
    body_bytes: Option<Bytes>,
}

impl Context {
    pub fn new(req: Request<Body>, params: Params) -> Context {
        Context {
            req,
            params,
            body_bytes: None,
        }
    }

    pub async fn body_json<T: serde::de::DeserializeOwned>(&mut self) -> Result<T, crate::Error> {
        let body_bytes = match self.body_bytes {
            Some(ref v) => v,
            _ => {
                let body = to_bytes(self.req.body_mut()).await?;
                self.body_bytes = Some(body);
                self.body_bytes.as_ref().expect("body_bytes was set above")
            }
        };

        Ok(serde_json::from_slice(&body_bytes)?)
    }
}
