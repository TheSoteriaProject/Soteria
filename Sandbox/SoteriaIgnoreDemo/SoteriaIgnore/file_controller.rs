use std::fs;
use std::env;
use std::process;
use std::path::Path;
use std::fs::read_to_string;

// Ignore File Comment Parser
fn remove_comments_and_empty_lines(lines: Vec<String>) -> Vec<String> {
    let mut result = Vec::new();

    for line in lines {
        // Removes comments after statement
        let line_without_comment: String = line.splitn(2, '#').next().unwrap_or("").trim().to_string();

        if !line_without_comment.is_empty() {
            result.push(line_without_comment);
        }
    }

    result
}

// Tokenize
fn tokenize_lines(lines: &Vec<String>) -> Vec<String> {
    let mut token_results = Vec::new();

    // Adjust so it removes or keeps
    for line in lines {
        if line.starts_with("-") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : SkipFolder");
                continue;
            }
            token_results.push(line.to_string() + " : SkipFile");
        }

        if line.starts_with("+") {
            if line.ends_with("/") {
                token_results.push(line.to_string() + " : IncludeFolder");
                continue;
            }
            token_results.push(line.to_string() + " : IncludeFile");
        }

        if line.starts_with("*") {
            token_results.push(line.to_string() + " : SkipExtension");
        }
    }

    token_results
}

// Read Lines of File
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



fn cmp_files(pre_files: Vec<String>, special_files: Vec<String>) -> Vec<String>{
    let mut final_files = Vec::new();
    
    for pfile in pre_files {
        for sfile in &special_files {
            // println!("{} <-> {}", pfile, sfile);
            // Split the string into parts based on ":"
            let split_token: Vec<&str> = sfile.split(":").collect();

            // Use the nth method to get the second part (index 1)
            if let Some(pre_token) = split_token.get(1) {
                // REWRITE this because its awful coding
                let token_str: &str = *pre_token;
                // Trim leading and trailing whitespaces
                let token: &str = token_str.trim();
                if token == "SkipFile" {
                    if let Some(pre_s_file) = split_token.get(0) {
                        let pre_s_file_str: &str = *pre_s_file;
                        let file_s_bad_character: &str = pre_s_file_str.trim_start_matches("-");
                        let file_s: &str = file_s_bad_character.trim();
                        if file_s == pfile {
                            println!("Skip File: {} <-> {}", file_s, pfile);
                        }
                    }
                } else if token == "SkipFolder" {
                    if let Some(_folder) = split_token.get(0) { 
                        /* 
                        let pre_file_str: &str = *pre_file;                             
                        let file_bad_character: &str = pre_file_str.trim_start_matches("-");
                        let file: &str = file_bad_character.trim();
                        // Deconstruct Path
                        if file == pfile {
                            println!("Skip Folder: {} <-> {}", file, pfile);                        
                        } */
                    }
                } else if token == "IncludeFile" {
                    if let Some(pre_i_file) = split_token.get(0) { 
                        let pre_i_file_str: &str = *pre_i_file;                             
                        let file_i_bad_character: &str = pre_i_file_str.trim_start_matches("+");
                        let file_i: &str = file_i_bad_character.trim();
                        if file_i == pfile {
                            println!("Include File: {} <-> {}", file_i, pfile);
                        }
                    }
                } else if token == "IncludeFolder" {
                    if let Some(_folder) = split_token.get(0) { 
                        /*
                        let pre_file_str: &str = *pre_file;                             
                        let file_bad_character: &str = pre_file_str.trim_start_matches("-");
                        let file: &str = file_bad_character.trim();
                        // Deconstruct Path
                        if file == pfile {
                            println!("Skip Files: {} <-> {}", file, pfile);
                        } */
                    }  
                } else if token == "SkipExtension" {
                    if let Some(_file) = split_token.get(0) { 
                        // Remove Extension *
                        // Special Case Write Later
                    }
                } else {
                    // Do Nothing For Now
                }
            } else {
                println!("Invalid format");
            }
        }
    }

    final_files
}

fn folder_traverse(dir: &Path) -> Vec<String> {
    let mut files = Vec::new();

    if dir.is_dir() {
        for entry in fs::read_dir(dir).unwrap() {
            let entry = entry.unwrap();
            let path = entry.path();

            if path.is_file() {
                // If it's a file (not a directory) and is also one of the correct extensions
                if let Some(path_str) = path.to_str() {
                    if is_dockerfile(path_str) { // Bad Coding
                        // println!("Docker File Path: {}", path.display());
                        files.push(path_str.to_owned());
                    } else if is_makefile(path_str) {
                        // println!("Makefile File Path: {}", path.display());
                        files.push(path_str.to_owned());
                    } else if is_bourne_shell_script(path_str) {
                        // println!("Bourne Shell File Path: {}", path.display());
                        files.push(path_str.to_owned());
                    } else {
                        // Nothing
                    }
                }
            } else if path.is_dir() {
                // Recursively visit subdirectories
                // folder_traverse(&path);
                files.extend(folder_traverse(&path).into_iter());
            }
        }
    }

    files
}

fn main() {

    /**************************
    ** Gather Special Cases **
    **************************/
    let _ignore_file = "./.soteriaignore";
    let special_files = read_lines(_ignore_file);
    // println!("Special-Files:\n{:?}", special_files);

    /**********************
    ** Gather File List **
    *********************/
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        // DO Nothing as of now
    } else {
        process::exit(1);
    }

    let root_dir = args[1].clone();
    let root_path = Path::new(&root_dir);

    let pre_files = folder_traverse(&root_path);
    // println!("Pre-Files:\n{:?}", pre_files);

    let files = cmp_files(pre_files, special_files);
    println!("Final Files:\n{:?}", files);
}
