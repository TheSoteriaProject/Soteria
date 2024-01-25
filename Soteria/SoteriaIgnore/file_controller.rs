use std::fs;
use std::env;
use std::process;
use std::path::Path;
use std::fs::read_to_string;

fn remove_comments_and_empty_lines(lines: Vec<String>) -> Vec<String> {
    let mut result = Vec::new();

    for line in lines {
        let line_without_comment: String = line.splitn(2, '#').next().unwrap_or("").trim().to_string();

        if !line_without_comment.is_empty() {
            result.push(line_without_comment);
        }
    }

    result
}

fn tokenize_lines(lines: &Vec<String>) -> Vec<String> {
    let mut token_results = Vec::new();

    for line in lines {
        if line.starts_with("-") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : SkipFolder");
            }
            token_results.push(line.to_string() + " : SkipFile");
        }

        if line.starts_with("+") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : AddFolder");
            }
            token_results.push(line.to_string() + " : AddFile");
        }

        if line.starts_with("*") {
            token_results.push(line.to_string() + " : SkipExtension");
        }
    }

    token_results
}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();
    let mut skip_multiline_comment = false;
    for line in read_to_string(filename).unwrap().lines() {
       
        if let Some(first_char) = line.chars().next() {
            if first_char == '#' {
                continue;
            }

            if line.starts_with("|-start") {
                skip_multiline_comment = true;
                continue;
            }

            if skip_multiline_comment && line.starts_with("|-end") {
                skip_multiline_comment = false;
                continue;
            }

            if skip_multiline_comment {
                continue;
            }
        }

        result.push(line.to_string())
    }
    
    let next_result = remove_comments_and_empty_lines(result);

    let next_next_result = tokenize_lines(&next_result);

    next_next_result
}

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

    /**************************
    ** Gather Special Cases **
    **************************/
    let _ignore_file = "./.soteriaignore";
    let content = read_lines(_ignore_file);
    println!("Content:\n{:?}", content);

    /**********************
    ** Gather File List **
    *********************/
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        println!("Files");
        println!("---------------------");
        /*for arg in args.iter().skip(2) { // Skip 2 because of directory name and also program name.
            println!("{}", arg);
        } */
    } else {
        process::exit(1);
    }

    let root_dir = args[1].clone();
    let root_path = Path::new(&root_dir);

    folder_traverse(&root_path);
}
