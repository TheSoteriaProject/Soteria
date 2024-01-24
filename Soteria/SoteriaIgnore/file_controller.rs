use std::fs;
use std::env;
use std::process;
use std::path::Path;

fn folder_traverse(dir: &Path) {
    if dir.is_dir() {
        for entry in fs::read_dir(dir).unwrap() {
            let entry = entry.unwrap();
            let path = entry.path();

            if path.is_file() {
                // If it's a file (not a directory), print or store information
                println!("{}", path.display());
            } else if path.is_dir() {
                // Recursively visit subdirectories
                folder_traverse(&path);
            }
        }
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        println!("Files");
        println!("---------------------");
        for arg in args.iter().skip(2) { // Skip 2 because of directory name and also program name.
            println!("{}", arg);
        }
    } else {
        process::exit(1);
    }

    let root_dir = args[1].clone();
    let root_path = Path::new(&root_dir);

    folder_traverse(&root_path);
}
