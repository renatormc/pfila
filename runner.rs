use std::fs::File;
use std::io::Write;
use std::path::Path;
use std::process::{Command, Stdio};

fn main() {
    let args = vec![
        "D:\\tests\\pfila\\iped\\iped-4.1.2\\jre\\bin\\java.exe",
        "-jar",
        "D:\\tests\\pfila\\iped\\iped-4.1.2\\iped.jar",
        "-profile",
        "fastmode",
        "-d",
        "D:\\tests\\pfila\\pen.E01",
        "-o",
        "D:\\tests\\pfila\\result",
        "--nogui",
    ];
    let path = Path::new("D:\\tests\\pfila\\console\\output.txt");
    let mut file = File::create(&path).expect("Failed to create output file");

    let mut cmd = Command::new(&args[0]);
    cmd.args(&args[1..])
        .stdout(Stdio::from(file.try_clone().expect("Failed to clone file handle")))
        .stderr(Stdio::from(file.try_clone().expect("Failed to clone file handle")));

    match cmd.spawn() {
        Ok(mut child) => {
            if let Err(err) = child.wait() {
                eprintln!("Failed to wait for child process: {}", err);
            }
        }
        Err(err) => {
            eprintln!("Failed to execute command: {}", err);
        }
    }
}
