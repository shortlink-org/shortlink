use std::env;
use tokio_postgres::{NoTls, Client, GenericClient};

type Error = Box<dyn std::error::Error + Send + Sync + 'static>;

struct NewsLetter {
    _id: i32,
    email: String
}

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

pub async fn list() -> std::result::Result<(Vec<String>), Error> {
    let mut newsletters = Vec::new();

    // client.batch_execute("
    //     SELECT * FROM shortlink.newsletter
    // ")?;
    //
    // for row in client.query("SELECT id, email FROM shortlink.newsletter", &[])? {
    //     let author = NewsLetter {
    //         _id: row.get(0),
    //         email: row.get(1),
    //     };
    //     newsletters.push(author.email);
    // }

    Ok(newsletters)
}

pub async fn add(email: &str) -> std::result::Result<(), Error> {
    // let newsletter = NewsLetter {
    //     _id: 0,
    //     email: email.to_string(),
    // };
    //
    // client.batch_execute(
    //     "INSERT INTO shortlink.newsletter (email) VALUES ($1)",
    //     &newsletter
    // )?;

    Ok(())
}

pub async fn delete() -> std::result::Result<(), Error> {
    // client.batch_execute("
    //     SELECT * FROM shortlink.newsletter
    // ")?;

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
