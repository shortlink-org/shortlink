use serde::{Serialize, Deserialize};

#[derive(Deserialize)]
pub(crate) struct SendSubscribeRequest {
    pub(crate) email: String,
    pub(crate) active: bool,
}

#[derive(Serialize, Deserialize)]
pub struct NewsLetter {
     pub(crate) _id: i32,
     pub(crate) email: String
}
