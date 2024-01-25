use std::fs;
use std::env;
use std::process;
use std::path::Path;

fn is_dockerfile(file_path: &str) -> bool {
    // Check if file extension is Dockerfile
    if let Some(extension) = Path::new(file_path).extension() {
        return extension.to_string_lossy().eq_ignore_ascii_case("Dockerfile");
    }
    // Check if file name is Dockerfile
    if let Some(file_name) = Path::new(file_path).file_name() {
        return file_name.to_string_lossy().eq_ignore_ascii_case("Dockerfile");
    }
    false
}

fn is_makefile(file_path: &str) -> bool {
    // Check if file extension is Makefile
    if let Some(extension) = Path::new(file_path).extension() {
        return extension.to_string_lossy().eq_ignore_ascii_case("Makefile");
    }
    // Check if file name is Makefile
    if let Some(file_name) = Path::new(file_path).file_name() {
        return file_name.to_string_lossy().eq_ignore_ascii_case("Makefile");
    }
    false
}

fn is_bourne_shell_script(file_path: &str) -> bool {
    // Check if file extension is Bourne Shell
    if let Some(extension) = Path::new(file_path).extension() {
        let extension_str = extension.to_string_lossy();
        return extension_str.eq_ignore_ascii_case("sh") || extension_str.eq_ignore_ascii_case("bash");
    }
    false
}

/* 
fn is_file_ignored() {}
*/

fn folder_traverse(dir: &Path) {
    if dir.is_dir() {
        for entry in fs::read_dir(dir).unwrap() {
            let entry = entry.unwrap();
            let path = entry.path();

            if path.is_file() {
                // If it's a file (not a directory) and is also one of the correct extensions
                if let Some(path_str) = path.to_str() {
                    if is_dockerfile(path_str) {
                        println!("Docker File Path: {}", path.display());
                    } else if is_makefile(path_str) {
                        println!("Makefile File Path: {}", path.display());
                    } else if is_bourne_shell_script(path_str) {
                        println!("Bourne Shell File Path: {}", path.display());
                    } else {
                        // Nothing
                    }
                }
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
