#[derive(Clone, PartialEq, Debug)]
pub struct Newsletter {
    // The unique identifier of the newsletter.
    pub email: String,

    // Status of the newsletter.
    pub active: bool,
}

#[derive(Clone, PartialEq, Debug)]
pub struct Newsletters {
    // List of newsletters.
    pub list: Vec<Newsletter>,
}
