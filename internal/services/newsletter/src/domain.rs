use serde::Deserialize;

#[derive(Deserialize)]
pub(crate) struct SendSubscribeRequest {
    pub(crate) email: String,
    pub(crate) active: bool,
}
