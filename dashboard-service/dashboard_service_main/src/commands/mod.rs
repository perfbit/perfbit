use clap::{ArgMathces, Command};

pub fn configuration(command: Command) -> Command {
    command.subcommand(Command::new("Hello").about("Hello World"))
}

pub fn handle(matches: &ArgMatches) -> anyhow::Result<()> {
    if let Some(_matches) = matches.subcommand_matches("hello") {
        println!("Hello World!");
    }

    Ok(())
}