### Newsletter

The service sends messages to everyone who has subscribed to the newsletter.

### Use

- Rust
- HTTP API
- [refinery](https://github.com/rust-db/refinery) is used to setup database migrations

### Roadmap

See  this [project](https://github.com/batazor/shortlink/projects/20)

### Running the Binaries

- `migrate` (which can be run from cargo run --bin migrate) is the binary to execute database migrations

If you're building this, please set your environment configuration from .env file (copied from .env.example). 
Then you can run cargo run --bin migrate to create the table.
