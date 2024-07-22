use anyhow::Ok;
use clap::{Arg, Command};
use dotenv::dotenv;



pub fn main() -> anyhow::Result<()>  {

    dotenv().ok();

    let command = Command::new("Dashboard Service Application")
        .version("1.0")
        .author("Maulik Dave <mdave.5191@gmail.com")
        .about("Dashboard microservice for perbit, experiment with Rust-based microservices")
        .arg(
            Arg::new("config")
                .short('c')
                .long("config")
                .help("Configuration file location")
                .default_value("config.json")
        );
    let _matches = command.get_matches();

    Ok(())
}
