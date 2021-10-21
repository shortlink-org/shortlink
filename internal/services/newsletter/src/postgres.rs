use std::env;
use tokio_postgres::{NoTls, Client, GenericClient};
use crate::domain::NewsLetter;

type Error = Box<dyn std::error::Error + Send + Sync + 'static>;

pub async fn new() -> Result<Client, Error> {
    let postgres_uri = env::var("STORE_POSTGRES_URI").unwrap();

    // Connect to the database.
    let (client, connection) =
        tokio_postgres::connect(&postgres_uri, NoTls).await?;

    // The connection object performs the actual communication with the database,
    // so spawn it off to run on its own.
    tokio::spawn(async move {
        if let Err(e) = connection.await {
            eprintln!("connection error: {}", e);
        }
    });

    Ok(client)
}

pub async fn list() -> std::result::Result<Vec<NewsLetter>, Error> {
    let mut client = new().await.unwrap();

    let rows = client.query("SELECT id, email FROM shortlink.newsletters", &[]).await;

    let mut newsletters = Vec::new();
    for row in rows.unwrap().as_slice() {
        newsletters.push(NewsLetter{
            _id: 0,
            email: row.get(1),
        });
    }

    Ok(newsletters)
}

pub async fn add(email: &str) -> std::result::Result<(), Error> {
    let mut client = new().await.unwrap();
    client.execute("INSERT INTO shortlink.newsletters (email) VALUES ($1)", &[&email]).await.ok();

    Ok(())
}

pub async fn delete(email: &str) -> std::result::Result<(), Error> {
    let mut client = new().await.unwrap();
    client.execute("DELETE FROM shortlink.newsletters WHERE email=$1", &[&email]).await.ok();

    Ok(())
}

mod embedded {
    use refinery::embed_migrations;
    embed_migrations!("src/migrations");
}

pub(crate) async fn run_migrations() -> Result<Client, Error> {
    let mut client = new().await.unwrap();

    println!("Running DB migrations...");

    let migration_report = embedded::migrations::runner()
        .run_async(&mut client)
        .await?;

    for migration in migration_report.applied_migrations() {
        println!(
            "Migration Applied -  Name: {}, Version: {}",
            migration.name(),
            migration.version()
        );
    }
    println!("DB migrations finished!");

    Ok(client)
}
