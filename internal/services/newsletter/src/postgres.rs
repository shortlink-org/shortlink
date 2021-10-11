use postgres::{Client, NoTls, Error};
use std::env;
use std::collections::HashMap;
use postgres::types::Kind::Array;

struct NewsLetter {
    _id: i32,
    email: String
}

fn new() -> Result<(), Error> {
    let postgres_uri = env::var("STORE_POSTGRES_URI").unwrap();

    let mut client = Client::connect(&postgres_uri, NoTls)?;
    Ok(())
}

pub async fn list() -> Result<(Vec<String>), Error> {
    let mut newsletters = Vec::new();

    client.batch_execute("
        SELECT * FROM shortlink.newsletter
    ")?;

    for row in client.query("SELECT id, email FROM shortlink.newsletter", &[])? {
        let author = NewsLetter {
            _id: row.get(0),
            email: row.get(1),
        };
        newsletters.push(author.email);
    }

    Ok(newsletters)
}

pub async fn add(email: &str) -> Result<(), Error> {
    let newsletter = NewsLetter {
        _id: 0,
        email: email.to_string(),
    };

    client.batch_execute(
        "INSERT INTO shortlink.newsletter (email) VALUES ($1)",
        &newsletter
    )?;

    Ok(())
}

pub async fn delete() -> Result<(), Error> {
    client.batch_execute("
        SELECT * FROM shortlink.newsletter
    ")?;

    Ok(())
}
